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

	// TODO:infra we probably need to maintain a consistent representation of all
	// visible partitions at all times.
	Prt map[object.Object]struct{}
}

// Screen returns the boundary information of this Vector's view, expressed in
// partition coordinates.
func (v *Vector) Screen() *Screen {
	return v.scr
}
