package engine

func (e *Engine) send() {
	for u := range e.uni.Length() {
		var cli chan<- []byte
		{
			cli = e.ply.cli[u]
		}

		// Skip all disconnected players, whether they are active or not.

		if cli == nil {
			continue
		}

		// Forward the fanout buffer of this update cycle to the client specific
		// goroutine for capacity aware processing. The buffer channels provided by
		// each client must never block.

		{
			cli <- e.ply.buf[u]
		}

		// Reset the player specific fanout buffer, but keep the existing sequence
		// byte and increment it for the next cycle. We have to allocate a new data
		// array in order to prevent race conditions between the engine and client.
		// Active players without a connected client are not processed at this
		// point, because their fanout channel will be nil as checked above. That is
		// why the fanout buffers of those types of players must not be written to
		// as long as they remain active in this disconnected state.

		{
			e.ply.buf[u] = []byte{e.ply.buf[u][0] + 1}
		}
	}
}
