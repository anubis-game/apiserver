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

	// New is the newly occupied partition.
	New object.Object
	// Old is the partition that this Vector left most recently.
	Old object.Object

	// Prt only contains the vectors initially occupied partition coordinates.
	// Meaning, this list is only filled once during Vector creation. In order to
	// see the Vector's updated occupied partitions, check the values of New and
	// Old.
	Prt []object.Object
}

// Occupy returns those partition coordinates that are currently occupied by
// this Vector's body parts.
func (v *Vector) Occupy() *Occupy {
	return v.occ
}
