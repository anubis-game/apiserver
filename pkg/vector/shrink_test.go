package vector

import (
	"slices"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Shrink(t *testing.T) {
	var one object.Object
	var two object.Object
	var thr object.Object
	var fou object.Object
	{
		one = object.Object{X: 100, Y: 100}
		two = object.Object{X: 100, Y: 150}
		thr = object.Object{X: 100, Y: 200}
		fou = object.Object{X: 100, Y: 250}
	}

	var vec *Vector
	{
		vec = New(Config{
			Obj: []object.Object{
				one, // T
				two,
				thr,
				fou, // H
			},
		})
	}

	if vec.len != 4 {
		t.Fatalf("expected %#v got %#v", 4, vec.len)
	}

	var act []object.Object
	vec.Ranger(func(x object.Object) {
		act = append(act, x)
	})

	var exp []object.Object
	{
		exp = []object.Object{
			one,
			two,
			thr,
			fou,
		}
	}

	if !slices.Equal(act, exp) {
		t.Fatalf("expected %#v got %#v", exp, act)
	}

	{
		vec.Shrink()
	}

	if vec.len != 3 {
		t.Fatalf("expected %#v got %#v", 3, vec.len)
	}

	act = nil
	vec.Ranger(func(x object.Object) {
		act = append(act, x)
	})

	{
		exp = []object.Object{
			two,
			thr,
			fou,
		}
	}

	if !slices.Equal(act, exp) {
		t.Fatalf("expected %#v got %#v", exp, act)
	}
}
