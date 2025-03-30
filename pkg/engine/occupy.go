package engine

import (
	"slices"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// occupy searches for all Vectors located around the given Vector v's occupied
// area. The area we are searching through here contains a single layer of large
// partitions around v's occupied area. Despite the fact that Vector v is also
// located within the given search area, the provided callback f is not executed
// for the Vector v itself.
func (e *Engine) occupy(v *vector.Vector, f func(byte, *vector.Vector)) {
	var p matrix.Partition

	a := []byte{v.Uid()}
	t, r, b, l := v.Occupy(matrix.Pt8)

	for y := t; y >= b; y -= int(matrix.Pt8) {
		for x := l; x <= r; x += int(matrix.Pt8) {
			p.X = x
			p.Y = y

			for u := range e.lkp.pt8[p] {
				if !slices.Contains(a, u) {
					f(u, e.mem.vec[u])
					a = append(a, u)
				}
			}
		}
	}
}
