package worker

type Config[T any] struct {
	Don <-chan struct{}
	Rtr Router[T]
}

type Worker[T any] struct {
	don <-chan struct{}
	que chan T
	rtr Router[T]
	sem chan struct{}
}

func NewWorker[T any](c Config[T]) *Worker[T] {
	return &Worker[T]{
		don: c.Don,
		que: make(chan T, 1000),
		rtr: c.Rtr,
		sem: make(chan struct{}, 10),
	}
}

func (w *Worker[T]) Ensure(pac T) {
	w.que <- pac
}

func (w *Worker[T]) Daemon() {
	for {
		select {
		case <-w.don:
			// The injected global done channel may signal a program shutdown. In that
			// case we are not accepting any new packets anymore. Once the global done
			// channel got closed, we simply return. Note that there is an option to
			// explicitly wait for the last packets to be processed, if the worker
			// daemon were to be synchronously integrated as a blocking element.
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
			// process packets at the same time. Every time we receive a packet, we
			// push a ticket into the semaphore before doing the work.
			w.sem <- struct{}{}

			// A new goroutine is created for every piece of work. That way we can
			// work on packets in parallel. Note that the received packet must be
			// injected into the goroutine as an argument, in order to work on the
			// exact packet that we received in this asynchronous environment.
			go func(pac T) {
				// Ensure we remove our ticket from the semaphore once all work was
				// completed.
				defer func() {
					<-w.sem
				}()

				// Forward the current packet to the configured router and wait for the
				// work to be done. The boolean returned will indicate whether the
				// processed packet ought to be processed once more.
				var req bool
				{
					pac, req = w.rtr.Router(pac)
				}

				// Once a packet was processed, we may receive the instruction to
				// requeue that task. In that case, we add the given packet back to the
				// end of the queue, for it to be processed again, until it concludes
				// successfully eventually.
				if req {
					w.que <- pac
				}
			}(x)
		}
	}
}
