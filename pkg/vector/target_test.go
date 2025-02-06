package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Target(t *testing.T) {
	testCases := []struct {
		obj []object.Object
		mot Motion
		hea object.Object
	}{
		// Case 000
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(1),   // quadrant 1
				AGL: byte(108), // 38.12° from 0°
				VLC: byte(1),   // 100% speed
			},
			hea: object.Object{X: 621_362, Y: 539_077}, // x+3 y+4
		},
		// Case 001
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(2),   // quadrant 2
				AGL: byte(253), // 89.29° from 90°
				VLC: byte(4),   // 400% speed
			},
			hea: object.Object{X: 621_359, Y: 539_053}, // y-18
		},
		// Case 002
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(3),   // quadrant 3
				AGL: byte(253), // 89.29° from 180°
				VLC: byte(4),   // 400% speed
			},
			hea: object.Object{X: 621_339, Y: 539_073}, // x-18
		},
		// Case 003
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(4),   // quadrant 4
				AGL: byte(108), // 38.12° from 180°
				VLC: byte(1),   // 100% speed
			},
			hea: object.Object{X: 621_355, Y: 539_076}, // x-4 y+3
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

			for i := 0; i < int(tc.mot.VLC); i++ {
				trg := vec.Target(tc.mot)
				vec.lis.Remove(vec.lis.Front())
				vec.lis.PushFront(trg)
			}

			if !reflect.DeepEqual(vec.lis.Front().Value.(object.Object), tc.hea) {
				t.Fatalf("expected %#v got %#v", tc.hea, vec.lis.Front().Value.(object.Object))
			}
		})
	}
}

func Benchmark_Vector_Target(b *testing.B) {
	testCases := []struct {
		obj []object.Object
		mot Motion
	}{
		// Case 000, ~1.40 ns/op
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(1),   // quadrant 1
				AGL: byte(108), // 38.12° from 0°
				VLC: byte(1),   // 100% speed
			},
		},
		// Case 001, ~1.70 ns/op
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(2),   // quadrant 2
				AGL: byte(253), // 89.29° from 90°
				VLC: byte(4),   // 400% speed
			},
		},
		// Case 002, ~1.70 ns/op
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(3),   // quadrant 3
				AGL: byte(253), // 89.29° from 180°
				VLC: byte(4),   // 400% speed
			},
		},
		// Case 003, ~1.70 ns/op
		{
			obj: []object.Object{
				{X: 621_359, Y: 539_073},
			},
			mot: Motion{
				QDR: byte(4),   // quadrant 4
				AGL: byte(108), // 38.12° from 180°
				VLC: byte(1),   // 100% speed
			},
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
				vec.Target(tc.mot)
			}
		})
	}
}
