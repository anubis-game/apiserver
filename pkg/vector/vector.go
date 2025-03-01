package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Mot Motion
	Obj []object.Object
}

type Vector struct {
	// crx
	crx setter.Interface[Charax]
	// mot contains this Vector's current direction of travel.
	mot Motion

	//
	hea *Linker
	tai *Linker
	len int

	//
	tl1 object.Object
	tl2 object.Object

	// occ contains information about partition coordinates that this Vector
	// occupies.
	occ *Occupy

	// scr
	scr *Screen

	xfr map[int]int
	yfr map[int]int
}

func New(c Config) *Vector {
	if len(c.Obj) == 0 {
		tracer.Panic(fmt.Errorf("%T.Obj must not be empty", c))
	}

	var vec *Vector
	{
		vec = &Vector{
			crx: setter.New[Charax](),
			mot: c.Mot,

			occ: &Occupy{
				Prt: make(map[object.Object][]object.Object),
			},
			scr: &Screen{
				Prt: make(map[object.Object]struct{}),
			},

			xfr: map[int]int{},
			yfr: map[int]int{},
		}
	}

	// Ensure the character setter tracks the player's default values.

	var crx Charax
	{
		crx = Charax{
			Als: byte(Ai),
			Alr: byte(Ai / 2),
			Prt: int(Pi),
			Rad: Rad,
			Siz: Siz,
			Typ: 0, // TODO:game configure the player skin/suit, randomly or preference
		}
	}

	{
		vec.crx.Set(crx)
	}

	// Initialize the vector with the first coordinate partition, so that we are
	// able to further track the Vector's occupied boundaries. Note that the first
	// coordinate within Config.Obj represents the Vector's tail, which means we
	// have to index the initial coordinates in reverse order.

	var hea object.Object
	var prt object.Object
	{
		hea = c.Obj[len(c.Obj)-1]
		prt = hea.Prt()
	}

	{
		vec.occ.Top = prt.Y
		vec.occ.Rig = prt.X
		vec.occ.Bot = prt.Y
		vec.occ.Lef = prt.X
	}

	{
		vec.xfr[prt.X] = 1
		vec.yfr[prt.Y] = 1
	}

	// Setting the head and tail elements to the very same pointer reference
	// allows Vector.Expand() below to set the tail's next neighbour, effectively
	// pointing to the expanded head.

	var lin *Linker
	{
		lin = &Linker{
			val: hea,
		}
	}

	{
		vec.hea = lin
		vec.tai = lin
		vec.len = 1
	}

	var vpb int
	{
		vpb = crx.Prt * matrix.Prt
	}

	{
		vec.scr.Top = prt.Y + vpb
		vec.scr.Rig = prt.X + vpb
		vec.scr.Bot = prt.Y - vpb
		vec.scr.Lef = prt.X - vpb
	}

	// Add all injected objects properly to this vector by registering the
	// injected coordinates and expanding this vector's window accordingly.

	for i := len(c.Obj) - 2; i >= 0; i-- {
		vec.Expand(c.Obj[i])
	}

	return vec
}
