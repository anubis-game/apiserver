package engine

import "time"

func (e *Engine) tick(tic time.Time) {
	{
		e.move()
		e.send()
	}

	// TODO:metrics monitor the time since tic to see how regular our fanout
	// procedure executes throughout the program lifetime.

	{
		time.Since(tic)
	}
}
