package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

type lookup struct {
	// TODO:infra we should remove the second map value and put the coordinates
	// and bytes into slices.
	nrg map[matrix.Partition]map[matrix.Coordinate]struct{}
	pt1 map[matrix.Partition]map[byte]struct{}
	pt8 map[matrix.Partition]map[byte]struct{}
}

func newLookup(_ int) *lookup {
	return &lookup{
		nrg: map[matrix.Partition]map[matrix.Coordinate]struct{}{},
		pt1: map[matrix.Partition]map[byte]struct{}{},
		pt8: map[matrix.Partition]map[byte]struct{}{},
	}
}

func (l *lookup) add(u byte, c matrix.Coordinate) {
	// Add the given node coordinate to the small partitions.

	{
		p := c.Pt1()

		m, e := l.pt1[p]
		if !e {
			m = map[byte]struct{}{u: {}}
		} else {
			m[u] = struct{}{}
		}

		l.pt1[p] = m
	}

	// Add the given node coordinate to the large partitions.

	{
		p := c.Pt8()

		m, e := l.pt8[p]
		if !e {
			m = map[byte]struct{}{u: {}}
		} else {
			m[u] = struct{}{}
		}

		l.pt8[p] = m
	}
}

func (l *lookup) rem(u byte, c matrix.Coordinate) {
	// Remove the given node coordinate from the small partitions.

	{
		delete(l.pt1[c.Pt1()], u)
	}

	// Remove the given node coordinate from the large partitions.

	{
		delete(l.pt8[c.Pt8()], u)
	}
}
