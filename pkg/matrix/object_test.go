package matrix

import (
	"fmt"
	"testing"
)

func Test_Matrix_Intersect(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		osz byte
		ebc Bucket
		epx Pixel
		esz byte
		sct bool
	}{
		// Case 000
		//
		//      o |
		//     ---+---
		//        | e
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 107, 101}, opx: Pixel{111, 129}, osz: 80,
			ebc: Bucket{115, 123, 108, 100}, epx: Pixel{157, 118}, esz: 20,
			sct: false,
		},
		// Case 001
		//
		//      o |
		//     ---+---
		//        | e
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 107, 101}, opx: Pixel{111, 129}, osz: 80,
			ebc: Bucket{115, 123, 108, 100}, epx: Pixel{157, 118}, esz: 60,
			sct: true,
		},
		// Case 002
		//
		//        | o
		//     ---+---
		//      e |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 108, 101}, opx: Pixel{135, 109}, osz: 40,
			ebc: Bucket{115, 123, 107, 100}, epx: Pixel{159, 114}, esz: 20,
			sct: false,
		},
		// Case 003
		//
		//        | o
		//     ---+---
		//      e |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 108, 101}, opx: Pixel{135, 109}, osz: 50,
			ebc: Bucket{115, 123, 107, 100}, epx: Pixel{159, 114}, esz: 20,
			sct: true,
		},
		// Case 004
		//
		//      e | o
		//     ---+---
		//        |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 110, 101}, opx: Pixel{142, 117}, osz: 40,
			ebc: Bucket{115, 123, 109, 101}, epx: Pixel{115, 112}, esz: 20,
			sct: false,
		},
		// Case 005
		//
		//      e | o
		//     ---+---
		//        |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 110, 101}, opx: Pixel{142, 117}, osz: 50,
			ebc: Bucket{115, 123, 109, 101}, epx: Pixel{115, 112}, esz: 50,
			sct: true,
		},
		// Case 006
		//
		//      o | e
		//     ---+---
		//        |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 101}, epx: Pixel{142, 117}, esz: 50,
			ebc: Bucket{115, 123, 110, 101}, opx: Pixel{115, 112}, osz: 50,
			sct: true,
		},
		// Case 007
		//
		//      o |
		//     ---+---
		//      e |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 100}, epx: Pixel{142, 117}, esz: 40,
			ebc: Bucket{115, 122, 109, 163}, opx: Pixel{115, 112}, osz: 10,
			sct: false,
		},
		// Case 008
		//
		//      o |
		//     ---+---
		//      e |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 100}, epx: Pixel{142, 117}, esz: 40,
			ebc: Bucket{115, 122, 109, 163}, opx: Pixel{115, 112}, osz: 30,
			sct: false,
		},
		// Case 009
		//
		//      o |
		//     ---+---
		//      e |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 100}, epx: Pixel{142, 115}, esz: 40,
			ebc: Bucket{115, 122, 109, 163}, opx: Pixel{115, 114}, osz: 30,
			sct: true,
		},
		// Case 010
		//
		//     oe |
		//     ---+---
		//        |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 144}, epx: Pixel{142, 115}, esz: 15,
			ebc: Bucket{115, 123, 109, 144}, opx: Pixel{115, 114}, osz: 12,
			sct: false,
		},
		// Case 011
		//
		//     oe |
		//     ---+---
		//        |
		//
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 144}, epx: Pixel{142, 115}, esz: 15,
			ebc: Bucket{115, 123, 109, 144}, opx: Pixel{115, 114}, osz: 13,
			sct: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			sct := Intersect(tc.obc, tc.opx, tc.osz, tc.ebc, tc.epx, tc.esz)

			if sct != tc.sct {
				t.Fatal("expected", tc.sct, "got", sct)
			}
		})
	}
}

func Benchmark_Matrix_Intersect(b *testing.B) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		osz byte
		ebc Bucket
		epx Pixel
		esz byte
	}{
		// Case 000, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 107, 101}, opx: Pixel{111, 129}, osz: 80,
			ebc: Bucket{115, 123, 108, 100}, epx: Pixel{157, 118}, esz: 20,
		},
		// Case 001, ~5.50 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 107, 101}, opx: Pixel{111, 129}, osz: 80,
			ebc: Bucket{115, 123, 108, 100}, epx: Pixel{157, 118}, esz: 60,
		},
		// Case 002, ~5.50 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 108, 101}, opx: Pixel{135, 109}, osz: 40,
			ebc: Bucket{115, 123, 107, 100}, epx: Pixel{159, 114}, esz: 20,
		},
		// Case 003, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 108, 101}, opx: Pixel{135, 109}, osz: 50,
			ebc: Bucket{115, 123, 107, 100}, epx: Pixel{159, 114}, esz: 20,
		},
		// Case 004, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 110, 101}, opx: Pixel{142, 117}, osz: 40,
			ebc: Bucket{115, 123, 109, 101}, epx: Pixel{115, 112}, esz: 20,
		},
		// Case 005, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 110, 101}, opx: Pixel{142, 117}, osz: 50,
			ebc: Bucket{115, 123, 109, 101}, epx: Pixel{115, 112}, esz: 50,
		},
		// Case 006, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 101}, epx: Pixel{142, 117}, esz: 50,
			ebc: Bucket{115, 123, 110, 101}, opx: Pixel{115, 112}, osz: 50,
		},
		// Case 007, ~5.80 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 100}, epx: Pixel{142, 117}, esz: 40,
			ebc: Bucket{115, 122, 109, 163}, opx: Pixel{115, 112}, osz: 10,
		},
		// Case 008, ~5.70 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 100}, epx: Pixel{142, 117}, esz: 40,
			ebc: Bucket{115, 122, 109, 163}, opx: Pixel{115, 112}, osz: 30,
		},
		// Case 009, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 100}, epx: Pixel{142, 115}, esz: 40,
			ebc: Bucket{115, 122, 109, 163}, opx: Pixel{115, 114}, osz: 30,
		},
		// Case 010, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 144}, epx: Pixel{142, 115}, esz: 15,
			ebc: Bucket{115, 123, 109, 144}, opx: Pixel{115, 114}, osz: 12,
		},
		// Case 011, ~5.60 ns/op
		{
			//           x0,  y0,  x1,  y1,              x2,  y2
			obc: Bucket{115, 123, 109, 144}, epx: Pixel{142, 115}, esz: 15,
			ebc: Bucket{115, 123, 109, 144}, opx: Pixel{115, 114}, osz: 13,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Intersect(tc.obc, tc.opx, tc.osz, tc.ebc, tc.epx, tc.esz)
			}
		})
	}
}
