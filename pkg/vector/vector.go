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
	//
	mot setter.Interface[Motion]

	//
	hea *Linker
	tai *Linker
	len int

	// btp, brg, bbt and blf are the outer boundaries of this Vector's body,
	// expressed in partition coordinates, that this Vector keeps track of.
	//
	//                btp
	//         +---------------+
	//         |          #### |
	//         |          #    |
	//     blf | #######  #    | brg
	//         |       #  #    |
	//         |       ####    |
	//         +---------------+
	//                bbt
	//
	btp int
	brg int
	bbt int
	blf int

	// TODO add Vector.Occupy() []object.Object to return the partition
	// coordinates that this Vector occupies right now and ensure that Vector.bpb
	// is properly maintained over this Vector's lifetime.
	bpb []object.Object

	// vtp, vrg, vbt and vlf are the outer boundaries of this Vector's view,
	// expressed in partition coordinates, that this Vector keeps track of.
	//
	//                vtp
	//         +---------------+
	//         |               |
	//         |               |
	//     vlf |    ####       | vrg
	//         |    #          |
	//     #######  #          |
	//         +-#--#----------+
	//           #### vbt
	//
	vtp int
	vrg int
	vbt int
	vlf int

	// vpb is this Vector's movement based list of partition boundaries. The
	// partitions listed here represent the Vector's most recently reveiled view
	// due to a player's movement across the coordinate system. In other words,
	// the partitions here represent what the player can now see, because the
	// player was moving towards a direction that was unexplored before.
	vpb []object.Object

	//
	xfr map[int]int
	yfr map[int]int

	// buf contains the prepared fanout buffers grouped by this Vector's occupied
	// coordinate partitions.
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
			crx: setter.New[Charax](),
			mot: setter.New[Motion](),

			hea: nil,
			tai: nil,
			len: 1,

			btp: 0,
			brg: 0,
			bbt: 0,
			blf: 0,

			vtp: 0,
			vrg: 0,
			vbt: 0,
			vlf: 0,

			xfr: map[int]int{},
			yfr: map[int]int{},
			buf: map[object.Object][]byte{},
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
			Typ: 0, // TODO randomize or configure the player suit based on the user's preference
		}
	}

	{
		vec.crx.Set(crx)
	}

	// Ensure the motion setter tracks the injected configuration.

	{
		vec.mot.Set(c.Mot)
	}

	// Initialize the vector with the first coordinate partition, so that we are
	// able to further track the vector's occupied boundaries. Note that this
	// first coordinate represents the Vector's tail.

	var prt object.Object
	var byt [object.Len]byte
	{
		tai := c.Obj[0]
		prt = tai.Prt()
		byt = tai.Byt()
	}

	{
		vec.btp = prt.Y
		vec.brg = prt.X
		vec.bbt = prt.Y
		vec.blf = prt.X
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

	// Expand the screen boundaries of this Vector based on the header coordinates
	// that we come up with after having expanded this Vector using all of the
	// injected coordinates. Note that technically the last coordinate object
	// becomes the head of our linked list.

	{
		prt = vec.hea.val.Prt()
	}

	var vpb int
	{
		vpb = crx.Prt * matrix.Prt
	}

	{
		vec.vtp = prt.Y + vpb
		vec.vrg = prt.X + vpb
		vec.vbt = prt.Y - vpb
		vec.vlf = prt.X - vpb
	}

	// We have to reset the partition boundary slice because the initial calls to
	// Vector.Expand() fill it with invalid values that would pollute the correct
	// view otherwise.

	{
		vec.vpb = nil
	}

	for x := vec.vlf; x <= vec.vrg; x += matrix.Prt {
		for y := vec.vbt; y <= vec.vtp; y += matrix.Prt {
			vec.vpb = append(vec.vpb, object.Object{X: x, Y: y})
		}
	}

	return vec
}
