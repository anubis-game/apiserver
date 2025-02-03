package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Vector_Target(t *testing.T) {
	testCases := []struct {
		vec *Vector
		mot Motion
		trg Object
	}{
		// Case 000
		{
			vec: &Vector{
				obj: []Object{
					{621_359, 539_073},
				},
			},
			mot: Motion{
				QDR: byte(1),   // quadrant 1
				AGL: byte(108), // 38.12° from 0°
				VLC: byte(1),   // 100% speed
			},
			trg: Object{621_362, 539_077}, // x+3 y+4
		},
		// Case 001
		{
			vec: &Vector{
				obj: []Object{
					{621_359, 539_073},
				},
			},
			mot: Motion{
				QDR: byte(2),   // quadrant 2
				AGL: byte(253), // 89.29° from 90°
				VLC: byte(4),   // 400% speed
			},
			trg: Object{621_359, 539_053}, // y-18
		},
		// Case 002
		{
			vec: &Vector{
				obj: []Object{
					{621_359, 539_073},
				},
			},
			mot: Motion{
				QDR: byte(3),   // quadrant 3
				AGL: byte(253), // 89.29° from 180°
				VLC: byte(4),   // 400% speed
			},
			trg: Object{621_339, 539_073}, // x-18
		},
		// Case 003
		{
			vec: &Vector{
				obj: []Object{
					{621_359, 539_073},
				},
			},
			mot: Motion{
				QDR: byte(4),   // quadrant 4
				AGL: byte(108), // 38.12° from 180°
				VLC: byte(1),   // 100% speed
			},
			trg: Object{621_355, 539_076}, // x-4 y+3
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			for i := 0; i < int(tc.mot.VLC); i++ {
				tc.vec.obj[HD] = tc.vec.Target(tc.mot)
			}

			if !reflect.DeepEqual(tc.vec.obj[HD], tc.trg) {
				t.Fatalf("expected %#v got %#v", tc.trg, tc.vec.obj[HD])
			}
		})
	}
}

func Benchmark_Vector_Target(b *testing.B) {
	testCases := []struct {
		vec *Vector
		mot Motion
	}{
		// Case 000, ~2.00 ns/op
		{
			vec: &Vector{
				obj: []Object{
					{621_359, 539_073},
				},
			},
			mot: Motion{
				QDR: byte(1),   // quadrant 1
				AGL: byte(108), // 38.12° from 0°
				VLC: byte(1),   // 100% speed
			},
		},
		// Case 001, ~2.00 ns/op
		{
			vec: &Vector{
				obj: []Object{
					{621_359, 539_073},
				},
			},
			mot: Motion{
				QDR: byte(2),   // quadrant 2
				AGL: byte(253), // 89.29° from 90°
				VLC: byte(4),   // 400% speed
			},
		},
		// Case 002, ~2.00 ns/op
		{
			vec: &Vector{
				obj: []Object{
					{621359, 539073},
				},
			},
			mot: Motion{
				QDR: byte(3),   // quadrant 3
				AGL: byte(253), // 89.29° from 180°
				VLC: byte(4),   // 400% speed
			},
		},
		// Case 003, ~2.00 ns/op
		{
			vec: &Vector{
				obj: []Object{
					{621_359, 539_073},
				},
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
			b.ResetTimer()
			for range b.N {
				tc.vec.Target(tc.mot)
			}
		})
	}
}
