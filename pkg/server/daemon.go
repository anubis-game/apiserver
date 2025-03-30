package server

func (s *Server) Daemon() {
	// Initialize several random queues. Note that those services are used in
	// various places, but they are exclusively initialized here in the server.

	{
		go s.agl.Daemon()
		go s.crd.Daemon()
		go s.qdr.Daemon()
	}

	{
		s.config()
		s.router()
		s.listen()
	}

	// Wait for the done channel to close, and then close the listener of this
	// server to prevent any further connections to be opened.

	{
		<-s.don
	}

	{
		s.lis.Close() // nolint:errcheck
	}
}
