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
func (e *Engine) allpt8(u byte, v *vector.Vector) []byte {
	all := []byte{}

	for _, p := range v.Layers(1, matrix.Pt8) {
		for b := range e.lkp.pt8[p] {
			if u != b && !slices.Contains(all, b) {
				all = append(all, b)
			}
		}
	}

	return all
}
