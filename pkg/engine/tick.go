package engine

import (
	"time"
)

func (e *Engine) tick(tic time.Time) {
	{
		e.food()
		e.move()
		e.send()
	}

	{
		e.cyc++
	}

	// TODO:metrics monitor the time since tic to see how regular our fanout
	// procedure executes throughout the program lifetime.

	{
		time.Since(tic)
	}
}
