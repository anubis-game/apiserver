package stream

func (s *Stream) delete(add string) {
	var old Client

	{
		s.mut.Lock()
		old = s.cli[add]
		delete(s.cli, add)
		s.mut.Unlock()
	}

	{
		old.Close()
	}
}
