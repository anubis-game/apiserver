package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Screen returns the outer boundaries of this Vector's current view measured in
// partition coordinates. The returned boundaries describe top, right, bottom
// and left, in this order. Everything a Vector can see is based on the Vector's
// head position, which is marked with H in the illustration below.
//
//	           top
//	    +---------------+
//	    |               |
//	    |               |
//	lef |    ###H       | rig
//	    |    #          |
//	T######  #          |
//	    +-#--#----------+
//	      #### bot
func (v *Vector) Screen(prt matrix.PartitionLength, fos ...byte) (int, int, int, int) {
	// The Vector's view is based on the Vector's current head node partition.

	var chp matrix.Partition
	if prt == matrix.Pt8 {
		chp = v.hea.crd.Pt8()
	} else {
		chp = v.hea.crd.Pt1()
	}

	// The default factor of sight is used of there was no parameter provided.

	var pxl int
	if len(fos) == 1 {
		pxl = int(fos[0]) * int(prt)
	} else {
		pxl = int(v.crx.Fos) * int(prt)
	}

	return chp.Y + pxl, chp.X + pxl, chp.Y - pxl, chp.X - pxl
}
