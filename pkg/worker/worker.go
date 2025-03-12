package worker

import (
	"fmt"
	"runtime"
	"time"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/transaction"
	"github.com/anubis-game/apiserver/pkg/worker/action"
	"github.com/anubis-game/apiserver/pkg/worker/record"
	"github.com/ethereum/go-ethereum/common"
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
	seq chan action.Interface
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
		all: make([]action.Interface, 0, 5000),
		don: c.Don,
		log: c.Log,
		que: make(chan action.Interface, 5000),
		reg: c.Reg,
		seq: make(chan action.Interface, 1),
		sig: sig,
	}
}

// TODO:test write some unit tests for this entire reconciliation complexity

func (w *Worker) Daemon() {
	// We need to persist every action whether it is going to succeed or not,
	// so that we can provide a full record of all worker actions being
	// processed throughout the Guardian's lifetime.

	// TODO:infra we have to expose and maybe even backup all of those actions,
	// because those actions represent the gains and losses of the winners and
	// losers.

	go func() {
		for a := range w.seq {
			w.all = append(w.all, a)
		}
	}()

	// Create a static worker pool to distribute work across all available host
	// CPUs.

	// TODO:infra the game has to stop if the action queue ever fills up due to
	// network outages, because the static worker pool may saturate at some point.

	// TODO:infra add some RPC failover mechanism so we can switch to a healthy
	// RPC provider.

	for range runtime.NumCPU() {
		go func() {
			for a := range w.que {
				// Trying to attach status updates to an empty record set would result
				// in a panic. So before we can process the given action, we have to add
				// a new record for us to always be able to attach any relevant status
				// updates.

				{
					a.Rec().Add()
				}

				// Process the given action and attach any relevant status updates to
				// the latest record.

				{
					err := w.worker(a)
					if err != nil {
						{
							a.Rec().Err().Set(err)
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

				if a.Rec().Sta().Get() != record.Success {
					var fai int

					// TODO:test verify that we count the number of failed actions
					// properly.

					for i := range a.Rec().Len() {
						if a.Rec().Get(i).Sta().Get() == record.Failure {
							fai++
						}
					}

					// We do not want to sign more than 3 transactions per action if
					// signing keeps failing. So, when an action failed twice already, we
					// re-schedule for another attempt until the given action either
					// succeeds, or fails a third time, which is when we give up on it.

					if fai < 3 {
						time.AfterFunc(a.Rec().Wai().Get(), func() {
							w.que <- a
						})
					}
				}
			}
		}()
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
			w.seq <- act
			w.que <- act
		}
	}
}

func (w *Worker) worker(act action.Interface) error {
	var err error

	var txn common.Hash
	var emp bool
	{
		txn = act.Rec().Txn().Get()
		emp = transaction.Empty(txn)
	}

	// There are two distinct cases in which we want to sign a new transaction.
	// The first case happens for all new actions, for which we have to sign the
	// first transaction. The second case happens if an already signed transaction
	// was found to have failed. Important here is to not call act.Rec().Prv() in
	// the first case, because record.Slicer.Prv() goes back 2 records in the
	// given action's history. When we sign the first transaction, then there is
	// only 1 record in the action history, meaning going back 2 while there is
	// only 1 results in a panic. The same applies when we set the record status
	// to either "created" or "retried".

	if emp || act.Rec().Prv().Sta().Get() == record.Failure {
		{
			txn, err = w.sig[act.Typ()].Sign(act.Arg())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if emp {
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
		_, err = w.reg.Search(txn)
	}

	if registry.IsTransactionNotFoundError(err) {
		act.Rec().Sta().Set(record.Waiting)
	} else if registry.IsTransactionStillPending(err) {
		act.Rec().Sta().Set(record.Waiting)
	} else if registry.IsTransactionNotSuccessfulError(err) {
		act.Rec().Sta().Set(record.Failure)
	}

	// We want to return in case any error occurs at all.

	if err != nil {
		return tracer.Mask(err)
	}

	{
		act.Rec().Sta().Set(record.Success)
	}

	return nil
}
