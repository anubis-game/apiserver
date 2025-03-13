package engine

import (
	"slices"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// TODO:infra add the uid to Vector again

// TODO:infra rename Vector.Bounds() to Vector.Screen() again

// TODO:infra only take v in Engine.screen()

// TODO:infra implement Engine.occupy() just like Engine.screen(), but use
// Vector.Occupy() for the bounds. The screen is then a square, while the occupy
// is a rectangle.

// screen searches for all the unique byte IDs located around the given Vector's
// head node. The area we are searching through here contains a single layer of
// large partitions around the given Vector's head node. That is 9 large
// partitions in total. The provided callback f is not executed for the Vector
// associated with the byte ID u.
func (e *Engine) screen(u byte, h matrix.Partition, f func(byte, *vector.Vector)) {
	var prt matrix.Partition

	top := h.Y + matrix.Pt8
	rig := h.X + matrix.Pt8
	bot := h.Y - matrix.Pt8
	lef := h.X - matrix.Pt8

	all := []byte{u}
	for y := top; y >= bot; y -= matrix.Pt8 {
		for x := lef; x <= rig; x += matrix.Pt8 {
			prt.X = x
			prt.Y = y

			for b := range e.lkp.pt8[prt] {
				if !slices.Contains(all, b) {
					f(b, e.mem.vec[b])
					all = append(all, b)
				}
			}
		}
	}
}
