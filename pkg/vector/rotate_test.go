package vector

import (
	"fmt"
	"slices"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Rotate(t *testing.T) {
	testCases := []struct {
		obj []object.Object
		trg object.Object
		rot []object.Object
	}{
		// Case 000, x+3 y+4
		{
			obj: []object.Object{
				{X: 621_362, Y: 539_077},
				{X: 621_359, Y: 539_073},
			},
			trg: object.Object{X: 621_365, Y: 539_081},
			rot: []object.Object{
				{X: 621_365, Y: 539_081},
				{X: 621_362, Y: 539_077},
			},
		},
		// Case 001, y-5
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_068},
				{X: 621_359, Y: 539_073},
			},
			trg: object.Object{X: 621_359, Y: 539_063},
			rot: []object.Object{
				{X: 621_359, Y: 539_063},
				{X: 621_359, Y: 539_068},
			},
		},
		// Case 002, x-5
		{
			obj: []object.Object{
				{X: 621_354, Y: 539_073},
				{X: 621_349, Y: 539_073},
			},
			trg: object.Object{X: 621_359, Y: 539_073},
			rot: []object.Object{
				{X: 621_359, Y: 539_073},
				{X: 621_354, Y: 539_073},
			},
		},
		// Case 003, x-4 y+3
		{
			obj: []object.Object{
				{X: 621_355, Y: 539_076},
				{X: 621_351, Y: 539_079},
			},
			trg: object.Object{X: 621_359, Y: 539_073},
			rot: []object.Object{
				{X: 621_359, Y: 539_073},
				{X: 621_355, Y: 539_076},
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
		obj []object.Object
		trg object.Object
	}{
		// Case 000, 1.90 ns/op
		{
			obj: []object.Object{
				{X: 621_362, Y: 539_077},
				{X: 621_359, Y: 539_073},
			},
			trg: object.Object{X: 621_365, Y: 539_081},
		},
		// Case 001, 1.90 ns/op
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_068},
				{X: 621_359, Y: 539_073},
			},
			trg: object.Object{X: 621_359, Y: 539_063},
		},
		// Case 002, 1.90 ns/op
		{
			obj: []object.Object{
				{X: 621_354, Y: 539_073},
				{X: 621_349, Y: 539_073},
			},
			trg: object.Object{X: 621_359, Y: 539_073},
		},
		// Case 003, 1.90 ns/op
		{
			obj: []object.Object{
				{X: 621_355, Y: 539_076},
				{X: 621_351, Y: 539_079},
			},
			trg: object.Object{X: 621_359, Y: 539_073},
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
