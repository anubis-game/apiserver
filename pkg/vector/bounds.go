package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Bounds returns the outer boundaries of this Vector's current view measured in
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
func (v *Vector) Bounds(fos ...int) (int, int, int, int) {
	// The Vector's view is based on the Vector's current head node.

	var chp matrix.Partition
	{
		chp = v.hea.crd.Pt1()
	}

	// The default factor of sight is used of there was no parameter provided.

	var pxl int
	if len(fos) == 1 {
		pxl = fos[0] * matrix.Pt1
	} else {
		pxl = v.crx.Fos * matrix.Pt1
	}

	return chp.Y + pxl, chp.X + pxl, chp.Y - pxl, chp.X - pxl
}
