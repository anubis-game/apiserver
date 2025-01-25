package engine

func (e *Engine) Daemon() {
	// Initialize the first fanout tick so that we can keep track of the actually
	// executed interval moving forward.
	{
		e.tic = <-e.rtr.Fanout()
	}

	for {
		select {
		case <-e.don:
			return
		case x := <-e.rtr.Create():
			e.create(x)
		case x := <-e.rtr.Delete():
			e.delete(x)
		case x := <-e.rtr.Fanout():
			e.fanout(x)
		}
	}
}
