package client

func (c *Client) Daemon() {
	// Process every message that this client receives in its own goroutine. Any
	// write error causes the client connection to terminate. The buffer channel
	// defines a capacity of 1024 pending messages, but we terminate the client
	// connection already after 25% saturation. 40 messages may acumulate every
	// second at a frame rate of 25 ms/s. That means roughly 200 messages may
	// accumulate within a 5 second window, which means that we have at least 3
	// intervals to decide whether we want to terminate a congested client
	// connection.

	go func() {
		for b := range c.buf {
			if c.Stream(b) != nil {
				close(c.wri)
				return
			}
		}
	}()

	// Ensure the client connection gets terminated on excessive saturation.
	// Every 5 seconds we check for the current amount of buffer congestion. We
	// allow every client to accumulate 256 pending messages before we close the
	// ticker channel below, which then triggers the client termination in the
	// server handler.

	go func() {
		for range c.tiC {
			if len(c.buf) >= c.cap {
				close(c.tic)
				return
			}
		}
	}()
}
