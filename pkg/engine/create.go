package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func (e *Engine) create(pac router.Packet) {
	var ply *player.Player
	{
		ply = player.New(player.Config{
			Bck: matrix.Bucket{
				e.fil.crd.Random(), // x0
				e.fil.crd.Random(), // y0
				e.fil.crd.Random(), // x1
				e.fil.crd.Random(), // y1
			},
			Cli: pac.Cli,
			Pxl: matrix.Pixel{
				e.fil.crd.Random(), // x2
				e.fil.crd.Random(), // y2
			},
			Spc: matrix.Space{
				e.fil.qdr.Random(), // quadrant
				e.fil.ang.Random(), // angle
			},
			Uid: pac.Uid,
		})
	}

	{
		e.mem.ply[pac.Uid] = ply
	}

	// Put the player randomly onto the game map for every relevant player to see.
	// Adding the new player information to every buffer of every player causes
	// the next fanout cycle to push the new player into all relevant views. For
	// the confirmation of joining the game, we send the player wallet together
	// with its object information, so that wallet and uuid can be associated in
	// the client.

	var byt []byte
	{
		byt = schema.Encode(schema.Join, pac.Cli.Wallet().Bytes(), ply.Bytes())
	}

	for k, v := range e.mem.ply {
		// Only add an additional buffer to the existing player v if the new player
		// ply is in the view of v.
		if v.Win.Has(ply.Obj.Bck) {
			e.buf.ply.Compute(k, func(old [][]byte, _ bool) ([][]byte, bool) {
				return append(old, byt), false
			})
		}
	}

	// TODO we need to stream all relevant map details for the initial view to
	// ply. That means to find all relevant energy and player details on the
	// current map.

	// TODO add new player to the lookup map based on its current coordinates

}
