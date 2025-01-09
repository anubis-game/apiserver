package stream

func (s *Stream) cast(byt []byte) error {
	s.cli.Ranger(func(_ string, val Client) {
		val.Write(byt)
	})

	return nil
}
