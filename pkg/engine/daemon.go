package engine

func (e *Engine) Daemon() {
	// Initialize the first fanout tick so that we can keep track of the actually
	// executed interval moving forward.

	{
		e.tic = <-e.rtr.Tick()
	}

	// Run all engine processes concurrently in order to utilize all available
	// host CPUs.
	//
	//       <--join-->      <--join-->  <--join-->            <--join--><--join-->       <--join-->
	//
	//      <-move->   <-move-> <-move->   <-move->  <-move-><-move->  <-move->   <-move-><-move->
	//
	//                      <race>         <race>    <race>                    <race>       <race>
	//
	//     <------------tick------------><------------tick------------><------------tick------------>
	//
	//     0ms                           25ms                          50ms
	//

	for {
		select {
		case <-e.don:
			return
		case x := <-e.rtr.Join():
			go e.join(x)
		case x := <-e.rtr.Move():
			go e.move(x)
		case x := <-e.rtr.Race():
			go e.race(x)
		case x := <-e.rtr.Tick():
			go e.tick(x)
		}
	}
}
