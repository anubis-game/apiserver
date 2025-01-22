package stream

func (s *Stream) delete(pac Packet) {
	delete(s.cli, pac.Cli.Wallet())
}
