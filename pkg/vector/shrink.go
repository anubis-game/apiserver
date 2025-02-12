package vector

import "github.com/anubis-game/apiserver/pkg/object"

func (v *Vector) Shrink() {
	tai := v.tai // remember current tail for new tail

	{
		v.tai = tai.nxt // next of tail becomes new tail
		v.len--
	}

	{
		v.shrink(tai.val) // old tail gets cleaned up
	}
}

func (v *Vector) shrink(old object.Object) {
	prt := old.Prt()
	buf := v.buf[prt]

	// Always keep track of the amount of coordinates that do not occupy their
	// respective partitions anymore.

	{
		v.xfr[prt.X]--
		v.yfr[prt.Y]--
	}

	// Reduce the fanout buffer given any of the situations described below.

	if len(buf) == object.Len {

		// There is only one item left. That item is the object we are asked to
		// delete.

		{
			delete(v.buf, prt)
		}

		// Shrink the partition boundaries according to the direction of change as
		// specified by the old tail coordinates.

		tai := v.tai.val.Prt()

		// TODO shrink range of sight

		if prt.Y == v.btp && v.yfr[prt.Y] == 0 {
			{
				v.btp = tai.Y
			}

			{
				delete(v.yfr, prt.Y)
			}
		}
		if prt.X == v.brg && v.xfr[prt.X] == 0 {
			{
				v.brg = tai.X
			}

			{
				delete(v.xfr, prt.X)
			}
		}
		if prt.Y == v.bbt && v.yfr[prt.Y] == 0 {
			v.bbt = tai.Y

			{
				delete(v.yfr, prt.Y)
			}
		}
		if prt.X == v.blf && v.xfr[prt.X] == 0 {
			{
				v.blf = tai.X
			}

			{
				delete(v.xfr, prt.X)
			}
		}
	} else {

		// The item to remove is always the very first part of the buffer. Note that
		// we are only reslicing the given partition buffer, which means that we
		// keep the bytes of the deleted old tail allocated in the underlying data
		// array. This alone would imply a memory leak, but we are fixing this
		// memory leak in due time within Vector.expand() and also once the
		// partition buffer gets deleted entirely.

		{
			v.buf[prt] = buf[object.Len:]
		}
	}
}
