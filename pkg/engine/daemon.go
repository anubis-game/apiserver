package engine

func (e *Engine) Daemon() {
	// Initialize the first fanout tick so that we can keep track of the actually
	// executed interval moving forward.

	{
		e.tic = <-e.rtr.Tick()
	}

	// Defining a player's direction to move towards costs about 2.90 ns/op. This
	// cost is constant, and the player instructions are rate limitted per client
	// connection. It is critically important that Engine.turn() is only called
	// sequentially by a single goroutine.

	go func() {
		for t := range e.rtr.Turn() {
			e.turn(t)
		}
	}()

	// Switching a player's velocity costs about 2.00 ns/op. This cost is
	// constant, and the player instructions are rate limitted per client
	// connection. It is critically important that Engine.race() is only called
	// sequentially by a single goroutine.

	go func() {
		for b := range e.rtr.Race() {
			e.race(b)
		}
	}()

	// Since we have more than one function to call within the same fixed
	// interval, we are distributing the ticks across static goroutines using
	// dedicated channels.

	go func() {
		for {
			select {
			case u := <-e.rtr.Uuid():
				e.uuid(u)
			case t := <-e.rtr.Tick():
				e.tick()
				e.send(t)
			}
		}
	}()

	// Block the engine daemon until the program shuts down.

	{
		<-e.don
	}
}
