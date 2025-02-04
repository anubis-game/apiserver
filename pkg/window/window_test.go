package window

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Window_has_True(t *testing.T) {
	testCases := []struct {
		obj object.Object
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
			obj: object.Object{X: 100, Y: 200},
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
			obj: object.Object{X: 150, Y: 200},
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
			obj: object.Object{X: 200, Y: 200},
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
			obj: object.Object{X: 200, Y: 150},
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
			obj: object.Object{X: 200, Y: 100},
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
			obj: object.Object{X: 150, Y: 100},
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
			obj: object.Object{X: 100, Y: 100},
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
			obj: object.Object{X: 100, Y: 150},
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
			obj: object.Object{X: 137, Y: 143},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := &Window{
				cbl: object.Object{X: 100, Y: 100},
				ctr: object.Object{X: 200, Y: 200},
			}

			has := win.has(tc.obj)

			if has != true {
				t.Fatalf("expected %#v got %#v", true, has)
			}
		})
	}
}

func Test_Window_has_False(t *testing.T) {
	testCases := []struct {
		obj []object.Object
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
			obj: []object.Object{
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
			obj: []object.Object{
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
			obj: []object.Object{
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
			obj: []object.Object{
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
			obj: []object.Object{
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
			obj: []object.Object{
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
			obj: []object.Object{
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
			obj: []object.Object{
				{X: 99, Y: 149},
				{X: 99, Y: 150},
				{X: 92, Y: 151},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := &Window{
				cbl: object.Object{X: 100, Y: 100},
				ctr: object.Object{X: 200, Y: 200},
			}

			for _, x := range tc.obj {
				has := win.has(x)

				if has != false {
					t.Fatalf("expected %#v got %#v", false, has)
				}
			}
		})
	}
}

func Benchmark_Window_has(b *testing.B) {
	testCases := []struct {
		obj object.Object
	}{
		// Case 000, 0.30 ns/op
		{
			obj: object.Object{X: 100, Y: 200},
		},
		// Case 001, 0.30 ns/op
		{
			obj: object.Object{X: 150, Y: 200},
		},
		// Case 002, 0.30 ns/op
		{
			obj: object.Object{X: 200, Y: 200},
		},
		// Case 003, 0.30 ns/op
		{
			obj: object.Object{X: 200, Y: 150},
		},
		// Case 004, 0.30 ns/op
		{
			obj: object.Object{X: 200, Y: 100},
		},
		// Case 005, 0.30 ns/op
		{
			obj: object.Object{X: 150, Y: 100},
		},
		// Case 006, 0.30 ns/op
		{
			obj: object.Object{X: 100, Y: 100},
		},
		// Case 007, 0.30 ns/op
		{
			obj: object.Object{X: 100, Y: 150},
		},
		// Case 008, 0.30 ns/op
		{
			obj: object.Object{X: 137, Y: 143},
		},
		// Case 009, 0.30 ns/op
		{
			obj: object.Object{X: 86, Y: 206},
		},
		// Case 010, 0.30 ns/op
		{
			obj: object.Object{X: 149, Y: 201},
		},
		// Case 011, 0.30 ns/op
		{
			obj: object.Object{X: 201, Y: 200},
		},
		// Case 012, 0.30 ns/op
		{
			obj: object.Object{X: 264, Y: 149},
		},
		// Case 013, 0.30 ns/op
		{
			obj: object.Object{X: 201, Y: 99},
		},
		// Case 014, 0.30 ns/op
		{
			obj: object.Object{X: 149, Y: 99},
		},
		// Case 015, 0.30 ns/op
		{
			obj: object.Object{X: 100, Y: 99},
		},
		// Case 016, 0.30 ns/op
		{
			obj: object.Object{X: 99, Y: 150},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			win := &Window{
				cbl: object.Object{X: 100, Y: 100},
				ctr: object.Object{X: 200, Y: 200},
			}

			b.ResetTimer()
			for range b.N {
				win.has(tc.obj)
			}
		})
	}
}
