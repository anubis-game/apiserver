package vector

import (
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Window_Change(t *testing.T) {
	//
	//     +---------HR
	//     │
	//     │
	//     |
	//     T
	//
	var vec *Vector
	{
		vec = New(Config{
			Obj: []object.Object{
				{X: 100, Y: 100}, // T
				{X: 100, Y: 105},
				{X: 100, Y: 110},
				{X: 100, Y: 115},
				{X: 100, Y: 120},
				{X: 105, Y: 120},
				{X: 110, Y: 120},
				{X: 115, Y: 120},
				{X: 120, Y: 120}, // H
			},
		})
	}

	cbl := object.Object{X: 100, Y: 100}
	ctr := object.Object{X: 120, Y: 120}

	if !reflect.DeepEqual(vec.Window().CBL(), cbl) {
		t.Fatalf("expected %#v got %#v", cbl, vec.Window().CBL())
	}
	if !reflect.DeepEqual(vec.Window().CTR(), ctr) {
		t.Fatalf("expected %#v got %#v", ctr, vec.Window().CTR())
	}

	cbl = object.Object{X: 100, Y: 105} // new T
	ctr = object.Object{X: 125, Y: 120} // R

	{
		vec.Rotate(ctr) // R
	}

	if !reflect.DeepEqual(vec.Window().CBL(), cbl) {
		t.Fatalf("expected %#v got %#v", cbl, vec.Window().CBL())
	}
	if !reflect.DeepEqual(vec.Window().CTR(), ctr) {
		t.Fatalf("expected %#v got %#v", ctr, vec.Window().CTR())
	}
}

func Test_Vector_Window_No_Change(t *testing.T) {
	//
	//     +---------+
	//     │    R    │
	//     │    H    │
	//     |    │    │
	//     T    +----+
	//
	var vec *Vector
	{
		vec = New(Config{
			Obj: []object.Object{
				{X: 100, Y: 100}, // T
				{X: 100, Y: 105},
				{X: 100, Y: 110},
				{X: 100, Y: 115},
				{X: 100, Y: 120},
				{X: 105, Y: 120},
				{X: 110, Y: 120},
				{X: 115, Y: 120},
				{X: 120, Y: 120},
				{X: 120, Y: 115},
				{X: 120, Y: 110},
				{X: 120, Y: 105},
				{X: 120, Y: 100},
				{X: 115, Y: 100},
				{X: 110, Y: 100},
				{X: 110, Y: 105},
				{X: 110, Y: 110}, // H
			},
		})
	}

	cbl := object.Object{X: 100, Y: 100}
	ctr := object.Object{X: 120, Y: 120}

	if !reflect.DeepEqual(vec.Window().CBL(), cbl) {
		t.Fatalf("expected %#v got %#v", cbl, vec.Window().CBL())
	}
	if !reflect.DeepEqual(vec.Window().CTR(), ctr) {
		t.Fatalf("expected %#v got %#v", ctr, vec.Window().CTR())
	}

	{
		vec.Rotate(object.Object{X: 110, Y: 115}) // R
	}

	if !reflect.DeepEqual(vec.Window().CBL(), cbl) {
		t.Fatalf("expected %#v got %#v", cbl, vec.Window().CBL())
	}
	if !reflect.DeepEqual(vec.Window().CTR(), ctr) {
		t.Fatalf("expected %#v got %#v", ctr, vec.Window().CTR())
	}
}
