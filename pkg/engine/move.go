package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func (e *Engine) move() {
	for u := range e.uni.Length() {
		// Skip all inactive players.

		if !e.ply.act[u] {
			continue
		}

		// Get the Vector v that we are reconciling now.

		var v *vector.Vector
		{
			v = e.mem.vec[u]
		}

		// Update this Vector.

		{
			v.Update(0, e.ply.qdr[u], e.ply.agl[u], e.ply.rac[u]) // TODO:infra the eaten food must be added
		}

		// Get the Vector's changed coordinates.

		var c vector.Change
		{
			c = v.Change()
		}

		// Add its own Vector changes to its own fanout buffer.

		{
			e.bufslf(u, c)
		}

		// Add the Vector changes to the lookup tables.

		{
			e.lookup(u, c)
		}

		// Look for all byte IDs and their associated Vectors that are located near
		// v's new head node.

		e.screen(u, v.Change().Hea.Pt8(), func(b byte, w *vector.Vector) {
			// Check whether v or w gets killed upon impact.

			{
				e.impact(v, w)
			}

			// Add the Vector changes of v to the fanout buffer of w, but only if w is
			// known to be connected.

			if e.ply.cli[b] != nil {
				e.bufply(c, b, w)
			}
		})
	}
}

func (e *Engine) bufply(c vector.Change, b byte, w *vector.Vector) {
	l, m, n, o := w.Bounds()

	if c.Hea.Pt1().Ins(l, m, n, o) {
		e.ply.buf[b] = append(e.ply.buf[b], 0x0) // TODO:infra encode head create message properly
	}

	for _, t := range c.Rem {
		if t.Pt1().Ins(l, m, n, o) {
			e.ply.buf[b] = append(e.ply.buf[b], 0x0) // TODO:infra encode tail delete message properly
		}
	}
}

func (e *Engine) bufslf(u byte, c vector.Change) {
	{
		e.ply.buf[u] = append(e.ply.buf[u], 0x0) // TODO:infra encode head create message properly
	}

	for range c.Rem {
		e.ply.buf[u] = append(e.ply.buf[u], 0x0) // TODO:infra encode tail remove message properly
	}
}

func (e *Engine) impact(v *vector.Vector, w *vector.Vector) {
	// We need to define the partition coordinates to search through for the
	// impact check below using the factor of sight f. The search area can be
	// smaller as long as the body part radii of Vector v and w combined fit into
	// the required layer of small partitions. E.g. Rv=25 Rw=17 requires only a
	// single layer because 42 is smaller than 128.

	var l, m, n, o int
	{
		l, m, n, o = v.Bounds(((v.Charax().Fos + w.Charax().Fos) / matrix.Pt1) + 1)
	}

	// Iterate over all node coordinates of Vector w that are close to Vector v's
	// head node.

	w.Inside(l, m, n, o, func(c matrix.Coordinate) bool {
		if v.Impact(c, w.Charax().Rad) {
			// The head node of Vector v has collided with the node coordinate c of
			// Vector w. By default this means that Vector v gets killed, because that
			// player v was running into the body of another player w. The exception
			// here is that Vector w gets killed instead, if the colliding node c is
			// the head node of Vector w, while the head node of Vector v is larger.
			// In other words, if two heads collide, then the larger player wins.

			if c == w.Change().Hea && v.Charax().Rad > w.Charax().Rad {
				// TODO:infra kill w, break loops
				return false
			}

			{
				// TODO:infra kill v, break loops
				return false
			}
		}

		return true
	})
}

func (e *Engine) lookup(u byte, c vector.Change) {
	{
		e.lkp.add(u, c.Hea)
	}

	for _, t := range c.Rem {
		e.lkp.rem(u, t)
	}
}
