package engine

import (
	"slices"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// allpt8 searches for all the unique byte IDs located around the given Vector's
// head node. The area we are searching through here contains a single layer of
// large partitions around the given Vector's head node. That is 9 large
// partitions in total. The returned map of unique byte IDs does not include the
// given byte ID u.
func (e *Engine) allpt8(u byte, h matrix.Partition, p int, f func(byte, *vector.Vector)) {
	var prt matrix.Partition

	top := h.Y + p
	rig := h.X + p
	bot := h.Y - p
	lef := h.X - p

	all := []byte{u}
	for y := top; y >= bot; y -= p {
		for x := lef; x <= rig; x += p {
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
