package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func (e *Engine) move() {
	for ua := range e.uni.Length() {
		// Skip all inactive players.

		if !e.ply.act[ua] {
			continue
		}

		// TODO:infra show new energy packets within the relevant screens

		// Get the Vector va that we are reconciling now.

		var va *vector.Vector
		{
			va = e.mem.vec[ua]
		}

		// Update this Vector.

		{
			va.Update(0, e.ply.qdr[ua], e.ply.agl[ua], e.ply.rac[ua]) // TODO:infra the eaten food must be added
		}

		// Get the Vector's changed coordinates.

		var c vector.Change
		{
			c = va.Change()
		}

		// Add its own Vector changes to its own fanout buffer.

		{
			e.bufslf(ua, c)
		}

		// Add the Vector changes to the lookup tables.

		{
			e.lookup(ua, c)
		}

		// Look for all byte IDs and their associated Vectors that are located near
		// va's new head node.

		e.screen(va, func(ub byte, vb *vector.Vector) {
			// Check whether va or vb gets killed upon impact.

			{
				e.impact(va, vb)
			}

			// Add the Vector changes of va to the fanout buffer of vb, but only if vb is
			// known to be connected.

			if e.ply.cli[ub] != nil {
				e.bufply(c, ub, vb)
			}
		})
	}
}

func (e *Engine) bufply(va vector.Change, ub byte, vb *vector.Vector) {
	t, r, b, l := vb.Screen(matrix.Pt1)

	if va.Hea.Pt1().Ins(t, r, b, l) {
		e.ply.buf[ub] = append(e.ply.buf[ub], 0x0) // TODO:infra encode head create message properly
	}

	for _, c := range va.Rem {
		if c.Pt1().Ins(t, r, b, l) {
			e.ply.buf[ub] = append(e.ply.buf[ub], 0x0) // TODO:infra encode tail delete message properly
		}
	}
}

func (e *Engine) bufslf(ua byte, va vector.Change) {
	{
		e.ply.buf[ua] = append(e.ply.buf[ua], 0x0) // TODO:infra encode head create message properly
	}

	for range va.Rem {
		e.ply.buf[ua] = append(e.ply.buf[ua], 0x0) // TODO:infra encode tail remove message properly
	}
}

func (e *Engine) impact(va *vector.Vector, vb *vector.Vector) {
	// We need to define the partition coordinates to search through for the
	// impact check below using the factor of sight f. The search area can be
	// smaller as long as the body part radii of Vector va and vb combined fit into
	// the required layer of small partitions. E.g. Rv=25 Rw=17 requires only a
	// single layer because 42 is smaller than 128.

	var t, r, b, l int
	{
		t, r, b, l = va.Screen(matrix.Pt1, ((va.Charax().Fos+vb.Charax().Fos)/byte(matrix.Pt1))+1)
	}

	// Iterate over all node coordinates of Vector vb that are close to Vector va's
	// head node.

	vb.Inside(t, r, b, l, func(c matrix.Coordinate) bool {
		if va.Impact(c, vb.Charax().Rad) {
			// The head node of Vector va has collided with the node coordinate c of
			// Vector vb. By default this means that Vector va gets killed, because that
			// player va was running into the body of another player vb. The exception
			// here is that Vector vb gets killed instead, if the colliding node c is
			// the head node of Vector vb, while the head node of Vector va is larger.
			// In other words, if two heads collide, then the larger player wins.

			if c == vb.Change().Hea && va.Charax().Rad > vb.Charax().Rad {
				// TODO:infra kill vb, create energy in place, update surrounding screens
				return false
			}

			{
				// TODO:infra kill va, create energy in place, update surrounding screens
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
