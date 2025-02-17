package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Expand(t *testing.T) {
	testCases := []struct {
		obj []object.Object
		hea object.Object
		exp []object.Object
	}{
		// Case 000, x+3 y+4
		{
			obj: []object.Object{
				{X: 621_362, Y: 539_077},
				{X: 621_359, Y: 539_073},
			},
			hea: object.Object{X: 621_365, Y: 539_081},
			exp: []object.Object{
				{X: 621_365, Y: 539_081},
				{X: 621_362, Y: 539_077},
				{X: 621_359, Y: 539_073},
			},
		},
		// Case 001, y-5
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_068},
				{X: 621_359, Y: 539_073},
			},
			hea: object.Object{X: 621_359, Y: 539_063},
			exp: []object.Object{
				{X: 621_359, Y: 539_063},
				{X: 621_359, Y: 539_068},
				{X: 621_359, Y: 539_073},
			},
		},
		// Case 002, x-5
		{
			obj: []object.Object{
				{X: 621_354, Y: 539_073},
				{X: 621_349, Y: 539_073},
			},
			hea: object.Object{X: 621_359, Y: 539_073},
			exp: []object.Object{
				{X: 621_359, Y: 539_073},
				{X: 621_354, Y: 539_073},
				{X: 621_349, Y: 539_073},
			},
		},
		// Case 003, x-4 y+3
		{
			obj: []object.Object{
				{X: 621_355, Y: 539_076},
				{X: 621_351, Y: 539_079},
			},
			hea: object.Object{X: 621_359, Y: 539_073},
			exp: []object.Object{
				{X: 621_359, Y: 539_073},
				{X: 621_355, Y: 539_076},
				{X: 621_351, Y: 539_079},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = New(Config{
					Obj: tc.obj,
				})
			}

			if vec.len != len(tc.obj) {
				t.Fatalf("expected %#v got %#v", len(tc.obj), vec.len)
			}

			vec.Expand(tc.hea)

			if vec.len != len(tc.exp) {
				t.Fatalf("expected %#v got %#v", len(tc.exp), vec.len)
			}
			if !reflect.DeepEqual(vec.hea.val, tc.hea) {
				t.Fatalf("expected %#v got %#v", tc.hea, vec.hea.val)
			}
		})
	}
}
