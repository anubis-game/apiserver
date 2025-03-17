package engine

import (
	"slices"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// TODO:infra implement Engine.occupy() just like Engine.screen(), but use
// Vector.Occupy() for the bounds. The screen is then a square, while the occupy
// is a rectangle.

// screen searches for all Vectors w located around the given Vector v's head
// node. The area we are searching through here contains a single layer of large
// partitions around v's head node. That is 9 large partitions in total. Despite
// the fact that Vector v is also located within the given search area, the
// provided callback f is not executed for the Vector v itself.
func (e *Engine) screen(v *vector.Vector, f func(byte, *vector.Vector)) {
	var p matrix.Partition

	a := []byte{v.Uid()}
	t, r, b, l := v.Screen(matrix.Pt8, 1)

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
