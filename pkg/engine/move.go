package engine

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// Move can theoretically be run concurrently since we only read data to update
// each Vector once, but if we do this, we have to wait for every single Vector
// to be updated.
func (e *Engine) move() {
	for u := range e.uni.Length() {
		if e.fcn[u] == nil {
			continue
		}

		var v *vector.Vector
		{
			v = e.mvc[u]
		}

		var tur router.Turn
		{
			tur = e.tur[u]
		}

		// It may happen that a new player is being processed here, while said
		// player has not yet provided their own motion specific update. In such a
		// case the new player's quadrant byte is still empty, forcing us to move
		// the player along the game map using the current direction of travel. We
		// will consider the player choise of movement as soon as they provide their
		// own motion update.

		if tur.Qdr == 0 {
			mot := v.Motion()
			tur.Qdr = mot.Qdr
			tur.Agl = mot.Agl
		}

		var rac byte
		{
			rac = e.rac[u]
		}

		// TODO:infra manage energy delta, eating food increases size, racing mode decreases size
		//
		//     find energy objects in front of head
		//     use hypothetical head
		//     calculate via Impact()
		//     track energy delta for Vector
		//

		var del int
		{
			//
		}

		var hea object.Object
		var tl1 object.Object
		var tl2 object.Object
		{
			hea, tl1, tl2 = v.Adjust(del, tur.Qdr, tur.Agl, rac)
		}

		// Add the byte ID to the head's partition.

		{
			prt := hea.Prt()
			_, exi := e.pvc[prt]

			if !exi {
				e.pvc[prt] = map[byte]struct{}{u: struct{}{}}
			} else {
				e.pvc[prt][u] = struct{}{}
			}
		}

		// TODO:infra the code below is wrong. There are two things to consider.
		//
		//     1. The heads and tails represent changes that every relevant client
		//        screen has to be informed about.
		//
		//     2. A tail being removed does not always mean for a byte ID to be
		//        removed from the partition, because more Vector segments might
		//        still be in that partition.
		//

		// Remove the byte ID from the tail partitions, if any of the tails has been
		// removed from the Vector.

		if !tl1.Zer() {
			delete(e.pvc[tl1.Prt()], u)
		}

		if !tl2.Zer() {
			delete(e.pvc[tl2.Prt()], u)
		}
	}
}
