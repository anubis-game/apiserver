package engine

func (e *Engine) Daemon() {
	// Initialize the first fanout tick so that we can keep track of the actually
	// executed interval moving forward.

	{
		e.tic = <-e.rtr.Tick()
	}

	// Synchronize all state management for a lock free architecture.

	go func() {
		for {
			select {
			case u := <-e.rtr.Uuid():
				e.uuid(u)
			case b := <-e.rtr.Race():
				e.race(b)
			case b := <-e.rtr.Turn():
				e.turn(b)
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
