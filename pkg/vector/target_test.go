package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Target(t *testing.T) {
	testCases := []struct {
		obj object.Object
		qdr byte
		agl byte
		dis float64
		hea object.Object
	}{
		// Case 000
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(1),                               // quadrant 1
			agl: byte(108),                             // 38.12° from 0°
			dis: Dis,                                   // normal speed
			hea: object.Object{X: 621_362, Y: 539_077}, // x+3 y+4
		},
		// Case 001
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(2),                               // quadrant 2
			agl: byte(253),                             // 89.29° from 90°
			dis: Ris,                                   // racing speed
			hea: object.Object{X: 621_359, Y: 539_053}, // y-18
		},
		// Case 002
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(3),                               // quadrant 3
			agl: byte(253),                             // 89.29° from 180°
			dis: Ris,                                   // racing speed
			hea: object.Object{X: 621_339, Y: 539_073}, // x-18
		},
		// Case 003
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(4),                               // quadrant 4
			agl: byte(108),                             // 38.12° from 180°
			dis: Dis,                                   // normal speed
			hea: object.Object{X: 621_355, Y: 539_076}, // x-4 y+3
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = New(Config{
					Obj: []object.Object{
						tc.obj,
					},
				})
			}

			var hea object.Object
			{
				hea = vec.Target(tc.qdr, tc.agl, tc.dis)
			}

			if !reflect.DeepEqual(hea, tc.hea) {
				t.Fatalf("expected %#v got %#v", tc.hea, hea)
			}
		})
	}
}

func Benchmark_Vector_Target(b *testing.B) {
	testCases := []struct {
		obj object.Object
		qdr byte
		agl byte
		dis float64
	}{
		// Case 000, ~1.90 ns/op
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(1),   // quadrant 1
			agl: byte(108), // 38.12° from 0°
			dis: Dis,       // normal speed
		},
		// Case 001, ~1.90 ns/op
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(2),   // quadrant 2
			agl: byte(253), // 89.29° from 90°
			dis: Ris,       // racing speed
		},
		// Case 002, ~1.90 ns/op
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(3),   // quadrant 3
			agl: byte(253), // 89.29° from 180°
			dis: Ris,       // racing speed
		},
		// Case 003, ~1.90 ns/op
		{
			obj: object.Object{X: 621_359, Y: 539_073},
			qdr: byte(4),   // quadrant 4
			agl: byte(108), // 38.12° from 180°
			dis: Dis,       // normal speed
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			var vec *Vector
			{
				vec = New(Config{
					Obj: []object.Object{
						tc.obj,
					},
				})
			}

			for b.Loop() {
				vec.Target(tc.qdr, tc.agl, tc.dis)
			}
		})
	}
}
