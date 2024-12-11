package stream

func (s *Stream) remove(add string) {
	s.mut.Lock()
	old, exi := s.cli[add]
	delete(s.cli, add)
	s.mut.Unlock()

	if exi {
		old.Close(false)
	}
}
