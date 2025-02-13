package vector

import (
	"fmt"
	"testing"
)

func Test_Vector_Adjust_angle(t *testing.T) {
	testCases := []struct {
		siz float64
		agl byte
	}{
		// Case 000
		{
			siz: 10,
			agl: 30,
		},
		// Case 001
		{
			siz: 50,
			agl: 29,
		},
		// Case 002
		{
			siz: 100,
			agl: 28,
		},
		// Case 003
		{
			siz: 250,
			agl: 27,
		},
		// Case 004
		{
			siz: 500,
			agl: 26,
		},
		// Case 005
		{
			siz: 1_000,
			agl: 24,
		},
		// Case 006
		{
			siz: 2_500,
			agl: 21,
		},
		// Case 007
		{
			siz: 5_000,
			agl: 19,
		},
		// Case 008
		{
			siz: 10_000,
			agl: 15,
		},
		// Case 009
		{
			siz: 20_000,
			agl: 10,
		},
		// Case 010
		{
			siz: 30_000,
			agl: 6,
		},
		// Case 011
		{
			siz: 40_000,
			agl: 3,
		},
		// Case 012
		{
			siz: 50_000,
			agl: 1,
		},
		// Case 013
		{
			siz: 60_000,
			agl: 1,
		},
		// Case 014
		{
			siz: 70_000,
			agl: 1,
		},
		// Case 015
		{
			siz: 80_000,
			agl: 1,
		},
		// Case 016
		{
			siz: 90_000,
			agl: 1,
		},
		// Case 017
		{
			siz: 100_000,
			agl: 1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			agl := angle(tc.siz)

			if agl != tc.agl {
				t.Fatalf("expected %d got %d", tc.agl, agl)
			}
		})
	}
}

func Test_Vector_Adjust_length(t *testing.T) {
	testCases := []struct {
		siz float64
		len int
	}{
		// Case 000
		{
			siz: 10,
			len: 10,
		},
		// Case 001
		{
			siz: 50,
			len: 10,
		},
		// Case 002
		{
			siz: 100,
			len: 19,
		},
		// Case 003
		{
			siz: 250,
			len: 43,
		},
		// Case 004
		{
			siz: 500,
			len: 80,
		},
		// Case 005
		{
			siz: 1_000,
			len: 149,
		},
		// Case 006
		{
			siz: 2_500,
			len: 339,
		},
		// Case 007
		{
			siz: 5_000,
			len: 631,
		},
		// Case 008
		{
			siz: 10_000,
			len: 1_178,
		},
		// Case 009
		{
			siz: 20_000,
			len: 2_198,
		},
		// Case 010
		{
			siz: 30_000,
			len: 3_165,
		},
		// Case 011
		{
			siz: 40_000,
			len: 4_100,
		},
		// Case 012
		{
			siz: 50_000,
			len: 5_000,
		},
		// Case 013
		{
			siz: 60_000,
			len: 5_000,
		},
		// Case 014
		{
			siz: 70_000,
			len: 5_000,
		},
		// Case 015
		{
			siz: 80_000,
			len: 5_000,
		},
		// Case 016
		{
			siz: 90_000,
			len: 5_000,
		},
		// Case 017
		{
			siz: 100_000,
			len: 5_000,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			len := length(tc.siz)

			if len != tc.len {
				t.Fatalf("expected %#v got %#v", tc.len, len)
			}
		})
	}
}

func Test_Vector_Adjust_radius(t *testing.T) {
	testCases := []struct {
		siz float64
		rad int
	}{
		// Case 000
		{
			siz: 10,
			rad: 10,
		},
		// Case 001
		{
			siz: 50,
			rad: 10,
		},
		// Case 002
		{
			siz: 100,
			rad: 14,
		},
		// Case 003
		{
			siz: 250,
			rad: 22,
		},
		// Case 004
		{
			siz: 500,
			rad: 30,
		},
		// Case 005
		{
			siz: 1_000,
			rad: 41,
		},
		// Case 006
		{
			siz: 2_500,
			rad: 63,
		},
		// Case 007
		{
			siz: 5_000,
			rad: 88,
		},
		// Case 008
		{
			siz: 10_000,
			rad: 121,
		},
		// Case 009
		{
			siz: 20_000,
			rad: 168,
		},
		// Case 010
		{
			siz: 30_000,
			rad: 203,
		},
		// Case 011
		{
			siz: 40_000,
			rad: 232,
		},
		// Case 012
		{
			siz: 50_000,
			rad: 256,
		},
		// Case 013
		{
			siz: 60_000,
			rad: 256,
		},
		// Case 014
		{
			siz: 70_000,
			rad: 256,
		},
		// Case 015
		{
			siz: 80_000,
			rad: 256,
		},
		// Case 016
		{
			siz: 90_000,
			rad: 256,
		},
		// Case 017
		{
			siz: 100_000,
			rad: 256,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			rad := radius(tc.siz)

			if rad != tc.rad {
				t.Fatalf("expected %#v got %#v", tc.rad, rad)
			}
		})
	}
}

func Test_Vector_Adjust_sight(t *testing.T) {
	testCases := []struct {
		siz float64
		prt int
	}{
		// Case 000
		{
			siz: 10,
			prt: 2,
		},
		// Case 001
		{
			siz: 50,
			prt: 2,
		},
		// Case 002
		{
			siz: 100,
			prt: 2,
		},
		// Case 003
		{
			siz: 250,
			prt: 3,
		},
		// Case 004
		{
			siz: 500,
			prt: 3,
		},
		// Case 005
		{
			siz: 1_000,
			prt: 4,
		},
		// Case 006
		{
			siz: 2_500,
			prt: 4,
		},
		// Case 007
		{
			siz: 5_000,
			prt: 5,
		},
		// Case 008
		{
			siz: 10_000,
			prt: 6,
		},
		// Case 009
		{
			siz: 20_000,
			prt: 7,
		},
		// Case 010
		{
			siz: 30_000,
			prt: 7,
		},
		// Case 011
		{
			siz: 40_000,
			prt: 8,
		},
		// Case 012
		{
			siz: 50_000,
			prt: 8,
		},
		// Case 013
		{
			siz: 60_000,
			prt: 8,
		},
		// Case 014
		{
			siz: 70_000,
			prt: 8,
		},
		// Case 015
		{
			siz: 80_000,
			prt: 8,
		},
		// Case 016
		{
			siz: 90_000,
			prt: 8,
		},
		// Case 017
		{
			siz: 100_000,
			prt: 8,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			prt := sight(tc.siz)

			if prt != tc.prt {
				t.Fatalf("expected %#v got %#v", tc.prt, prt)
			}
		})
	}
}

func Test_Vector_Adjust_trgAgl(t *testing.T) {
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

func Benchmark_Vector_Adjust_trgAgl(b *testing.B) {
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
