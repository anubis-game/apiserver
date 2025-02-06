package vector

import (
	"container/list"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/anubis-game/apiserver/pkg/window"
)

type Config struct {
	Mot Motion
	Obj []object.Object
}

type Vector struct {
	lis *list.List
	mot setter.Interface[Motion]
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

	var lis *list.List
	{
		lis = list.New()
	}

	if len(c.Obj) != 0 {
		{
			win.Ini(c.Obj[0])
			lis.PushFront(c.Obj[0])
		}

		for _, x := range c.Obj[1:] {
			lis.PushFront(x)
			win.Inc(x)
		}
	}

	return &Vector{
		lis: lis,
		mot: mot,
		win: win,
	}
}
