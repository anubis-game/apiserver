package vector

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_target(t *testing.T) {
	testCases := []struct {
		hea matrix.Coordinate
		qdr byte
		agl byte
		dis float64
		trg matrix.Coordinate
	}{
		// Case 000
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(1),                                   // quadrant 1
			agl: byte(108),                                 // 38.12° from 0°
			dis: nrm,                                       // normal speed
			trg: matrix.Coordinate{X: 621_362, Y: 539_077}, // x+3 y+4
		},
		// Case 001
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(2),                                   // quadrant 2
			agl: byte(253),                                 // 89.29° from 90°
			dis: nrm * 4,                                   // racing speed
			trg: matrix.Coordinate{X: 621_359, Y: 539_053}, // y-18
		},
		// Case 002
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(3),                                   // quadrant 3
			agl: byte(253),                                 // 89.29° from 180°
			dis: nrm * 4,                                   // racing speed
			trg: matrix.Coordinate{X: 621_339, Y: 539_073}, // x-18
		},
		// Case 003
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(4),                                   // quadrant 4
			agl: byte(108),                                 // 38.12° from 180°
			dis: nrm,                                       // normal speed
			trg: matrix.Coordinate{X: 621_355, Y: 539_076}, // x-4 y+3
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = New(Config{
					Hea: tc.hea,
				})
			}

			var trg matrix.Coordinate
			{
				trg = vec.target(tc.qdr, tc.agl, tc.dis)
			}

			if trg != tc.trg {
				t.Fatalf("expected %#v got %#v", tc.trg, trg)
			}
		})
	}
}

func Benchmark_Vector_target(b *testing.B) {
	testCases := []struct {
		hea matrix.Coordinate
		qdr byte
		agl byte
		dis float64
	}{
		// Case 000, ~53 ns/op
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(1),   // quadrant 1
			agl: byte(108), // 38.12° from 0°
			dis: nrm,       // normal speed
		},
		// Case 001, ~53 ns/op
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(2),   // quadrant 2
			agl: byte(253), // 89.29° from 90°
			dis: nrm * 4,   // racing speed
		},
		// Case 002, ~53 ns/op
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(3),   // quadrant 3
			agl: byte(253), // 89.29° from 180°
			dis: nrm * 4,   // racing speed
		},
		// Case 003, ~53 ns/op
		{
			hea: matrix.Coordinate{X: 621_359, Y: 539_073},
			qdr: byte(4),   // quadrant 4
			agl: byte(108), // 38.12° from 180°
			dis: nrm,       // normal speed
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			var vec *Vector
			{
				vec = New(Config{
					Hea: tc.hea,
				})
			}

			for b.Loop() {
				vec.target(tc.qdr, tc.agl, tc.dis)
			}
		})
	}
}
