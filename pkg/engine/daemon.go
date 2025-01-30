package engine

func (e *Engine) Daemon() {
	// Initialize the first fanout tick so that we can keep track of the actually
	// executed interval moving forward.
	{
		e.tic = <-e.rtr.Push()
	}

	for {
		select {
		case <-e.don:
			return
		case x := <-e.rtr.Join():
			e.join(x)
		case x := <-e.rtr.Move():
			e.move(x)
		case x := <-e.rtr.Race():
			e.race(x)
		case x := <-e.rtr.Push():
			e.push(x)
		}
	}
}
