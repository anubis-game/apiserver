package vector

import (
	"fmt"
	"slices"
	"testing"
)

func Test_Vector_Rotate(t *testing.T) {
	testCases := []struct {
		obj []Object
		trg Object
		rot []Object
	}{
		// Case 000, x+3 y+4
		{
			obj: []Object{
				{621_362, 539_077},
				{621_359, 539_073},
			},
			trg: Object{621_365, 539_081},
			rot: []Object{
				{621_365, 539_081},
				{621_362, 539_077},
			},
		},
		// Case 001, y-5
		{
			obj: []Object{
				{621_359, 539_068},
				{621_359, 539_073},
			},
			trg: Object{621_359, 539_063},
			rot: []Object{
				{621_359, 539_063},
				{621_359, 539_068},
			},
		},
		// Case 002, x-5
		{
			obj: []Object{
				{621_354, 539_073},
				{621_349, 539_073},
			},
			trg: Object{621_359, 539_073},
			rot: []Object{
				{621_359, 539_073},
				{621_354, 539_073},
			},
		},
		// Case 003, x-4 y+3
		{
			obj: []Object{
				{621_355, 539_076},
				{621_351, 539_079},
			},
			trg: Object{621_359, 539_073},
			rot: []Object{
				{621_359, 539_073},
				{621_355, 539_076},
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

			vec.Rotate(tc.trg)

			if !slices.Equal(vec.obj[:vec.len], tc.rot) {
				t.Fatalf("expected %#v got %#v", tc.rot, vec.obj[:vec.len])
			}
		})
	}
}

func Benchmark_Vector_Rotate(b *testing.B) {
	testCases := []struct {
		obj []Object
		trg Object
	}{
		// Case 000, 1.90 ns/op
		{
			obj: []Object{
				{621_362, 539_077},
				{621_359, 539_073},
			},
			trg: Object{621_365, 539_081},
		},
		// Case 001, 1.90 ns/op
		{
			obj: []Object{
				{621_359, 539_068},
				{621_359, 539_073},
			},
			trg: Object{621_359, 539_063},
		},
		// Case 002, 1.90 ns/op
		{
			obj: []Object{
				{621_354, 539_073},
				{621_349, 539_073},
			},
			trg: Object{621_359, 539_073},
		},
		// Case 003, 1.90 ns/op
		{
			obj: []Object{
				{621_355, 539_076},
				{621_351, 539_079},
			},
			trg: Object{621_359, 539_073},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			var vec *Vector
			{
				vec = New(Config{
					Obj: tc.obj,
				})
			}

			b.ResetTimer()
			for range b.N {
				vec.Rotate(tc.trg)
			}
		})
	}
}
