package stream

import (
	"github.com/anubis-game/apiserver/pkg/schema"
)

func (s *Stream) update() {
	for k, v := range s.cli {
		go s.wrk.Worker(func() {
			// TODO add energy to wallet address, if any
			nrg, _ := s.nrg.Load(k)
			for _, x := range nrg {
				v.Stream(schema.Encode(schema.Food, x.Bytes()))
			}

			// TODO add player movement to wallet address, if any
			ply, _ := s.ply.Load(k)
			for _, x := range ply {
				v.Stream(schema.Encode(schema.Move, x.Bytes()))
			}
		})
	}
}
