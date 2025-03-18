package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// energy calls the given function with the energy bytes that are located within
// the given Vector's screen.
func (e *Engine) energy(v *vector.Vector, f func([]byte)) {
	var p matrix.Partition

	t, r, b, l := v.Screen(matrix.Pt1)

	for y := t; y >= b; y -= int(matrix.Pt1) {
		for x := l; x <= r; x += int(matrix.Pt1) {
			p.X = x
			p.Y = y

			for c := range e.lkp.nrg[p] {
				f(e.mem.nrg[c])
			}
		}
	}
}
