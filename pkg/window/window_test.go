package window

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Window_Has_True(t *testing.T) {
	testCases := []struct {
		obj matrix.Coordinate
	}{
		// Case 000
		//
		//                  200
		//         +----------+ 200
		//         | x        |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 100, Y: 200},
		},
		// Case 001
		//
		//                  200
		//         +----------+ 200
		//         |    x     |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 150, Y: 200},
		},
		// Case 002
		//
		//                  200
		//         +----------+ 200
		//         |        x |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 200, Y: 200},
		},
		// Case 003
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |        x |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 200, Y: 150},
		},
		// Case 004
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |        x |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 200, Y: 100},
		},
		// Case 005
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |    x     |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 150, Y: 100},
		},
		// Case 006
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         | x        |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 100, Y: 100},
		},
		// Case 007
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         | x        |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 100, Y: 150},
		},
		// Case 008
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |    x     |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: matrix.Coordinate{X: 137, Y: 143},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := &Window{
				cbl: matrix.Coordinate{X: 100, Y: 100},
				ctr: matrix.Coordinate{X: 200, Y: 200},
			}

			has := win.Has(tc.obj)

			if has != true {
				t.Fatalf("expected %#v got %#v", true, has)
			}
		})
	}
}

func Test_Window_Has_False(t *testing.T) {
	testCases := []struct {
		obj []matrix.Coordinate
	}{
		// Case 000
		//
		//       x          200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []matrix.Coordinate{
				{X: 86, Y: 206},
				{X: 99, Y: 201},
				{X: 99, Y: 200},
				{X: 100, Y: 201},
			},
		},
		// Case 001
		//
		//              x   200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []matrix.Coordinate{
				{X: 150, Y: 201},
				{X: 149, Y: 201},
				{X: 151, Y: 236},
			},
		},
		// Case 002
		//
		//                  200 x
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []matrix.Coordinate{
				{X: 208, Y: 299},
				{X: 201, Y: 201},
				{X: 201, Y: 200},
				{X: 200, Y: 201},
			},
		},
		// Case 003
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          | x
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []matrix.Coordinate{
				{X: 201, Y: 151},
				{X: 201, Y: 150},
				{X: 264, Y: 149},
			},
		},
		// Case 004
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100          x
		//
		{
			obj: []matrix.Coordinate{
				{X: 256, Y: 87},
				{X: 201, Y: 99},
				{X: 200, Y: 99},
				{X: 201, Y: 100},
			},
		},
		// Case 005
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100  x
		//
		{
			obj: []matrix.Coordinate{
				{X: 149, Y: 99},
				{X: 150, Y: 99},
				{X: 151, Y: 85},
			},
		},
		// Case 006
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//       x 100
		//
		{
			obj: []matrix.Coordinate{
				{X: 96, Y: 92},
				{X: 99, Y: 99},
				{X: 99, Y: 100},
				{X: 100, Y: 99},
			},
		},
		// Case 007
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//       x |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []matrix.Coordinate{
				{X: 99, Y: 149},
				{X: 99, Y: 150},
				{X: 92, Y: 151},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := &Window{
				cbl: matrix.Coordinate{X: 100, Y: 100},
				ctr: matrix.Coordinate{X: 200, Y: 200},
			}

			for _, x := range tc.obj {
				has := win.Has(x)

				if has != false {
					t.Fatalf("expected %#v got %#v", false, has)
				}
			}
		})
	}
}

func Benchmark_Window_Has(b *testing.B) {
	testCases := []struct {
		obj matrix.Coordinate
	}{
		// Case 000, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 100, Y: 200},
		},
		// Case 001, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 150, Y: 200},
		},
		// Case 002, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 200, Y: 200},
		},
		// Case 003, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 200, Y: 150},
		},
		// Case 004, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 200, Y: 100},
		},
		// Case 005, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 150, Y: 100},
		},
		// Case 006, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 100, Y: 100},
		},
		// Case 007, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 100, Y: 150},
		},
		// Case 008, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 137, Y: 143},
		},
		// Case 009, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 86, Y: 206},
		},
		// Case 010, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 149, Y: 201},
		},
		// Case 011, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 201, Y: 200},
		},
		// Case 012, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 264, Y: 149},
		},
		// Case 013, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 201, Y: 99},
		},
		// Case 014, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 149, Y: 99},
		},
		// Case 015, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 100, Y: 99},
		},
		// Case 016, 1.70 ns/op
		{
			obj: matrix.Coordinate{X: 99, Y: 150},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			win := &Window{
				cbl: matrix.Coordinate{X: 100, Y: 100},
				ctr: matrix.Coordinate{X: 200, Y: 200},
			}

			for b.Loop() {
				win.Has(tc.obj)
			}
		})
	}
}
