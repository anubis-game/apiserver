package window

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Window_Crx_True(t *testing.T) {
	testCases := []struct {
		win *Window
	}{
		// Case 000
		//
		//     +------+
		//     |   +--|-------+
		//     +---|--+       |
		//         |          |
		//         |          |
		//         +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 75, Y: 175},
				ctr: object.Object{X: 125, Y: 225},
			},
		},
		// Case 001
		//
		//           +------+
		//         +-|------|-+
		//         | +------+ |
		//         |          |
		//         |          |
		//         +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 175},
				ctr: object.Object{X: 175, Y: 225},
			},
		},
		// Case 002
		//
		//                 +------+
		//         +-------|--+   |
		//         |       +--|---+
		//         |          |
		//         |          |
		//         +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 175, Y: 175},
				ctr: object.Object{X: 225, Y: 225},
			},
		},
		// Case 003
		//
		//         +----------+
		//         |       +------+
		//         |       |  |   |
		//         |       +------+
		//         +----------+
		//
		//
		{
			win: &Window{
				cbl: object.Object{X: 175, Y: 125},
				ctr: object.Object{X: 225, Y: 175},
			},
		},
		// Case 004
		//
		//         +----------+
		//         |          |
		//         |          |
		//         |       +--|---+
		//         +-------|--+   |
		//                 +------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 175, Y: 75},
				ctr: object.Object{X: 225, Y: 125},
			},
		},
		// Case 005
		//
		//         +----------+
		//         |          |
		//         |          |
		//         | +------+ |
		//         +-|------|-+
		//           +------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 75},
				ctr: object.Object{X: 175, Y: 125},
			},
		},
		// Case 006
		//
		//         +----------+
		//         |          |
		//         |          |
		//     +---|--+       |
		//     |   +--|----- -+
		//     +------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 75, Y: 75},
				ctr: object.Object{X: 125, Y: 125},
			},
		},
		// Case 007
		//
		//         +----------+
		//     +------+       |
		//     |   |  |       |
		//     +------+       |
		//         +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 75, Y: 125},
				ctr: object.Object{X: 125, Y: 175},
			},
		},
		// Case 008
		//
		//         +----------+
		//         | +------+ |
		//         | |      | |
		//         | +------+ |
		//         +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 125},
				ctr: object.Object{X: 175, Y: 175},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := &Window{
				cbl: object.Object{X: 100, Y: 100},
				ctr: object.Object{X: 200, Y: 200},
			}

			{
				has := tc.win.Crx(win)
				if has != true {
					t.Fatalf("expected %#v got %#v", true, has)
				}
			}

			{
				has := win.Crx(tc.win)
				if has != true {
					t.Fatalf("expected %#v got %#v", true, has)
				}
			}
		})
	}
}

func Test_Window_Crx_False(t *testing.T) {
	testCases := []struct {
		win *Window
	}{
		// Case 000
		//
		//     +------+
		//     |      |
		//     +------+
		//              +----------+
		//              |          |
		//              |          |
		//              |          |
		//              +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 25, Y: 225},
				ctr: object.Object{X: 75, Y: 275},
			},
		},
		// Case 001
		//
		//                +------+
		//                |      |
		//                +------+
		//              +----------+
		//              |          |
		//              |          |
		//              |          |
		//              +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 225},
				ctr: object.Object{X: 175, Y: 275},
			},
		},
		// Case 002
		//
		//                           +------+
		//                           |      |
		//                           +------+
		//              +----------+
		//              |          |
		//              |          |
		//              |          |
		//              +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 225, Y: 225},
				ctr: object.Object{X: 275, Y: 275},
			},
		},
		// Case 003
		//
		//              +----------+
		//              |          | +------+
		//              |          | |      |
		//              |          | +------+
		//              +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 225, Y: 125},
				ctr: object.Object{X: 275, Y: 175},
			},
		},
		// Case 004
		//
		//              +----------+
		//              |          |
		//              |          |
		//              |          |
		//              +----------+
		//                           +------+
		//                           |      |
		//                           +------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 225, Y: 25},
				ctr: object.Object{X: 275, Y: 75},
			},
		},
		// Case 005
		//
		//              +----------+
		//              |          |
		//              |          |
		//              |          |
		//              +----------+
		//                +------+
		//                |      |
		//                +------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 25},
				ctr: object.Object{X: 175, Y: 75},
			},
		},
		// Case 006
		//
		//              +----------+
		//              |          |
		//              |          |
		//              |          |
		//              +----------+
		//     +------+
		//     |      |
		//     +------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 25, Y: 25},
				ctr: object.Object{X: 75, Y: 75},
			},
		},
		// Case 007
		//
		//              +----------+
		//     +------+ |          |
		//     |      | |          |
		//     +------+ |          |
		//              +----------+
		//
		{
			win: &Window{
				cbl: object.Object{X: 25, Y: 125},
				ctr: object.Object{X: 75, Y: 175},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := &Window{
				cbl: object.Object{X: 100, Y: 100},
				ctr: object.Object{X: 200, Y: 200},
			}

			{
				has := tc.win.Crx(win)
				if has != false {
					t.Fatalf("expected %#v got %#v", false, has)
				}
			}

			{
				has := win.Crx(tc.win)
				if has != false {
					t.Fatalf("expected %#v got %#v", false, has)
				}
			}
		})
	}
}

func Test_Window_Has_True(t *testing.T) {
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

			has := win.Has(tc.obj)

			if has != true {
				t.Fatalf("expected %#v got %#v", true, has)
			}
		})
	}
}

func Test_Window_Has_False(t *testing.T) {
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
				has := win.Has(x)

				if has != false {
					t.Fatalf("expected %#v got %#v", false, has)
				}
			}
		})
	}
}

func Test_Window_Key(t *testing.T) {
	testCases := []struct {
		win *Window
		key [4]object.Object
	}{
		// Case 000
		{
			win: &Window{
				cbl: object.Object{X: 100, Y: 200},
				ctr: object.Object{X: 150, Y: 250},
			},
			key: [4]object.Object{
				{X: 0, Y: 0},
				{X: 0, Y: 0},
				{X: 0, Y: 0},
				{X: 0, Y: 0},
			},
		},
		// Case 001
		{
			win: &Window{
				cbl: object.Object{X: 1_500, Y: 250},
				ctr: object.Object{X: 2_000, Y: 750},
			},
			key: [4]object.Object{
				{X: 1024, Y: 0},
				{X: 1536, Y: 0},
				{X: 1024, Y: 512},
				{X: 1536, Y: 512},
			},
		},
		// Case 002
		{
			win: &Window{
				cbl: object.Object{X: 2_366, Y: 17},
				ctr: object.Object{X: 3_646, Y: 1_297},
			},
			key: [4]object.Object{
				{X: 2048, Y: 0},
				{X: 3584, Y: 0},
				{X: 2048, Y: 1024},
				{X: 3584, Y: 1024},
			},
		},
		// Case 003
		{
			win: &Window{
				cbl: object.Object{X: 15_775, Y: 4_096},
				ctr: object.Object{X: 17_245, Y: 5_566},
			},
			key: [4]object.Object{
				{X: 15360, Y: 4096},
				{X: 16896, Y: 4096},
				{X: 15360, Y: 5120},
				{X: 16896, Y: 5120},
			},
		},
		// Case 004
		{
			win: &Window{
				cbl: object.Object{X: 365_283, Y: 22_547},
				ctr: object.Object{X: 369_379, Y: 26_643},
			},
			key: [4]object.Object{
				{X: 365056, Y: 22528},
				{X: 369152, Y: 22528},
				{X: 365056, Y: 26624},
				{X: 369152, Y: 26624},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := tc.win.Key()

			if !reflect.DeepEqual(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}

var crxSnk bool

func Benchmark_Window_Crx(b *testing.B) {
	testCases := []struct {
		win *Window
	}{
		// Case 000, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 75, Y: 175},
				ctr: object.Object{X: 125, Y: 225},
			},
		},
		// Case 001, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 175},
				ctr: object.Object{X: 175, Y: 225},
			},
		},
		// Case 002, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 175, Y: 175},
				ctr: object.Object{X: 225, Y: 225},
			},
		},
		// Case 003, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 175, Y: 125},
				ctr: object.Object{X: 225, Y: 175},
			},
		},
		// Case 004, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 175, Y: 75},
				ctr: object.Object{X: 225, Y: 125},
			},
		},
		// Case 005, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 75},
				ctr: object.Object{X: 175, Y: 125},
			},
		},
		// Case 006, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 75, Y: 75},
				ctr: object.Object{X: 125, Y: 125},
			},
		},
		// Case 007, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 75, Y: 125},
				ctr: object.Object{X: 125, Y: 175},
			},
		},
		// Case 008, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 125},
				ctr: object.Object{X: 175, Y: 175},
			},
		},
		// Case 009, ~0.60 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 25, Y: 225},
				ctr: object.Object{X: 75, Y: 275},
			},
		},
		// Case 010, ~0.90 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 225},
				ctr: object.Object{X: 175, Y: 275},
			},
		},
		// Case 011, ~0.70 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 225, Y: 225},
				ctr: object.Object{X: 275, Y: 275},
			},
		},
		// Case 012, ~0.70 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 225, Y: 125},
				ctr: object.Object{X: 275, Y: 175},
			},
		},
		// Case 013, ~0.60 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 225, Y: 25},
				ctr: object.Object{X: 275, Y: 75},
			},
		},
		// Case 014, ~0.60 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 125, Y: 25},
				ctr: object.Object{X: 175, Y: 75},
			},
		},
		// Case 015, ~0.60 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 25, Y: 25},
				ctr: object.Object{X: 75, Y: 75},
			},
		},
		// Case 016, ~0.60 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 25, Y: 125},
				ctr: object.Object{X: 75, Y: 175},
			},
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
				crxSnk = win.Crx(tc.win)
			}
		})
	}
}

var hasSnk bool

func Benchmark_Window_Has(b *testing.B) {
	testCases := []struct {
		obj object.Object
	}{
		// Case 000, 0.60 ns/op
		{
			obj: object.Object{X: 100, Y: 200},
		},
		// Case 001, 0.60 ns/op
		{
			obj: object.Object{X: 150, Y: 200},
		},
		// Case 002, 0.60 ns/op
		{
			obj: object.Object{X: 200, Y: 200},
		},
		// Case 003, 0.60 ns/op
		{
			obj: object.Object{X: 200, Y: 150},
		},
		// Case 004, 0.60 ns/op
		{
			obj: object.Object{X: 200, Y: 100},
		},
		// Case 005, 0.60 ns/op
		{
			obj: object.Object{X: 150, Y: 100},
		},
		// Case 006, 0.60 ns/op
		{
			obj: object.Object{X: 100, Y: 100},
		},
		// Case 007, 0.60 ns/op
		{
			obj: object.Object{X: 100, Y: 150},
		},
		// Case 008, 0.60 ns/op
		{
			obj: object.Object{X: 137, Y: 143},
		},
		// Case 009, 0.60 ns/op
		{
			obj: object.Object{X: 86, Y: 206},
		},
		// Case 010, 0.60 ns/op
		{
			obj: object.Object{X: 149, Y: 201},
		},
		// Case 011, 0.60 ns/op
		{
			obj: object.Object{X: 201, Y: 200},
		},
		// Case 012, 0.60 ns/op
		{
			obj: object.Object{X: 264, Y: 149},
		},
		// Case 013, 0.60 ns/op
		{
			obj: object.Object{X: 201, Y: 99},
		},
		// Case 014, 0.60 ns/op
		{
			obj: object.Object{X: 149, Y: 99},
		},
		// Case 015, 0.60 ns/op
		{
			obj: object.Object{X: 100, Y: 99},
		},
		// Case 016, 0.60 ns/op
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
				hasSnk = win.Has(tc.obj)
			}
		})
	}
}

var keySnk [4]object.Object

func Benchmark_Window_Key(b *testing.B) {
	testCases := []struct {
		win *Window
	}{
		// Case 000, 2.50 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 100, Y: 200},
				ctr: object.Object{X: 150, Y: 250},
			},
		},
		// Case 001, 2.40 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 1_500, Y: 250},
				ctr: object.Object{X: 2_000, Y: 750},
			},
		},
		// Case 002, 2.40 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 2_366, Y: 17},
				ctr: object.Object{X: 3_646, Y: 1_297},
			},
		},
		// Case 003, 2.40 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 15_775, Y: 4_096},
				ctr: object.Object{X: 17_245, Y: 5_566},
			},
		},
		// Case 004, 2.40 ns/op
		{
			win: &Window{
				cbl: object.Object{X: 365_283, Y: 22_547},
				ctr: object.Object{X: 369_379, Y: 26_643},
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for range b.N {
				keySnk = tc.win.Key()
			}
		})
	}
}
