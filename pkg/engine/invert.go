package engine

import (
	"slices"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// invert removes the given Vector from the game map and adds the equivalent
// amount of energy packets in its place.
func (e *Engine) invert(v *vector.Vector) {
	// Calculate the amount of energy that we should generate per Vector node.

	var b int
	{
		b = v.Charax().Siz / v.Length()
	}

	// Range over all Vector nodes.

	var a []matrix.Coordinate
	var i int

	v.Ranger(func(c matrix.Coordinate) {
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

			// TODO track new energy packets separately for this update cycle

			// TODO track new energy packets globally

			// TODO find all connected players that can see this inversion
			//
			//     * those players must be informed about a player being killed
			//
			//     * those players must render the new energy packets in place
			//

			return true
		})
	})
}
