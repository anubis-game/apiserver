package vector

import "github.com/anubis-game/apiserver/pkg/object"

type Config struct {
	Mot Motion
	Obj []Object
}

// TODO we also need to track the window expansion somehow

type Vector struct {
	len int
	mot object.Interface[Motion]
	obj []Object
}

func New(c Config) *Vector {
	v := &Vector{
		len: 0,
		mot: object.New[Motion](),
		// 100 points * 100 points per segment * 1000 segments = 10,000,000 points per player
		obj: make([]Object, 1000),
	}

	{
		v.mot.Set(c.Mot)
	}

	for _, x := range c.Obj {
		v.obj[v.len] = x
		v.len++
	}

	return v
}
