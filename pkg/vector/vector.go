package vector

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/anubis-game/apiserver/pkg/window"
)

type Config struct {
	Mot Motion
	Obj []object.Object
}

type Vector struct {
	ind int
	mot setter.Interface[Motion]
	obj []object.Object
	win *window.Window
}

func New(c Config) *Vector {
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
	}

	if len(c.Obj) != 0 {
		win.Ini(c.Obj[0])
	}

	// Add all injected objects properly to this vector by registering the
	// injected coordinates and expanding this vector's window accordingly.
	// Allocating 100 points per segment times 10,000 segments in total allows for
	// 1,000,000 points per player.

	var obj []object.Object
	{
		obj = make([]object.Object, 10_000)
	}

	var ind int
	for _, x := range c.Obj {
		win.Inc(x)
		obj[ind] = x
		ind++
	}

	return &Vector{
		ind: ind,
		mot: mot,
		obj: obj,
		win: win,
	}
}
