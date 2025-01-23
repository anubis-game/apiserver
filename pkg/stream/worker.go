package stream

import (
	"runtime"
)

type Worker struct {
	sem chan struct{}
}

// TODO move worker to its own package
func NewWorker() *Worker {
	return &Worker{
		sem: make(chan struct{}, runtime.NumCPU()),
	}
}

func (w *Worker) Worker(fnc func()) {
	// The semaphore controls the amount of workers that are allowed to process
	// packets at the same time. Every time we receive a packet, we push a ticket
	// into the semaphore before doing the work.
	{
		w.sem <- struct{}{}
	}

	{
		fnc()
	}

	// Ensure we remove our ticket from the semaphore once all work was completed.
	{
		<-w.sem
	}
}
