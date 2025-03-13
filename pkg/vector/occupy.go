package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

func (v *Vector) Occupy() (int, int, int, int) {
	return v.otp, v.org, v.obt, v.olf
}

// occAdd adds data about the partitions and coordinates that this Vector
// occupies. The coordinates we are working with internally represent the
// Vector's current and previous head node. It is therefore important that
// Vector.occAdd() is only called after a new head node got added to this
// Vector.
func (v *Vector) occAdd() {
	chp := v.hea.crd.Pt1()
	php := v.hea.prv.crd.Pt1()

	// Always keep track of the amount of partition coordinates that the nodes of
	// this Vector occupy along the X and Y axis.

	{
		v.ofx[chp.X]++
		v.ofy[chp.Y]++
	}

	// If the new head node breaks into an unoccupied partition, then set the
	// new value.

	if !chp.Eql(php) {
		if chp.Y > php.Y {
			v.otp = chp.Y
		}

		if chp.X > php.X {
			v.org = chp.X
		}

		if chp.Y < php.Y {
			v.obt = chp.Y
		}

		if chp.X < php.X {
			v.olf = chp.X
		}
	}
}

// occRem removes data about the partitions and coordinates that this Vector
// occupies. The given coordinate represents the Vector's old tail node, that
// got removed from the underlying linked list.
func (v *Vector) occRem(prv matrix.Coordinate) {
	ctp := v.tai.crd.Pt1()
	ptp := prv.Pt1()

	// Always keep track of the amount of partition coordinates that the nodes of
	// this Vector occupy along the X and Y axis.

	{
		v.ofx[ptp.X]--
		v.ofy[ptp.Y]--
	}

	// Shrink the occupied partition coordinates according to the direction of
	// change as specified by the previous tail node.

	if !ctp.Eql(ptp) {
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
