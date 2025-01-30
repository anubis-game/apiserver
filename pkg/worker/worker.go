package worker

import (
	"fmt"
	"runtime"
	"time"

	"github.com/anubis-game/apiserver/pkg/cache"
	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/worker/action"
	"github.com/anubis-game/apiserver/pkg/worker/record"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don <-chan struct{}
	Log logger.Interface
	Reg *registry.Registry
	Sig []Signer
}

type Worker struct {
	all []action.Interface
	don <-chan struct{}
	log logger.Interface
	que chan action.Interface
	reg *registry.Registry
	req *cache.Time[uuid.UUID]
	sem chan struct{}
	sig map[string]Signer
}

func New(c Config) *Worker {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}
	if c.Reg == nil {
		tracer.Panic(fmt.Errorf("%T.Reg must not be empty", c))
	}
	if len(c.Sig) == 0 {
		tracer.Panic(fmt.Errorf("%T.Sig must not be empty", c))
	}

	sig := map[string]Signer{}
	for _, x := range c.Sig {
		sig[x.Type()] = x
	}

	return &Worker{
		all: []action.Interface{},
		don: c.Don,
		log: c.Log,
		que: make(chan action.Interface, 5000),
		reg: c.Reg,
		req: cache.NewTime[uuid.UUID](),
		sem: make(chan struct{}, runtime.NumCPU()),
		sig: sig,
	}
}

// TODO write some unit tests for this entire reconciliation complexity

func (w *Worker) Daemon() {
	// Setup the re-queue cache to check all expiration callbacks every so often.

	{
		go w.req.Expire(time.Second)
	}

	for {
		select {
		case <-w.don:
			// The injected global done channel may signal a program shutdown. In that
			// case we are not accepting any new actions anymore. Once the global done
			// channel got closed, we simply return below. Note that there is an
			// option to explicitly wait for the last action to be processed, if the
			// program's process would not give ample time to gracefully shutdown the
			// entire system.
			//
			//     for len(w.sem) > 0 {
			//       time.Sleep(500 * time.Millisecond)
			//     }
			//
			//     {
			//       close(w.sem)
			//     }
			//

			return
		case x := <-w.que:
			// The semaphore controls the amount of workers that are allowed to
			// process actions at the same time. Every time we receive an action, we
			// push a ticket into the semaphore before doing the work.

			{
				w.sem <- struct{}{}
			}

			// We need to persist every action whether it is going to succeed or not,
			// so that we can provide a full record of all worker actions being
			// processed throughout the Guardian's lifetime.

			{
				w.all = append(w.all, x)
			}

			// A new goroutine is created for every piece of work. That way we can
			// work on actions in parallel. Note that the received action must be
			// injected into the goroutine as an argument, in order to work on the
			// exact action that we received in this asynchronous environment.

			go func(act action.Interface) {
				// Trying to attach status updates to an empty record set would result
				// in a panic. So before we can process the given action, we have to add
				// a new record for us to always be able to attach any relevant status
				// updates.

				{
					act.Rec().Add()
				}

				// Process the given action and attach any relevant status updates to
				// the latest record.

				{
					err := w.worker(act)
					if err != nil {
						{
							act.Rec().Err().Set(err)
						}

						w.log.Log(
							"level", "error",
							"message", err.Error(),
							"stack", tracer.Stack(err),
						)
					}
				}

				// Unless we receive a success status, we add the given action back to
				// the internal queue, effectively retrying it once more.

				if act.Rec().Sta().Get() != record.Success {
					var fai int

					for i := 0; i < act.Rec().Len(); i++ {
						if act.Rec().Sta().Get() == record.Failure {
							fai++
						}
					}

					// We do not want to sign more than 3 transactions per action if
					// signing keeps failing. So when an action failed twice already, we
					// re-schedule for another attempt until the given action either
					// succeeds, or fails a third time, which is when we give up on it.

					if fai < 3 {
						w.req.Ensure(act.Uid(), act.Rec().Wai().Get(), func() {
							w.que <- act
						})
					}
				}

				// Ensure we remove our ticket from the semaphore once all work was
				// completed.

				{
					<-w.sem
				}
			}(x)
		}
	}
}

func (w *Worker) Ensure(act action.Interface) {
	var exi bool
	{
		_, exi = w.sig[act.Typ()]
	}

	if !exi {
		w.log.Log(
			"level", "warning",
			"message", fmt.Sprintf("invalid action type %q", act.Typ()),
		)
	} else {
		{
			w.que <- act
		}
	}
}

func (w *Worker) worker(act action.Interface) error {
	var err error

	// There are two distinct cases in which we want to sign a new transaction.
	// The first case happens for all new actions, for which we have to sign the
	// first transaction. The second case happens if an already signed transaction
	// was found to have failed. Important here is to not call act.Rec().Prv() in
	// the first case, because record.Slicer.Prv() goes back 2 records in the
	// given action's history. When we sign the first transaction, then there is
	// only 1 record in the action history, meaning going back 2 while there is
	// only 1 results in a panic. The same applies when we set the record status
	// to either "created" or "retried".

	if act.Rec().Txn().Emp() || act.Rec().Prv().Sta().Get() == record.Failure {
		var txn common.Hash
		{
			txn, err = w.sig[act.Typ()].Sign(act.Arg())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if act.Rec().Txn().Emp() {
			act.Rec().Sta().Set(record.Created)
		} else {
			act.Rec().Sta().Set(record.Retried)
		}

		{
			act.Rec().Txn().Set(txn)
		}

		return nil
	}

	// At this point we have to verify the latest transaction of the given action,
	// which has either the status "created", "waiting", or "retried".

	{
		_, err = w.reg.Search(act.Rec().Txn().Get())
	}

	if registry.IsTransactionNotFoundError(err) {
		act.Rec().Sta().Set(record.Waiting)
	} else if registry.IsTransactionStillPending(err) {
		act.Rec().Sta().Set(record.Waiting)
	} else if registry.IsTransactionNotSuccessfulError(err) {
		act.Rec().Sta().Set(record.Failure)
	}

	if err != nil {
		return tracer.Mask(err)
	}

	{
		act.Rec().Sta().Set(record.Success)
	}

	return nil
}
