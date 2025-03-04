package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

func (v *Vector) Oclist(prt matrix.Partition) []matrix.Coordinate {
	return v.ocl[prt]
}

func (v *Vector) Ocdiff(prt matrix.Partition) []matrix.Coordinate {
	return v.ocd[prt]
}

// occRem removes data about the partitions and coordinates that this Vector
// occupies. The given coordinate represents the Vector's old tail node, that
// got removed from the underlying linked list.
func (v *Vector) occRem(old matrix.Coordinate) {
	prt := old.Prt()

	if len(v.ocl[prt]) == 1 {
		// There is only one coordinate left in the removed tail partition. Since we
		// have been asked to delete remove the given tail, we delete the reference
		// to the entire partition.

		{
			delete(v.ocl, prt)
		}

		// At this point the Vector is leaving the occupied partition, because the
		// tail that we are removing now was the last coordinate occupying it. We
		// track the tail coordinates internally, so that we can report on the delta
		// created during this update cycle.

		{
			v.ocd[prt] = append(v.ocd[prt], old)
		}

	} else {
		// The coordinate to remove is always at index zero. Note that we are only
		// reslicing the existing coordinate array, without deleting the remaining
		// data still allocated. This by itself does usually bear the risk of a
		// memory leak. We are working around this memory leak by deleting the
		// coordinate array, as soon as this Vector moves out of the occupied
		// partition when playing the game.

		{
			v.ocl[prt] = v.ocl[prt][1:]
		}
	}
}

// occAdd adds data about the partitions and coordinates that this Vector
// occupies. The coordinates we are working with internally represent the
// Vector's current and previous head node. It is therefore important that
// Vector.occAdd() is only called after a new head node got added to this
// Vector.
func (v *Vector) occAdd(rsz bool) {
	cur := v.hea.crd
	prv := v.hea.prv.crd

	chp := cur.Prt()
	php := prv.Prt()

	if len(v.ocl[chp]) == 0 {
		// Initialize a new slice of occupied partition coordinates and add the
		// current head node to it.

		{
			v.ocl[chp] = []matrix.Coordinate{cur}
		}

		// Track the head coordinates internally, so that we can report on the delta
		// created during this update cycle.

		{
			v.ocd[chp] = append(v.ocd[chp], cur)
		}

		// Only if the new head breaks into an unoccupied partition, and only if the
		// player's range of sight remains unchanged, only then do we have to check
		// in which direction the new head node is overflowing, because then we want
		// to adjust the player's visible range of sight.

		if !rsz {
			row := (v.crx.Fos * 2) + 1

			if chp.Y > php.Y {
				{
					copy(v.scr[row:], v.scr[:]) // move the first rows one row south
				}
				for i := range row {
					v.scr[i].Y += matrix.Prt // update the top row
				}
			}

			if chp.X > php.X {
				for i := 0; i < len(v.scr); i += row {
					copy(v.scr[i:i+row-1], v.scr[i+1:i+row]) // move all rows one column west
				}
				for i := row - 1; i < len(v.scr); i += row {
					v.scr[i].X += matrix.Prt // update the right column
				}
			}

			if chp.Y < php.Y {
				{
					copy(v.scr[:], v.scr[row:]) // move the last rows one row north
				}
				for i := range len(v.scr) - row {
					v.scr[i].Y -= matrix.Prt // update the bottom row
				}
			}

			if chp.X < php.X {
				for i := 0; i < len(v.scr); i += row {
					copy(v.scr[i+1:i+row], v.scr[i:i+row-1]) // move all rows one column east
				}
				for i := 0; i < len(v.scr); i += row {
					v.scr[i].X -= matrix.Prt // update the left column
				}
			}
		}
	} else {
		// Extend the current buffer with the compressed 6 byte version of the given
		// coordinates. Using preallocated slices via copy safes about 5 ns/op
		// compared to using append. Note that using a new preallocated byte slice
		// fixes the memory leak incurred during Vector.shrink() where we merely
		// reslice the partition buffer.

		{
			v.ocl[chp] = append(v.ocl[chp], cur)
		}
	}
}
