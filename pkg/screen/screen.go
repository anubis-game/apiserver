package screen

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/schema"
)

type Screen struct {
	obj [][]object.Object
}

func New() *Screen {
	obj := make([][]object.Object, 255)
	for i := range obj {
		obj[i] = []object.Object{}
	}

	return &Screen{
		obj: obj,
	}
}

// TODO:test ensure that we do not miss any updates if Buffer() is called
// sequentially and Update() is called concurrently.

func (s *Screen) Buffer() []byte {
	var buf []byte

	for u := range byte(255) {
		// TODO:bench check if the extra handling of v is faster than simply appending to buf

		l := len(s.obj[u])
		v := make([]byte, 3+(l*object.Len))

		{
			v[0] = byte(schema.Body)
			v[1] = u
			v[2] = byte(l)
		}

		for i, o := range s.obj[u] {
			p := 3 + (i * object.Len)
			b := o.Byt()
			copy(v[p:p+object.Len], b[:])
		}

		{
			buf = append(buf, v...)
		}
	}

	return buf
}

func (s *Screen) Create(uid byte, obj []object.Object) {
	//
}

func (s *Screen) Delete(uid byte) {
	//
}

func (s *Screen) Update(uid byte, hea object.Object, tl1 object.Object, tl2 object.Object) {
	// Update() is called iteratively
	// within the same cycle, Update() may be called multiple times with the same uid
	// we need to create an initial representation using schema.Body
	// consecutive calls need to find the delta using schema.Head and schema.Tail
}
