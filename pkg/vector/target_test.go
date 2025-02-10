package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Target(t *testing.T) {
	testCases := []struct {
		vec *Vector
		mot Motion
		hea object.Object
	}{
		// Case 000
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(1),
					Agl: byte(8),
					Vlc: byte(1),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(1),   // quadrant 1
				Agl: byte(108), // 38.12° from 0°
				Vlc: byte(1),   // 100% speed
			},
			hea: object.Object{X: 621_362, Y: 539_077}, // x+3 y+4
		},
		// Case 001
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(2),
					Agl: byte(249),
					Vlc: byte(1),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(2),   // quadrant 2
				Agl: byte(253), // 89.29° from 90°
				Vlc: byte(4),   // 400% speed
			},
			hea: object.Object{X: 621_359, Y: 539_053}, // y-18
		},
		// Case 002
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(3),
					Agl: byte(222),
					Vlc: byte(4),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(3),   // quadrant 3
				Agl: byte(253), // 89.29° from 180°
				Vlc: byte(4),   // 400% speed
			},
			hea: object.Object{X: 621_339, Y: 539_073}, // x-18
		},
		// Case 003
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(4),
					Agl: byte(199),
					Vlc: byte(4),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(4),   // quadrant 4
				Agl: byte(108), // 38.12° from 180°
				Vlc: byte(1),   // 100% speed
			},
			hea: object.Object{X: 621_355, Y: 539_076}, // x-4 y+3
		},
		// Case 004, the same as 003, but with an invalid velocity, the resulting
		// head coordinate must be the same.
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(4),
					Agl: byte(199),
					Vlc: byte(1),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(4),   // quadrant 4
				Agl: byte(108), // 38.12° from 180°
				Vlc: byte(2),   // invalid speed
			},
			hea: object.Object{X: 621_355, Y: 539_076}, // x-4 y+3
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var hea object.Object
			{
				hea = tc.vec.Target(tc.mot)
			}

			if !reflect.DeepEqual(hea, tc.hea) {
				t.Fatalf("expected %#v got %#v", tc.hea, hea)
			}
		})
	}
}

func Test_Vector_trgAgl(t *testing.T) {
	testCases := []struct {
		pqd byte
		pag byte
		nqd byte
		nag byte
		lim byte
		qdr byte
		agl byte
	}{
		// Case 000, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to remain inside the 1st
		// quadrant, towards the 2nd quadrant.
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(2),
			nag: byte(87),
			lim: byte(100),
			qdr: byte(1),
			agl: byte(207),
		},
		// Case 001, the desired range of motion is granted, because the maximum
		// angle deviation of 100 bytes is respected, which then does not force the
		// allowed range of motion to be restricted.
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(1),
			nag: byte(87),
			lim: byte(100),
			qdr: byte(1),
			agl: byte(87),
		},
		// Case 002, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to remain inside the 1st
		// quadrant, towards the 2nd quadrant.
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(4), // under 180°, move clockwise
			lim: byte(100),
			qdr: byte(1),
			agl: byte(105),
		},
		// Case 003, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to remain inside the 1st
		// quadrant, towards the 2nd quadrant.
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(5), // exactly 180°, move clockwise
			lim: byte(100),
			qdr: byte(1),
			agl: byte(105),
		},
		// Case 004, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to underflow into the 4th
		// quadrant.
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(6), // over 180°, move counter clockwise
			lim: byte(100),
			qdr: byte(4),
			agl: byte(161),
		},
		// Case 005
		{
			pqd: byte(4),
			pag: byte(161),
			nqd: byte(3),
			nag: byte(1), // under 180°, move counter clockwise
			lim: byte(100),
			qdr: byte(4),
			agl: byte(61),
		},
		// Case 006
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(77), // under 180°, move clockwise
			lim: byte(75),
			qdr: byte(2),
			agl: byte(163),
		},
		// Case 007
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(99), // above 180°, move counter clockwise
			lim: byte(75),
			qdr: byte(2),
			agl: byte(13),
		},
		// Case 008
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(90), // under 180°, move clockwise
			lim: byte(175),
			qdr: byte(2),
			agl: byte(90),
		},
		// Case 009
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(85), // under 180°, move clockwise
			lim: byte(175),
			qdr: byte(2),
			agl: byte(85),
		},
		// Case 010
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(1),
			nag: byte(2), // under 180°, move counter clockwise
			lim: byte(175),
			qdr: byte(1),
			agl: byte(169),
		},
		// Case 011
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(88), // exactly 180°, move clockwise
			lim: byte(175),
			qdr: byte(3),
			agl: byte(7),
		},
		// Case 012
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(84), // under 180°, move clockwise
			lim: byte(175),
			qdr: byte(3),
			agl: byte(7),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var qdr byte
			var agl byte
			{
				qdr, agl = trgAgl(tc.pqd, tc.pag, tc.nqd, tc.nag, tc.lim)
			}

			if qdr != tc.qdr {
				t.Fatalf("expected %#v got %#v", tc.qdr, qdr)
			}
			if agl != tc.agl {
				t.Fatalf("expected %#v got %#v", tc.agl, agl)
			}
		})
	}
}

func Benchmark_Vector_Target(b *testing.B) {
	testCases := []struct {
		vec *Vector
		mot Motion
	}{
		// Case 000, ~5.00 ns/op
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(1),
					Agl: byte(8),
					Vlc: byte(1),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(1),   // quadrant 1
				Agl: byte(108), // 38.12° from 0°
				Vlc: byte(1),   // 100% speed
			},
		},
		// Case 001, ~4.70 ns/op
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(2),
					Agl: byte(249),
					Vlc: byte(1),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(2),   // quadrant 2
				Agl: byte(253), // 89.29° from 90°
				Vlc: byte(4),   // 400% speed
			},
		},
		// Case 002, ~4.80 ns/op
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(3),
					Agl: byte(222),
					Vlc: byte(4),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(3),   // quadrant 3
				Agl: byte(253), // 89.29° from 180°
				Vlc: byte(4),   // 400% speed
			},
		},
		// Case 003, ~4.70 ns/op
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(4),
					Agl: byte(199),
					Vlc: byte(4),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(4),   // quadrant 4
				Agl: byte(108), // 38.12° from 180°
				Vlc: byte(1),   // 100% speed
			},
		},
		// Case 004, ~4.80 ns/op
		{
			vec: New(Config{
				Mot: Motion{
					Qdr: byte(4),
					Agl: byte(199),
					Vlc: byte(1),
				},
				Obj: []object.Object{
					{X: 621_359, Y: 539_073},
				},
			}),
			mot: Motion{
				Qdr: byte(4),   // quadrant 4
				Agl: byte(108), // 38.12° from 180°
				Vlc: byte(2),   // invalid speed
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

func Benchmark_Vector_trgAgl(b *testing.B) {
	testCases := []struct {
		pqd byte
		pag byte
		nqd byte
		nag byte
		lim byte
	}{
		// Case 000, ~1.80 ns/op
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(2),
			nag: byte(87),
			lim: byte(100),
		},
		// Case 001, ~1.80 ns/op
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(1),
			nag: byte(87),
			lim: byte(100),
		},
		// Case 002, ~1.80 ns/op
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(4), // under 180°, move clockwise
			lim: byte(100),
		},
		// Case 003, ~1.80 ns/op
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(5), // exactly 180°, move clockwise
			lim: byte(100),
		},
		// Case 004, ~1.80 ns/op
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(6), // over 180°, move counter clockwise
			lim: byte(100),
		},
		// Case 005, ~1.70 ns/op
		{
			pqd: byte(4),
			pag: byte(161),
			nqd: byte(3),
			nag: byte(1), // under 180°, move counter clockwise
			lim: byte(100),
		},
		// Case 006, ~1.80 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(77), // under 180°, move clockwise
			lim: byte(75),
		},
		// Case 007, ~1.80 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(99), // above 180°, move counter clockwise
			lim: byte(75),
		},
		// Case 008, ~1.80 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(90), // under 180°, move clockwise
			lim: byte(175),
		},
		// Case 009, ~1.80 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(85), // under 180°, move clockwise
			lim: byte(175),
		},
		// Case 010, ~1.80 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(1),
			nag: byte(2), // under 180°, move counter clockwise
			lim: byte(175),
		},
		// Case 011, ~1.80 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(88), // exactly 180°, move clockwise
			lim: byte(175),
		},
		// Case 012, ~1.80 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(84), // under 180°, move clockwise
			lim: byte(175),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for range b.N {
				trgAgl(tc.pqd, tc.pag, tc.nqd, tc.nag, tc.lim)
			}
		})
	}
}
