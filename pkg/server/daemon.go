package server

func (s *Server) Daemon() {
	s.config()
	s.router()
	s.listen()
}
