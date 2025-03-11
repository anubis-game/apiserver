package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Screen returns the outer partition coordinates of this Vector's current view.
// The returned boundaries describe top, right, bottom and left, in this order.
// Everything a Vector can see is based on the Vector's head position, which is
// marked with H in the illustration below.
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
func (v *Vector) Screen(fos int) (int, int, int, int) {
	chp := v.hea.crd.Pt1()
	pxl := fos * matrix.Pt1

	return chp.Y + pxl, chp.X + pxl, chp.Y - pxl, chp.X - pxl
}
