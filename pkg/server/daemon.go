package server

func (s *Server) Daemon() {
	s.config()
	s.router()
	s.listen()

	// Wait for the done channel to close, and then close the listener of this
	// server to prevent any further connections to be opened.

	<-s.don

	s.lis.Close() // nolint:errcheck
}
