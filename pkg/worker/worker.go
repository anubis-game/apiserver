package worker

import (
	"runtime"
	"time"

	"github.com/anubis-game/apiserver/pkg/cache"
)

type Config[K comparable, P any] struct {
	Don <-chan struct{}
	Ens Ensure[K, P]
}

type Worker[K comparable, P any] struct {
	don <-chan struct{}
	ens Ensure[K, P]
	que chan P
	req *cache.Time[K]
	sem chan struct{}
}

func New[K comparable, P any](c Config[K, P]) *Worker[K, P] {
	return &Worker[K, P]{
		don: c.Don,
		ens: c.Ens,
		que: make(chan P, 1000),
		req: cache.NewTime[K](),
		sem: make(chan struct{}, runtime.NumCPU()),
	}
}

func (o *Worker[K, P]) Create(pac P) {
	o.que <- pac
}

func (o *Worker[K, P]) Daemon() {
	// Setup the re-queue cache to check all expiration callbacks every so often.
	{
		go o.req.Expire(time.Second)
	}

	for {
		select {
		case <-o.don:
			// The injected global done channel may signal a program shutdown. In that
			// case we are not accepting any new packets anymore. Once the global done
			// channel got closed, we simply return. Note that there is an option to
			// explicitly wait for the last packets to be processed, if the worker
			// daemon were to be synchronously integrated as a blocking element.
			//
			//     for len(o.sem) > 0 {
			//       time.Sleep(500 * time.Millisecond)
			//     }
			//
			//     {
			//       close(o.sem)
			//     }
			//
			return
		case x := <-o.que:
			// The semaphore controls the amount of workers that are allowed to
			// process packets at the same time. Every time we receive a packet, we
			// push a ticket into the semaphore before doing the work.
			{
				o.sem <- struct{}{}
			}

			// A new goroutine is created for every piece of work. That way we can
			// work on packets in parallel. Note that the received packet must be
			// injected into the goroutine as an argument, in order to work on the
			// exact packet that we received in this asynchronous environment.
			go func(pac P) {
				// Forward the current packet to the configured router and wait for the
				// work to be done. The timeout returned will indicate whether the
				// processed packet ought to be processed once more. If no timeout is
				// returned, then the given packet is considered processed successfully.
				var key K
				var ttl time.Duration
				{
					pac, key, ttl = o.ens.Ensure(pac)
				}

				// Once a packet was processed, we may receive the instruction to
				// requeue that task. In that case, we add the given packet back to the
				// end of the queue once the given timeout passed, for the given packet
				// to be processed again.
				if ttl != 0 {
					o.req.Ensure(key, ttl, func() {
						o.que <- pac
					})
				}

				// Ensure we remove our ticket from the semaphore once all work was
				// completed.
				{
					<-o.sem
				}
			}(x)
		}
	}
}
