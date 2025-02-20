package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Mot Motion
	Obj []object.Object
	Uid [2]byte
}

type Vector struct {
	// buf contains the prepared fanout buffers grouped by this Vector's occupied
	// coordinate partitions.
	//
	//                          0                               6
	//     [X: 128, Y: 512]    [0x0, 0x0, 0x2, 0x8, 0x10, 0x2c, 0x0, 0x0, 0x2, 0x8, 0xc, 0x28]
	//
	buf map[object.Object][]byte

	// crx
	crx setter.Interface[Charax]
	// mot
	mot setter.Interface[Motion]

	//
	hea *Linker
	tai *Linker
	len int

	// occ contains information about partition coordinates that this Vector
	// occupies.
	occ *Occupy

	// scr
	scr *Screen

	// Uid is the 2 byte unique identifier for this particular Vector across its
	// entire lifetime.
	uid [2]byte

	// TODO:refactor move those counters to Vector.Screen
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
			buf: map[object.Object][]byte{},

			crx: setter.New[Charax](),
			mot: setter.New[Motion](),

			occ: &Occupy{},
			scr: &Screen{},

			uid: c.Uid,

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
		siz := make([]byte, 4)

		siz[0] = byte(schema.Size)
		copy(siz[1:3], vec.uid[:])
		siz[3] = crx.Rad

		crx.siz = siz
	}

	{
		typ := make([]byte, 4)

		typ[0] = byte(schema.Type)
		copy(typ[1:3], vec.uid[:])
		typ[3] = crx.Rad

		crx.typ = typ
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
		vec.occ.Top = prt.Y
		vec.occ.Rig = prt.X
		vec.occ.Bot = prt.Y
		vec.occ.Lef = prt.X
	}

	{
		vec.xfr[prt.X] = 1
		vec.yfr[prt.Y] = 1
	}

	{
		buf := make([]byte, 4+object.Len)

		buf[0] = byte(schema.Body)
		copy(buf[1:3], vec.uid[:])
		buf[3] = 0x1
		copy(buf[4:], byt[:])

		vec.buf[prt] = buf
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
		vec.len = 1
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
		vec.occ.Old = object.Object{}
		vec.occ.New = object.Object{}
		vec.occ.Prt = nil
	}

	{
		vec.scr.Top = prt.Y + vpb
		vec.scr.Rig = prt.X + vpb
		vec.scr.Bot = prt.Y - vpb
		vec.scr.Lef = prt.X - vpb
		vec.scr.Prt = nil
	}

	for x := vec.scr.Lef; x <= vec.scr.Rig; x += matrix.Prt {
		for y := vec.scr.Bot; y <= vec.scr.Top; y += matrix.Prt {
			vec.scr.Prt = append(vec.scr.Prt, object.Object{X: x, Y: y})
		}
	}

	for x := vec.occ.Lef; x <= vec.occ.Rig; x += matrix.Prt {
		for y := vec.occ.Bot; y <= vec.occ.Top; y += matrix.Prt {
			vec.occ.Prt = append(vec.occ.Prt, object.Object{X: x, Y: y})
		}
	}

	return vec
}
