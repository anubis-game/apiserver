package vector

import (
	"fmt"
	"slices"
	"testing"
)

func Test_Vector_Expand(t *testing.T) {
	testCases := []struct {
		obj []Object
		trg Object
		exp []Object
	}{
		// Case 000, x+3 y+4
		{
			obj: []Object{
				{621_362, 539_077},
				{621_359, 539_073},
			},
			trg: Object{621_365, 539_081},
			exp: []Object{
				{621_365, 539_081},
				{621_362, 539_077},
				{621_359, 539_073},
			},
		},
		// Case 001, y-5
		{
			obj: []Object{
				{621_359, 539_068},
				{621_359, 539_073},
			},
			trg: Object{621_359, 539_063},
			exp: []Object{
				{621_359, 539_063},
				{621_359, 539_068},
				{621_359, 539_073},
			},
		},
		// Case 002, x-5
		{
			obj: []Object{
				{621_354, 539_073},
				{621_349, 539_073},
			},
			trg: Object{621_359, 539_073},
			exp: []Object{
				{621_359, 539_073},
				{621_354, 539_073},
				{621_349, 539_073},
			},
		},
		// Case 003, x-4 y+3
		{
			obj: []Object{
				{621_355, 539_076},
				{621_351, 539_079},
			},
			trg: Object{621_359, 539_073},
			exp: []Object{
				{621_359, 539_073},
				{621_355, 539_076},
				{621_351, 539_079},
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

			vec.Expand(tc.trg)

			if !slices.Equal(vec.obj[:vec.len], tc.exp) {
				t.Fatalf("expected %#v got %#v", tc.exp, vec.obj[:vec.len])
			}
		})
	}
}
