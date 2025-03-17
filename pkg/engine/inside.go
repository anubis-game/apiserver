package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func (e *Engine) inside(f []byte, v *vector.Vector, u byte, t int, r int, b int, l int) []byte {
	{
		f = append(f, byte(schema.Body), u, 0x0)
	}

	var a int
	{
		a = len(f) - 1
	}

	v.Inside(t, r, b, l, func(c matrix.Coordinate) bool {
		{
			f[a]++
		}

		{
			x := c.Byt()
			f = append(f, x[:]...)
		}

		return true
	})

	return f
}
