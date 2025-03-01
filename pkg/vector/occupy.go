package vector

import "github.com/anubis-game/apiserver/pkg/object"

type Occupy struct {
	// Top, Rig, Bot and Lef are the outer partition coordinates of this Vector's
	// body parts.
	//
	//                Top
	//         +---------------+
	//         |          ###H |
	//         |          #    |
	//     Lef | T######  #    | Rig
	//         |       #  #    |
	//         |       ####    |
	//         +---------------+
	//                Bot
	//
	Top int
	Rig int
	Bot int
	Lef int

	// Prt contains a consistent representation of all occupied partitions at all
	// times. Keys are partition coordinates. Values are Vector coordinates.
	Prt map[object.Object][]object.Object
}

// Occupy returns those partition coordinates that are currently occupied by
// this Vector's body parts.
func (v *Vector) Occupy() map[object.Object][]object.Object {
	return v.occ.Prt
}
