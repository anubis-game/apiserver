package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

type Config struct {
	// Hea contains the new Vector's head.
	Hea matrix.Coordinate
	// Mot contains the new Vector's direction of travel.
	Mot Motion
	// Uid is the unique byte ID associated to the player operating this Vector.
	Uid byte
}

type Vector struct {
	// crx
	crx Charax
	// mot contains this Vector's current direction of travel.
	mot Motion
	// ocd, occupied coordinate diff, contains the coordinates of the new head
	// node, the new tail node, and any removed tail node, as generated by any
	// given update cycle, in this order, grouped by occupied partition.
	ocd Change
	// otp, org, obt and olf describe the outer partition coordinates that the nodes
	// of this Vector occupy. Everything a Vector occupies is based on every
	// single Vector node, which are illustrated from head to tail in the diagram
	// below.
	//
	//                otp
	//         +---------------+
	//         |          ###H |
	//         |          #    |
	//     olf | T######  #    | org
	//         |       #  #    |
	//         |       ####    |
	//         +---------------+
	//                obt
	//
	otp, org, obt, olf int

	len int
	// mhn, maximum hidden nodes, describes the maximum amount of hidden
	// nodes that this Vector allowes to exist between the nodes of its
	// underlying linked list.
	mhn int8

	// hea is this Vector's first node. The Vector's head defines a player's
	// direction of travel.
	hea *Linker
	// tai is this Vector's last node.
	tai *Linker

	// ofx and ofy contain the frequency counts of the occupied partition
	// coordinates for X and Y respectively.
	ofx, ofy map[int]int

	// uid
	uid byte
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

	var prt matrix.Partition
	var fos byte
	{
		prt = c.Hea.Pt1()
		fos = sight(0)
	}

	return &Vector{
		crx: Charax{Fos: fos},
		mot: c.Mot,

		hea: hea,
		tai: tai,
		mhn: 3,
		len: 1,

		ocd: Change{Hea: hea.crd, Tai: tai.crd},
		otp: prt.Y,
		org: prt.X,
		obt: prt.Y,
		olf: prt.X,
		ofx: map[int]int{prt.X: 1},
		ofy: map[int]int{prt.Y: 1},

		uid: c.Uid,
	}
}
