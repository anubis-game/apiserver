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
		tai object.Object
		exp []object.Object
	}{
		// Case 000, x+3 y+4
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
				{X: 621_362, Y: 539_077},
			},
			hea: object.Object{X: 621_365, Y: 539_081},
			tai: object.Object{X: 621_359, Y: 539_073},
		},
		// Case 001, y-5
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
				{X: 621_359, Y: 539_068},
			},
			hea: object.Object{X: 621_359, Y: 539_063},
			tai: object.Object{X: 621_359, Y: 539_073},
		},
		// Case 002, x-5
		{
			obj: []object.Object{
				{X: 621_349, Y: 539_073},
				{X: 621_354, Y: 539_073},
			},
			hea: object.Object{X: 621_359, Y: 539_073},
			tai: object.Object{X: 621_349, Y: 539_073},
		},
		// Case 003, x-4 y+3
		{
			obj: []object.Object{
				{X: 621_351, Y: 539_079},
				{X: 621_355, Y: 539_076},
			},
			hea: object.Object{X: 621_359, Y: 539_073},
			tai: object.Object{X: 621_351, Y: 539_079},
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

			l := vec.lis.Len()

			vec.Expand(tc.hea)

			if vec.lis.Len() != l+1 {
				t.Fatalf("expected %#v got %#v", l+1, vec.lis.Len())
			}
			if !reflect.DeepEqual(vec.lis.Front().Value.(object.Object), tc.hea) {
				t.Fatalf("expected %#v got %#v", tc.hea, vec.lis.Front().Value.(object.Object))
			}
			if !reflect.DeepEqual(vec.lis.Back().Value.(object.Object), tc.tai) {
				t.Fatalf("expected %#v got %#v", tc.tai, vec.lis.Back().Value.(object.Object))
			}
		})
	}
}
