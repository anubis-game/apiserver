package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Mot Motion
	Obj []object.Object
}

type Vector struct {
	//
	mot setter.Interface[Motion]

	//
	hea *Linker
	tai *Linker
	siz int

	// top, rig, bot and lef are the outer boundaries, expressed in partition
	// coordinates, that this Mapper keeps track of.
	//
	//                top
	//         +---------------+
	//         |          #### |
	//         |          #    |
	//     lef | #######  #    | rig
	//         |       #  #    |
	//         |       ####    |
	//         +---------------+
	//                bot
	//
	top int
	rig int
	bot int
	lef int

	//
	xfr map[int]int
	yfr map[int]int

	// buf contains the prepared fanout buffers grouped by coordinate partitions.
	//
	//                          0                               6
	//     [X: 128, Y: 512]    [0x0, 0x0, 0x2, 0x8, 0x10, 0x2c, 0x0, 0x0, 0x2, 0x8, 0xc, 0x28]
	//
	buf map[object.Object][]byte
}

func New(c Config) *Vector {
	if len(c.Obj) == 0 {
		tracer.Panic(fmt.Errorf("%T.Obj must not be empty", c))
	}

	var vec *Vector
	{
		vec = &Vector{
			mot: setter.New[Motion](),

			hea: nil,
			siz: 1,
			tai: nil,

			top: 0,
			rig: 0,
			bot: 0,
			lef: 0,

			xfr: map[int]int{},
			yfr: map[int]int{},
			buf: map[object.Object][]byte{},
		}
	}

	// Ensure the motion setter tracks the injected configuration.

	{
		vec.mot.Set(c.Mot)
	}

	// Initialize the vector with the first coordinate partition, so that we are
	// able to further track the vector's occupied boundaries.

	var hea object.Object
	var prt object.Object
	var byt [6]byte
	{
		hea = c.Obj[0]
		prt = hea.Prt()
		byt = hea.Byt()
	}

	{
		vec.top = prt.Y
		vec.rig = prt.X
		vec.bot = prt.Y
		vec.lef = prt.X
	}

	{
		vec.xfr[prt.X] = 1
		vec.yfr[prt.Y] = 1
	}

	{
		vec.buf[prt] = byt[:]
	}

	// Setting the head and tail elements to the very same pointer reference
	// allows Vector.Expand() below to set the tail's next neighbour, effectively
	// pointing to the expanded head.

	var lin *Linker
	{
		lin = &Linker{
			val: c.Obj[0],
		}
	}

	{
		vec.hea = lin
		vec.tai = lin
	}

	// Add all injected objects properly to this vector by registering the
	// injected coordinates and expanding this vector's window accordingly.

	for _, x := range c.Obj[1:] {
		vec.Expand(x)
	}

	return vec
}
