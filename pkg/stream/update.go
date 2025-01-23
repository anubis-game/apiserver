package stream

import (
	"github.com/anubis-game/apiserver/pkg/schema"
)

func (s *Stream) update() {
	for k, v := range s.cli {
		go s.wrk.Worker(func() {
			// Send player movements of the enemies first so that every player can
			// react based on the full picture of the current frame.

			// TODO add prepared player bytes to wallet address, if any
			ply, _ := s.ply.Load(k)
			for _, x := range ply {
				v.Stream(schema.Encode(schema.Move, x))
			}

			// TODO check for wallet specific movement and calculate Target(). We
			// cannot just fanout a prepared byte slice here, since we have to force
			// the player movement in either the currently chosen, or the latest known
			// direction.
			v.Window()

			// Send energy changes last, since player updates are more relevant.

			// TODO add prepared energy bytes to wallet address, if any
			nrg, _ := s.nrg.Load(k)
			for _, x := range nrg {
				v.Stream(schema.Encode(schema.Food, x))
			}
		})
	}
}
