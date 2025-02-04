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
	len int
	mot setter.Interface[Motion]
	obj []object.Object
	win *window.Window
}

func New(c Config) *Vector {
	v := &Vector{
		len: 0,
		mot: setter.New[Motion](),
		// 100 points per segment * 10,000 segments = 1,000,000 points per player
		obj: make([]object.Object, 10_000),
		win: &window.Window{},
	}

	{
		v.mot.Set(c.Mot)
	}

	for _, x := range c.Obj {
		v.win.Inc(x)
		v.obj[v.len] = x
		v.len++
	}

	return v
}
