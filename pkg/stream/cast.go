package stream

import "github.com/ethereum/go-ethereum/common"

func (s *Stream) cast(byt []byte) error {
	s.cli.Ranger(func(_ common.Address, val Client) {
		val.Write(byt)
	})

	return nil
}
