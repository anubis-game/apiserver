package vector

import (
	"fmt"
	"testing"
)

func Test_Vector_Inside_True(t *testing.T) {
	testCases := []struct {
		vec *Vector
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
			vec: &Vector{
				btp: 225,
				brg: 125,
				bbt: 175,
				blf: 75,
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
			vec: &Vector{
				btp: 225,
				brg: 175,
				bbt: 175,
				blf: 125,
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
			vec: &Vector{
				btp: 225,
				brg: 225,
				bbt: 175,
				blf: 175,
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
			vec: &Vector{
				btp: 175,
				brg: 225,
				bbt: 125,
				blf: 175,
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
			vec: &Vector{
				btp: 125,
				brg: 225,
				bbt: 75,
				blf: 175,
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
			vec: &Vector{
				btp: 125,
				brg: 175,
				bbt: 75,
				blf: 125,
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
			vec: &Vector{
				btp: 125,
				brg: 125,
				bbt: 75,
				blf: 75,
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
			vec: &Vector{
				btp: 175,
				brg: 125,
				bbt: 125,
				blf: 75,
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
			vec: &Vector{
				btp: 175,
				brg: 175,
				bbt: 125,
				blf: 125,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			vec := &Vector{
				vtp: 200,
				vrg: 200,
				vbt: 100,
				vlf: 100,
			}

			has := tc.vec.Inside(vec.Screen())

			if has != true {
				t.Fatalf("expected %#v got %#v", true, has)
			}
		})
	}
}

func Test_Vector_Inside_False(t *testing.T) {
	testCases := []struct {
		vec *Vector
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
			vec: &Vector{
				btp: 275,
				brg: 75,
				bbt: 225,
				blf: 25,
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
			vec: &Vector{
				btp: 275,
				brg: 175,
				bbt: 225,
				blf: 125,
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
			vec: &Vector{
				btp: 275,
				brg: 275,
				bbt: 225,
				blf: 225,
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
			vec: &Vector{
				btp: 175,
				brg: 275,
				bbt: 125,
				blf: 225,
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
			vec: &Vector{
				btp: 75,
				brg: 275,
				bbt: 25,
				blf: 225,
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
			vec: &Vector{
				btp: 75,
				brg: 175,
				bbt: 25,
				blf: 125,
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
			vec: &Vector{
				btp: 75,
				brg: 75,
				bbt: 25,
				blf: 25,
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
			vec: &Vector{
				btp: 175,
				brg: 75,
				bbt: 125,
				blf: 25,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			vec := &Vector{
				vtp: 200,
				vrg: 200,
				vbt: 100,
				vlf: 100,
			}

			has := tc.vec.Inside(vec.Screen())

			if has != false {
				t.Fatalf("expected %#v got %#v", false, has)
			}
		})
	}
}

var crxSnk bool

func Benchmark_Vector_Inside(b *testing.B) {
	testCases := []struct {
		vec *Vector
	}{
		// Case 000, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 225,
				brg: 125,
				bbt: 175,
				blf: 75,
			},
		},
		// Case 001, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 225,
				brg: 175,
				bbt: 175,
				blf: 125,
			},
		},
		// Case 002, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 225,
				brg: 225,
				bbt: 175,
				blf: 175,
			},
		},
		// Case 003, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 175,
				brg: 225,
				bbt: 125,
				blf: 175,
			},
		},
		// Case 004, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 125,
				brg: 225,
				bbt: 75,
				blf: 175,
			},
		},
		// Case 005, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 125,
				brg: 175,
				bbt: 75,
				blf: 125,
			},
		},
		// Case 006, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 125,
				brg: 125,
				bbt: 75,
				blf: 75,
			},
		},
		// Case 007, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 175,
				brg: 125,
				bbt: 125,
				blf: 75,
			},
		},
		// Case 008, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 175,
				brg: 175,
				bbt: 125,
				blf: 125,
			},
		},
		// Case 009, ~0.60 ns/op
		{
			vec: &Vector{
				btp: 275,
				brg: 75,
				bbt: 225,
				blf: 25,
			},
		},
		// Case 010, ~0.60 ns/op
		{
			vec: &Vector{
				btp: 275,
				brg: 175,
				bbt: 225,
				blf: 125,
			},
		},
		// Case 011, ~0.60 ns/op
		{
			vec: &Vector{
				btp: 275,
				brg: 275,
				bbt: 225,
				blf: 225,
			},
		},
		// Case 012, ~0.60 ns/op
		{
			vec: &Vector{
				btp: 175,
				brg: 275,
				bbt: 125,
				blf: 225,
			},
		},
		// Case 013, ~0.60 ns/op
		{
			vec: &Vector{
				btp: 75,
				brg: 275,
				bbt: 25,
				blf: 225,
			},
		},
		// Case 014, ~0.90 ns/op
		{
			vec: &Vector{
				btp: 75,
				brg: 175,
				bbt: 25,
				blf: 125,
			},
		},
		// Case 015, ~0.70 ns/op
		{
			vec: &Vector{
				btp: 75,
				brg: 75,
				bbt: 25,
				blf: 25,
			},
		},
		// Case 016, ~0.70 ns/op
		{
			vec: &Vector{
				btp: 175,
				brg: 75,
				bbt: 125,
				blf: 25,
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			vec := &Vector{
				vtp: 200,
				vrg: 200,
				vbt: 100,
				vlf: 100,
			}

			b.ResetTimer()
			for range b.N {
				crxSnk = tc.vec.Inside(vec.Screen())
			}
		})
	}
}
