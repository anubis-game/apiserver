package vector

import (
	"fmt"
	"slices"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_Inside(t *testing.T) {
	testCases := []struct {
		scr screen
		ins []matrix.Coordinate
	}{
		// Case 000
		{
			scr: screen{},
			ins: nil,
		},
		// Case 001
		//
		//     +---------------+
		//     |               |
		//     |               |
		//     |       s       |
		//     |               |
		//     |               | P   P   P   P
		//     +---------------+
		//                       P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 768, bot: 1280, lef: 256},
			ins: nil,
		},
		// Case 002
		//
		//     +---------------+
		//     |               |
		//     |               |
		//     |       s       |
		//     |               |
		//     |             P | P   P   P
		//     +---------------+
		//                   P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 896, bot: 1280, lef: 384},
			ins: []matrix.Coordinate{
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
			},
		},
		// Case 003
		//
		//     +---------------+
		//     |               |
		//     |               |
		//     |       s       |
		//     |               |
		//     |         P   P | P   P
		//     +---------------+
		//               P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 1024, bot: 1280, lef: 512},
			ins: []matrix.Coordinate{
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
			},
		},
		// Case 004
		//
		//     +---------------+
		//     |               |
		//     |               |
		//     |       s       |
		//     |               |
		//     |     P   P   P | P
		//     +---------------+
		//           P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 1152, bot: 1280, lef: 640},
			ins: []matrix.Coordinate{
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
			},
		},
		// Case 005
		//
		//     +---------------+
		//     |               |
		//     |               |
		//     |       s       |
		//     |               |
		//     | P   P   P   P |
		//     +---------------+
		//       P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 1280, bot: 1280, lef: 768},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
			},
		},
		// Case 006
		//
		//       +---------------+
		//       |               |
		//       |               |
		//       |       s       |
		//       |               |
		//     P | P   P   P     |
		//       +---------------+
		//     P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 1536, bot: 1280, lef: 1024},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
			},
		},
		// Case 007
		//
		//           +---------------+
		//           |               |
		//           |               |
		//           |       s       |
		//           |               |
		//     P   P | P   P         |
		//           +---------------+
		//     P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 1664, bot: 1280, lef: 1152},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
			},
		},
		// Case 008
		//
		//               +---------------+
		//               |               |
		//               |               |
		//               |       s       |
		//               |               |
		//     P   P   P | P             |
		//               +---------------+
		//     P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 1792, bot: 1280, lef: 1280},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
			},
		},
		// Case 009
		//
		//                   +---------------+
		//                   |               |
		//                   |               |
		//                   |       s       |
		//                   |               |
		//     P   P   P   P |               |
		//                   +---------------+
		//     P   p   p   p
		//
		{
			scr: screen{top: 1792, rig: 1920, bot: 1280, lef: 1408},
			ins: nil,
		},
		// Case 010
		//
		//                   +---------------+
		//                   |               |
		//     P   P   P   P |               |
		//                   |       s       |
		//     P   p   p   p |               |
		//                   |               |
		//                   +---------------+
		//
		{
			scr: screen{top: 1536, rig: 1920, bot: 1024, lef: 1408},
			ins: nil,
		},
		// Case 011
		//
		//               +---------------+
		//               |               |
		//     P   P   P | P             |
		//               |       s       |
		//     P   p   p | p             |
		//               |               |
		//               +---------------+
		//
		{
			scr: screen{top: 1536, rig: 1792, bot: 1024, lef: 1280},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
			},
		},
		// Case 012
		//
		//           +---------------+
		//           |               |
		//     P   P | P   P         |
		//           |       s       |
		//     P   p | p   p         |
		//           |               |
		//           +---------------+
		//
		{
			scr: screen{top: 1536, rig: 1664, bot: 1024, lef: 1152},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
			},
		},
		// Case 013
		//
		//       +---------------+
		//       |               |
		//     P | P   P   P     |
		//       |       s       |
		//     P | p   p   p     |
		//       |               |
		//       +---------------+
		//
		{
			scr: screen{top: 1536, rig: 1536, bot: 1024, lef: 1024},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
			},
		},
		// Case 014
		//
		//     +---------------+
		//     |               |
		//     | P   P   P   P |
		//     |       s       |
		//     | P   p   p   p |
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1536, rig: 1408, bot: 1024, lef: 896},
			ins: []matrix.Coordinate{
				// x=1280 y=1280
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 015
		//
		//     +---------------+
		//     |               |
		//     |     P   P   P | P
		//     |       s       |
		//     |     P   p   p | p
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1536, rig: 1152, bot: 1024, lef: 640},
			ins: []matrix.Coordinate{
				// x=1152 y=1280
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 016
		//
		//     +---------------+
		//     |               |
		//     |         P   P | P   P
		//     |       s       |
		//     |         P   p | p   p
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1536, rig: 1024, bot: 1024, lef: 512},
			ins: []matrix.Coordinate{
				// x=1024 y=1280
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 017
		//
		//     +---------------+
		//     |               |
		//     |             P | P   P   P
		//     |       s       |
		//     |             P | p   p   p
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1536, rig: 896, bot: 1024, lef: 384},
			ins: []matrix.Coordinate{
				// x=896 y=1280
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 018
		//
		//     +---------------+
		//     |               |
		//     |               | P   P   P   P
		//     |       s       |
		//     |               | P   p   p   p
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1536, rig: 768, bot: 1024, lef: 256},
			ins: nil,
		},
		// Case 019
		//
		//                       P   P   P   P
		//     +---------------+
		//     |               | P   p   p   p
		//     |               |
		//     |       s       |
		//     |               |
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1152, rig: 768, bot: 640, lef: 256},
			ins: nil,
		},
		// Case 020
		//
		//                   P   P   P   P
		//     +---------------+
		//     |             P | p   p   p
		//     |               |
		//     |       s       |
		//     |               |
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1152, rig: 896, bot: 640, lef: 384},
			ins: []matrix.Coordinate{
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 021
		//
		//               P   P   P   P
		//     +---------------+
		//     |         P   p | p   p
		//     |               |
		//     |       s       |
		//     |               |
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1152, rig: 1024, bot: 640, lef: 512},
			ins: []matrix.Coordinate{
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 022
		//
		//           P   P   P   P
		//     +---------------+
		//     |     P   p   p | p
		//     |               |
		//     |       s       |
		//     |               |
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1152, rig: 1152, bot: 640, lef: 640},
			ins: []matrix.Coordinate{
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 023
		//
		//       P   P   P   P
		//     +---------------+
		//     | P   p   p   p |
		//     |               |
		//     |       s       |
		//     |               |
		//     |               |
		//     +---------------+
		//
		{
			scr: screen{top: 1152, rig: 1280, bot: 640, lef: 768},
			ins: []matrix.Coordinate{
				// x=896 y=1152
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
		},
		// Case 024
		//
		//     P   P   P   P
		//       +---------------+
		//     P | p   p   p     |
		//       |               |
		//       |       s       |
		//       |               |
		//       |               |
		//       +---------------+
		//
		{
			scr: screen{top: 1152, rig: 1536, bot: 640, lef: 1024},
			ins: nil,
		},
		// Case 025
		//
		//     P   P   P   P
		//           +---------------+
		//     P   p | p   p         |
		//           |               |
		//           |       s       |
		//           |               |
		//           |               |
		//           +---------------+
		//
		{
			scr: screen{top: 1152, rig: 1664, bot: 640, lef: 1152},
			ins: nil,
		},
		// Case 026
		//
		//     P   P   P   P
		//               +---------------+
		//     P   p   p | p             |
		//               |               |
		//               |       s       |
		//               |               |
		//               |               |
		//               +---------------+
		//
		{
			scr: screen{top: 1152, rig: 1792, bot: 640, lef: 1280},
			ins: nil,
		},
		// Case 027
		//
		//     P   P   P   P
		//                   +---------------+
		//     P   p   p   p |               |
		//                   |               |
		//                   |       s       |
		//                   |               |
		//                   |               |
		//                   +---------------+
		//
		{
			scr: screen{top: 1152, rig: 1920, bot: 640, lef: 1408},
			ins: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = tesVec()
			}

			{
				tesUpd(vec)
			}

			ins := vec.Inside(tc.scr.top, tc.scr.rig, tc.scr.bot, tc.scr.lef)

			if !slices.Equal(ins, tc.ins) {
				t.Fatalf("expected %#v got %#v", tc.ins, ins)
			}
		})
	}
}

func Benchmark_Vector_Inside(b *testing.B) {
	testCases := []struct {
		scr screen
	}{
		// Case 000, ~2 ns/op
		{
			scr: screen{},
		},
		// Case 001, ~2 ns/op
		{
			scr: screen{top: 1792, rig: 768, bot: 1280, lef: 256},
		},
		// Case 002, ~42 ns/op, 2 allocs/op
		{
			scr: screen{top: 1792, rig: 896, bot: 1280, lef: 384},
		},
		// Case 003, ~117 ns/op, 5 allocs/op
		{
			scr: screen{top: 1792, rig: 1024, bot: 1280, lef: 512},
		},
		// Case 004, ~119 ns/op, 5 allocs/op
		{
			scr: screen{top: 1792, rig: 1152, bot: 1280, lef: 640},
		},
		// Case 005, ~177 ns/op, 6 allocs/op
		{
			scr: screen{top: 1792, rig: 1280, bot: 1280, lef: 768},
		},
		// Case 006, ~124 ns/op, 5 allocs/op
		{
			scr: screen{top: 1792, rig: 1536, bot: 1280, lef: 1024},
		},
		// Case 007, ~90 ns/op, 4 allocs/op
		{
			scr: screen{top: 1792, rig: 1664, bot: 1280, lef: 1152},
		},
		// Case 008, ~47 ns/op, 2 allocs/op
		{
			scr: screen{top: 1792, rig: 1792, bot: 1280, lef: 1280},
		},
		// Case 009, ~2 ns/op
		{
			scr: screen{top: 1792, rig: 1920, bot: 1280, lef: 1408},
		},
		// Case 010, ~2 ns/op
		{
			scr: screen{top: 1536, rig: 1920, bot: 1024, lef: 1408},
		},
		// Case 011, ~47 ns/op, 2 allocs/op
		{
			scr: screen{top: 1536, rig: 1792, bot: 1024, lef: 1280},
		},
		// Case 012, ~88 ns/op, 4 allocs/op
		{
			scr: screen{top: 1536, rig: 1664, bot: 1024, lef: 1152},
		},
		// Case 013, ~122 ns/op, 5 allocs/op
		{
			scr: screen{top: 1536, rig: 1536, bot: 1024, lef: 1024},
		},
		// Case 014, ~182 ns/op, 6 allocs/op
		{
			scr: screen{top: 1536, rig: 1408, bot: 1024, lef: 896},
		},
		// Case 015, ~182 ns/op, 6 allocs/op
		{
			scr: screen{top: 1536, rig: 1152, bot: 1024, lef: 640},
		},
		// Case 016, ~120 ns/op, 5 allocs/op
		{
			scr: screen{top: 1536, rig: 1024, bot: 1024, lef: 512},
		},
		// Case 017, ~83 ns/op, 4 allocs/op
		{
			scr: screen{top: 1536, rig: 896, bot: 1024, lef: 384},
		},
		// Case 018, ~2 ns/op
		{
			scr: screen{top: 1536, rig: 768, bot: 1024, lef: 256},
		},
		// Case 019, ~2 ns/op
		{
			scr: screen{top: 1152, rig: 768, bot: 640, lef: 256},
		},
		// Case 020, ~60 ns/op, 3 allocs/op
		{
			scr: screen{top: 1152, rig: 896, bot: 640, lef: 384},
		},
		// Case 021, ~60 ns/op, 3 allocs/op
		{
			scr: screen{top: 1152, rig: 1024, bot: 640, lef: 512},
		},
		// Case 022, ~60 ns/op, 3 allocs/op
		{
			scr: screen{top: 1152, rig: 1152, bot: 640, lef: 640},
		},
		// Case 023, ~70 ns/op, 3 allocs/op
		{
			scr: screen{top: 1152, rig: 1280, bot: 640, lef: 768},
		},
		// Case 024, ~15 ns/op
		{
			scr: screen{top: 1152, rig: 1536, bot: 640, lef: 1024},
		},
		// Case 025, ~15 ns/op
		{
			scr: screen{top: 1152, rig: 1664, bot: 640, lef: 1152},
		},
		// Case 026, ~15 ns/op
		{
			scr: screen{top: 1152, rig: 1792, bot: 640, lef: 1280},
		},
		// Case 027, ~2 ns/op
		{
			scr: screen{top: 1152, rig: 1920, bot: 640, lef: 1408},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			var vec *Vector
			{
				vec = tesVec()
			}

			{
				tesUpd(vec)
			}

			for b.Loop() {
				vec.Inside(tc.scr.top, tc.scr.rig, tc.scr.bot, tc.scr.lef)
			}
		})
	}
}
