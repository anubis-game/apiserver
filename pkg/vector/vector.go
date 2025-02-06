package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/anubis-game/apiserver/pkg/window"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Mot Motion
	Obj []object.Object
}

type Vector struct {
	hea *Linker
	mot setter.Interface[Motion]
	siz int
	tai *Linker
	win *window.Window
}

func New(c Config) *Vector {
	if len(c.Obj) == 0 {
		tracer.Panic(fmt.Errorf("%T.Obj must not be empty", c))
	}

	var mot setter.Interface[Motion]
	{
		mot = setter.New[Motion]()
	}

	{
		mot.Set(c.Mot)
	}

	// If there is at least one coordinate object for this vector, then we have to
	// initialize the vector's window with any available object, so that the
	// window itself is further able to keep track of this vector's occupied
	// boundaries.

	var win *window.Window
	{
		win = window.New()
		win.Ini(c.Obj[0])
	}

	var hea *Linker
	var siz int
	var tai *Linker
	{
		lin := &Linker{val: c.Obj[0]}
		hea = lin
		tai = lin
		siz = 1
	}

	// Add all injected objects properly to this vector by registering the
	// injected coordinates and expanding this vector's window accordingly.

	for _, x := range c.Obj[1:] {
		lin := &Linker{val: x}

		hea.nxt = lin
		hea = lin
		siz++

		win.Inc(x)
	}

	return &Vector{
		hea: hea,
		mot: mot,
		siz: siz,
		tai: tai,
		win: win,
	}
}
