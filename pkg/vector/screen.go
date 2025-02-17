package vector

import "github.com/anubis-game/apiserver/pkg/object"

type Screen struct {
	// Top, Rig, Bot and Lef are the outer partition coordinates of this Vector's
	// view that this Vector can see right now, based on the Vector's header
	// position, which is marked with H in the illustration below.
	//
	//                Top
	//         +---------------+
	//         |               |
	//         |               |
	//     Lef |    ###H       | Rig
	//         |    #          |
	//     T######  #          |
	//         +-#--#----------+
	//           #### Bot
	//
	Top int
	Rig int
	Bot int
	Lef int

	// Prt represents the most recently discovered partition coordinates. For a
	// new Vector those partitions represent the Vector's entire range of sight.
	// For Vectors being expanded or rotated, those partitions represent the slice
	// of partition coordinates on the screen that have just been revealed by
	// movement towards any given direction.
	Prt []object.Object
}

// Screen returns the boundary information of this Vector's view, expressed in
// partition coordinates.
func (v *Vector) Screen() *Screen {
	return v.scr
}
