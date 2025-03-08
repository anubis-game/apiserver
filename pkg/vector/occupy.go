package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

// occAdd adds data about the partitions and coordinates that this Vector
// occupies. The coordinates we are working with internally represent the
// Vector's current and previous head node. It is therefore important that
// Vector.occAdd() is only called after a new head node got added to this
// Vector.
func (v *Vector) occAdd() {
	cur := v.hea.crd
	prv := v.hea.prv.crd

	chp := cur.Prt()
	php := prv.Prt()

	// Always keep track of the amount of partition coordinates that the nodes of
	// this Vector occupy along the X and Y axis.

	{
		v.ofx[chp.X]++
		v.ofy[chp.Y]++
	}

	if !chp.Eql(php) {
		// Only if the new head breaks into an unoccupied partition, only then do we
		// have to check in which direction the new head node is overflowing,
		// because then we want to adjust the player's visible range of sight.
		//
		//     1. If the current head breaks north, then increment top and bottom.
		//
		//     2. If the current head breaks east, then increment left and right.
		//
		//     3. If the current head breaks south, then decrement top and bottom.
		//
		//     4. If the current head breaks west, then decrement left and right.
		//

		if chp.Y > php.Y {
			v.otp = chp.Y
			v.stp += matrix.Prt
			v.sbt += matrix.Prt
		}

		if chp.X > php.X {
			v.org = chp.X
			v.slf += matrix.Prt
			v.srg += matrix.Prt
		}

		if chp.Y < php.Y {
			v.obt = chp.Y
			v.stp -= matrix.Prt
			v.sbt -= matrix.Prt
		}

		if chp.X < php.X {
			v.olf = chp.X
			v.slf -= matrix.Prt
			v.srg -= matrix.Prt
		}
	}
}

// occRem removes data about the partitions and coordinates that this Vector
// occupies. The given coordinate represents the Vector's old tail node, that
// got removed from the underlying linked list.
func (v *Vector) occRem(prv matrix.Coordinate) {
	ctp := v.tai.crd.Prt()
	ptp := prv.Prt()

	// Always keep track of the amount of partition coordinates that the nodes of
	// this Vector occupy along the X and Y axis.

	{
		v.ofx[ptp.X]--
		v.ofy[ptp.Y]--
	}

	if !ctp.Eql(ptp) {
		// Shrink the occupied partition coordinates according to the direction of
		// change as specified by the previous tail node.

		if ptp.Y == v.otp && v.ofy[ptp.Y] == 0 {
			v.otp = ctp.Y
			delete(v.ofy, ptp.Y)
		}

		if ptp.X == v.org && v.ofx[ptp.X] == 0 {
			v.org = ctp.X
			delete(v.ofx, ptp.X)
		}

		if ptp.Y == v.obt && v.ofy[ptp.Y] == 0 {
			v.obt = ctp.Y
			delete(v.ofy, ptp.Y)
		}

		if ptp.X == v.olf && v.ofx[ptp.X] == 0 {
			v.olf = ctp.X
			delete(v.ofx, ptp.X)
		}
	}
}
