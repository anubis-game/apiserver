package stream

func (s *Stream) add(add string, cli Client) {
	s.mut.Lock()
	old, exi := s.cli[add]
	s.cli[add] = cli
	s.mut.Unlock()

	if exi {
		old.Close(true)
	}
}
