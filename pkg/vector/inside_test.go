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
				occ: &Occupy{
					Top: 225,
					Rig: 125,
					Bot: 175,
					Lef: 75,
				},
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
				occ: &Occupy{
					Top: 225,
					Rig: 175,
					Bot: 175,
					Lef: 125,
				},
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
				occ: &Occupy{
					Top: 225,
					Rig: 225,
					Bot: 175,
					Lef: 175,
				},
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
				occ: &Occupy{
					Top: 175,
					Rig: 225,
					Bot: 125,
					Lef: 175,
				},
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
				occ: &Occupy{
					Top: 125,
					Rig: 225,
					Bot: 75,
					Lef: 175,
				},
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
				occ: &Occupy{
					Top: 125,
					Rig: 175,
					Bot: 75,
					Lef: 125,
				},
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
				occ: &Occupy{
					Top: 125,
					Rig: 125,
					Bot: 75,
					Lef: 75,
				},
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
				occ: &Occupy{
					Top: 175,
					Rig: 125,
					Bot: 125,
					Lef: 75,
				},
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
				occ: &Occupy{
					Top: 175,
					Rig: 175,
					Bot: 125,
					Lef: 125,
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			vec := &Vector{
				scr: &Screen{
					Top: 200,
					Rig: 200,
					Bot: 100,
					Lef: 100,
				},
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
				occ: &Occupy{
					Top: 275,
					Rig: 75,
					Bot: 225,
					Lef: 25,
				},
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
				occ: &Occupy{
					Top: 275,
					Rig: 175,
					Bot: 225,
					Lef: 125,
				},
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
				occ: &Occupy{
					Top: 275,
					Rig: 275,
					Bot: 225,
					Lef: 225,
				},
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
				occ: &Occupy{
					Top: 175,
					Rig: 275,
					Bot: 125,
					Lef: 225,
				},
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
				occ: &Occupy{
					Top: 75,
					Rig: 275,
					Bot: 25,
					Lef: 225,
				},
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
				occ: &Occupy{
					Top: 75,
					Rig: 175,
					Bot: 25,
					Lef: 125,
				},
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
				occ: &Occupy{
					Top: 75,
					Rig: 75,
					Bot: 25,
					Lef: 25,
				},
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
				occ: &Occupy{
					Top: 175,
					Rig: 75,
					Bot: 125,
					Lef: 25,
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			vec := &Vector{
				scr: &Screen{
					Top: 200,
					Rig: 200,
					Bot: 100,
					Lef: 100,
				},
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
		// Case 000, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 225,
					Rig: 125,
					Bot: 175,
					Lef: 75,
				},
			},
		},
		// Case 001, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 225,
					Rig: 175,
					Bot: 175,
					Lef: 125,
				},
			},
		},
		// Case 002, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 225,
					Rig: 225,
					Bot: 175,
					Lef: 175,
				},
			},
		},
		// Case 003, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 175,
					Rig: 225,
					Bot: 125,
					Lef: 175,
				},
			},
		},
		// Case 004, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 125,
					Rig: 225,
					Bot: 75,
					Lef: 175,
				},
			},
		},
		// Case 005, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 125,
					Rig: 175,
					Bot: 75,
					Lef: 125,
				},
			},
		},
		// Case 006, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 125,
					Rig: 125,
					Bot: 75,
					Lef: 75,
				},
			},
		},
		// Case 007, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 175,
					Rig: 125,
					Bot: 125,
					Lef: 75,
				},
			},
		},
		// Case 008, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 175,
					Rig: 175,
					Bot: 125,
					Lef: 125,
				},
			},
		},
		// Case 009, ~0.60 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 275,
					Rig: 75,
					Bot: 225,
					Lef: 25,
				},
			},
		},
		// Case 010, ~0.60 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 275,
					Rig: 175,
					Bot: 225,
					Lef: 125,
				},
			},
		},
		// Case 011, ~0.60 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 275,
					Rig: 275,
					Bot: 225,
					Lef: 225,
				},
			},
		},
		// Case 012, ~0.60 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 175,
					Rig: 275,
					Bot: 125,
					Lef: 225,
				},
			},
		},
		// Case 013, ~0.60 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 75,
					Rig: 275,
					Bot: 25,
					Lef: 225,
				},
			},
		},
		// Case 014, ~1.00 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 75,
					Rig: 175,
					Bot: 25,
					Lef: 125,
				},
			},
		},
		// Case 015, ~0.80 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 75,
					Rig: 75,
					Bot: 25,
					Lef: 25,
				},
			},
		},
		// Case 016, ~0.80 ns/op
		{
			vec: &Vector{
				occ: &Occupy{
					Top: 175,
					Rig: 75,
					Bot: 125,
					Lef: 25,
				},
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			vec := &Vector{
				scr: &Screen{
					Top: 200,
					Rig: 200,
					Bot: 100,
					Lef: 100,
				},
			}

			b.ResetTimer()
			for range b.N {
				crxSnk = tc.vec.Inside(vec.Screen())
			}
		})
	}
}
