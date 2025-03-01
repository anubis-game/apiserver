package engine

import "time"

func (e *Engine) Daemon() {
	// We are constantly filling the game map with energy packages and place
	// different sizes of different types randomly onto the game map. Smaller
	// energy appears more often than bigger energy.

	go func() {
		for {
			select {
			case <-time.Tick(2 * time.Second):
				e.food(e.fil.Energy(2))
			case <-time.Tick(5 * time.Second):
				e.food(e.fil.Energy(5))
			case <-time.Tick(10 * time.Second):
				e.food(e.fil.Energy(10))
			}
		}
	}()

	// Joining a game incurs at least 8,000 ns/op for the first player. This
	// process becomes more expensive the more players are active within the
	// same partition coordinates.

	go func() {
		for u := range e.rtr.Uuid() {
			e.uuid(u)
		}
	}()

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

	// Initialize the first fanout tick so that we can keep track of the actually
	// executed interval moving forward.

	{
		e.tic = time.Now().UTC()
	}

	// TODO:docs move, diff, send

	go func() {
		for range e.rtr.Tick() {
			e.move() // move every player on the map      (read turn/race, write Vector)
			e.send() // forward fanout buffers to writer  (read/write fanout buffers)
		}
	}()

	// Block the engine daemon until the program shuts down.

	{
		<-e.don
	}
}
