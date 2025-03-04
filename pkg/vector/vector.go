package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

type Config struct {
	// Hea contains the new Vector's head.
	Hea matrix.Coordinate
	// Mot contains the new Vector's direction of travel.
	Mot Motion
}

type Vector struct {
	// crx
	crx Charax
	// mot contains this Vector's current direction of travel.
	mot Motion

	// ocl, occupied coordinate list, contains all coordinates occupied by this
	// Vector, grouped by occupied partition.
	ocl map[matrix.Partition][]matrix.Coordinate
	// ocd, occupied coordinate diff,
	ocd map[matrix.Partition][]matrix.Coordinate

	// hea is this Vector's first node. The Vector's head defines a player's
	// direction of travel.
	hea *Linker
	// tai is this Vector's last node.
	tai *Linker
	// mhn, maximum hidden nodes, describes the maximum amount of hidden
	// nodes that this Vector allowes to exist between the nodes of its
	// underlying linked list.
	mhn int8
	len int

	// scr contains all partition coordinates of this Vector's view, which entails
	// everything that this Vector can see right now. A Vector's screen partitions
	// are based on the Vector's head position, which is marked with H in the
	// illustration below.
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
	scr []matrix.Partition
}

func New(c Config) *Vector {
	// Ensure that we start with unique *Linker pointers for the head and tail
	// nodes, so that we do not overwrite the tail when we update the head during
	// initial Vector expansion. Note that we have to initialize the head *Linker
	// with a hidden node count of -1, so that the very first Vector expansion can
	// split the root coordinate without misrepresenting the head's hidden nodes.

	var hea *Linker
	var tai *Linker
	{
		hea = &Linker{crd: c.Hea, hid: -1}
		tai = &Linker{crd: c.Hea}
	}

	{
		hea.prv = tai
		tai.nxt = hea
	}

	return &Vector{
		mot: c.Mot,

		ocl: map[matrix.Partition][]matrix.Coordinate{
			c.Hea.Prt(): {
				c.Hea,
			},
		},
		ocd: map[matrix.Partition][]matrix.Coordinate{},

		hea: hea,
		tai: tai,
		mhn: 3,
		len: 1,
	}
}
