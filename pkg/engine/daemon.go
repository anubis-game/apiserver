package engine

import "time"

func (e *Engine) Daemon() {
	// Initialize the first fanout tick so that we can keep track of the actually
	// executed interval moving forward.

	{
		e.tic = <-e.rtr.Tick()
	}

	// Joining a game incurs at least 8,000 ns/op for the first player. This
	// process becomes more expensive the more players are active within the
	// same partition coordinates.

	go func() {
		for p := range e.rtr.Uuid() {
			e.uuid(p)
		}
	}()

	// Defining a player's direction to move towards costs about 2.90 ns/op. This
	// cost is constant, and the player instructions are rate limitted per client
	// connection. It is critically important that Engine.turn() is only called
	// sequentially by a single goroutine.

	go func() {
		for p := range e.rtr.Turn() {
			e.turn(p)
		}
	}()

	// Switching a player's velocity costs about 2.00 ns/op. This cost is
	// constant, and the player instructions are rate limitted per client
	// connection. It is critically important that Engine.race() is only called
	// sequentially by a single goroutine.

	go func() {
		for p := range e.rtr.Race() {
			e.race(p)
		}
	}()

	// Since we have more than one function to call within the same fixed
	// interval, we are distributing the ticks across static goroutines using
	// dedicated channels.

	var sen chan time.Time
	var tic chan struct{}
	{
		sen = make(chan time.Time, 1)
		tic = make(chan struct{}, 1)
	}

	go func() {
		for t := range e.rtr.Tick() {
			sen <- t
			tic <- struct{}{}
		}
	}()

	// Sending the prepared fanout buffer to a single player costs about 3,900
	// ns/op, which is mainly due to some quirky websocket overhead. We want to
	// serve about 250 players concurrently, which means that we have to
	// distribute code execution across all available host CPUs. The way this
	// distribution works is by sending the prepared fanout buffers to a client
	// specific goroutine, which is specifically maintained throughout the
	// client's lifetime, for the sole purpose of writing to the client's own
	// websocket connection. Important here is that Engine.send() is guaranteed to
	// execute sequentially, in order to guarantee the accurate reading and
	// resetting of the client specific fanout buffers.

	go func() {
		for t := range sen {
			e.send(t)
		}
	}()

	// Adjust the game state on every tick.

	go func() {
		for range tic {
			e.tick()
		}
	}()

	// Block the engine daemon until the program shuts down.

	{
		<-e.don
	}
}
