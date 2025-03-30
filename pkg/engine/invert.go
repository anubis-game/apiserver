package engine

import (
	"slices"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// invert removes the Vector of the losing player from the game map and adds the
// equivalent amount of energy packets in its place. Vector va is the winner and
// Vector vb is the loser.
func (e *Engine) invert(va *vector.Vector, vb *vector.Vector) {
	// Calculate the amount of energy that we should generate per Vector node.

	var a []matrix.Coordinate
	var v []*vector.Vector
	var b int
	var i int
	{
		b = vb.Charax().Siz / vb.Length()
	}

	var ua byte
	var ub byte
	{
		ua = va.Uid()
		ub = vb.Uid()
	}

	e.occupy(vb, func(uc byte, vc *vector.Vector) {
		if e.ply.cli[uc] != nil {
			// Collect the Vectors of all connected players so that they may render
			// the new energy packets in place of the losing player.

			{
				v = append(v, vc)
			}

			// All connected players watching the kill must be informed about the
			// winning and losing player.

			{
				e.ply.buf[uc] = append(e.ply.buf[uc], byte(schema.Kill), ua, ub)
			}
		}
	})

	// Range over all Vector nodes.

	vb.Ranger(func(c matrix.Coordinate) {
		// We use a slice for a more performant existence check, so that we
		// guarantee the generation of unique energy packets. In order to keep this
		// deduplication mechanism most runtime performant, we reset the slice once
		// Vector nodes do not overlap anymore. That works because nodes that do not
		// overlap cannot produce duplicated coordinates.

		if i < 3 {
			i++
		} else {
			a = nil
			i = 0
		}

		vector.Circle(c, b, 5, func(c matrix.Coordinate) bool {
			// We may generated duplicated coordinates because the radii of
			// consecutive Vector nodes do always overlap. If a coordinate was
			// generated already, then we do not deduct from the per node energy
			// budget, and by consequence generate more coordinates by circling
			// further out.

			if slices.Contains(a, c) {
				return false
			}

			{
				a = append(a, c)
			}

			var n []byte
			{
				n = energy.New(energy.Config{
					Crd: c,
					Siz: 5,
					Typ: 0, // TODO:infra randomize energy type
				})
			}

			// Add the new energy packet to the memory table so that we can find all
			// energy packets in existence later on.

			{
				e.mem.nrg[c] = n
			}

			// Collect all energy packets in a single byte slice, so that we can
			// inform all relevant players about the inversion, and increment the
			// amount of energy packets collected.

			for _, x := range v {
				if c.Pt1().Ins(x.Screen(matrix.Pt1)) {
					f = []byte{byte(schema.Food), 0x0}
					f[1]++
					f = append(f, n...)
				}
			}

			return true
		})
	})
}
