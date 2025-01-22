package matrix

import (
	"fmt"
	"testing"
)

func Test_Matrix_Target_Quadrant_1(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		tbc Bucket
		tpx Pixel
	}{
		// Case 000, move +3 along x and +4 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1),   // quadrant 1
				byte(108), // 38.12° from 0°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			tbc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			tpx: Pixel{
				114, // x2
				133, // y2
			},
		},
		// Case 001, move +19 along x and 0 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				163, // x1
				101, // y1
			},
			opx: Pixel{
				151, // x2
				129, // y2
			},
			spc: Space{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			tbc: Bucket{
				151, // x0
				130, // y0
				100, // x1
				101, // y1
			},
			tpx: Pixel{
				106, // x2
				129, // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != tc.tbc[X0] {
				t.Fatal("expected", tc.tbc[X0], "got", tbc[X0])
			}
			if tbc[Y0] != tc.tbc[Y0] {
				t.Fatal("expected", tc.tbc[Y0], "got", tbc[Y0])
			}

			if tbc[X1] != tc.tbc[X1] {
				t.Fatal("expected", tc.tbc[X1], "got", tbc[X1])
			}
			if tbc[Y1] != tc.tbc[Y1] {
				t.Fatal("expected", tc.tbc[Y1], "got", tbc[Y1])
			}

			if tpx[X2] != tc.tpx[X2] {
				t.Fatal("expected", tc.tpx[X2], "got", tpx[X2])
			}
			if tpx[Y2] != tc.tpx[Y2] {
				t.Fatal("expected", tc.tpx[Y2], "got", tpx[Y2])
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Matrix_Target_Quadrant_2(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		mil float64
		tbc Bucket
		tpx Pixel
	}{
		// Case 000, move +4 along x and -3 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(2),   // quadrant 2
				byte(108), // 38.12° from 90°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			tbc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			tpx: Pixel{
				115, // x2
				126, // y2
			},
		},
		// Case 001, move 0 along x and -19 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(2),   // quadrant 2
				byte(253), // 89.29° from 90°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			tbc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			tpx: Pixel{
				111, // x2
				110, // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != tc.tbc[X0] {
				t.Fatal("expected", tc.tbc[X0], "got", tbc[X0])
			}
			if tbc[Y0] != tc.tbc[Y0] {
				t.Fatal("expected", tc.tbc[Y0], "got", tbc[Y0])
			}

			if tbc[X1] != tc.tbc[X1] {
				t.Fatal("expected", tc.tbc[X1], "got", tbc[X1])
			}
			if tbc[Y1] != tc.tbc[Y1] {
				t.Fatal("expected", tc.tbc[Y1], "got", tbc[Y1])
			}

			if tpx[X2] != tc.tpx[X2] {
				t.Fatal("expected", tc.tpx[X2], "got", tpx[X2])
			}
			if tpx[Y2] != tc.tpx[Y2] {
				t.Fatal("expected", tc.tpx[Y2], "got", tpx[Y2])
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Matrix_Target_Quadrant_3(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		mil float64
		tbc Bucket
		tpx Pixel
	}{
		// Case 000, move -3 along x and -4 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3),   // quadrant 3
				byte(108), // 38.12° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			tbc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			tpx: Pixel{
				108, // x2
				125, // y2
			},
		},
		// Case 001, move -19 along x and 0 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			tbc: Bucket{
				150, // x0
				130, // y0
				106, // x1
				101, // y1
			},
			tpx: Pixel{
				156, // x2
				129, // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != tc.tbc[X0] {
				t.Fatal("expected", tc.tbc[X0], "got", tbc[X0])
			}
			if tbc[Y0] != tc.tbc[Y0] {
				t.Fatal("expected", tc.tbc[Y0], "got", tbc[Y0])
			}

			if tbc[X1] != tc.tbc[X1] {
				t.Fatal("expected", tc.tbc[X1], "got", tbc[X1])
			}
			if tbc[Y1] != tc.tbc[Y1] {
				t.Fatal("expected", tc.tbc[Y1], "got", tbc[Y1])
			}

			if tpx[X2] != tc.tpx[X2] {
				t.Fatal("expected", tc.tpx[X2], "got", tpx[X2])
			}
			if tpx[Y2] != tc.tpx[Y2] {
				t.Fatal("expected", tc.tpx[Y2], "got", tpx[Y2])
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Matrix_Target_Quadrant_4(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		mil float64
		tbc Bucket
		tpx Pixel
	}{
		// Case 000, move -4 along x and +3 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(4),   // quadrant 4
				byte(108), // 38.12° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			tbc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			tpx: Pixel{
				107, // x2
				132, // y2
			},
		},
		// Case 001, move 0 along x and +19 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			tbc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			tpx: Pixel{
				111, // x2
				148, // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != tc.tbc[X0] {
				t.Fatal("expected", tc.tbc[X0], "got", tbc[X0])
			}
			if tbc[Y0] != tc.tbc[Y0] {
				t.Fatal("expected", tc.tbc[Y0], "got", tbc[Y0])
			}

			if tbc[X1] != tc.tbc[X1] {
				t.Fatal("expected", tc.tbc[X1], "got", tbc[X1])
			}
			if tbc[Y1] != tc.tbc[Y1] {
				t.Fatal("expected", tc.tbc[Y1], "got", tbc[Y1])
			}

			if tpx[X2] != tc.tpx[X2] {
				t.Fatal("expected", tc.tpx[X2], "got", tpx[X2])
			}
			if tpx[Y2] != tc.tpx[Y2] {
				t.Fatal("expected", tc.tpx[Y2], "got", tpx[Y2])
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Matrix_Target_Overflow_top(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		mil float64
		ovr byte
	}{
		// Case 000, move +1920 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 270°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			ovr: byte('o'),
		},
		// Case 001, move -4 along x and +3 along y
		{
			obc: Bucket{
				150, // x0
				163, // y0
				107, // x1
				163, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1), // quadrant 1
				byte(2), // 0.71° from 0°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('t'),
		},
		// Case 002,
		{
			obc: Bucket{
				150, // x0
				179, // y0, out of bounds
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1), // quadrant 1
				byte(2), // 0.71° from 0°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('t'),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != 0 {
				t.Fatal("expected", 0, "got", tbc[X0])
			}
			if tbc[Y0] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y0])
			}

			if tbc[X1] != 0 {
				t.Fatal("expected", 0, "got", tbc[X1])
			}
			if tbc[Y1] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y1])
			}

			if tpx[X2] != 0 {
				t.Fatal("expected", 0, "got", tpx[X2])
			}
			if tpx[Y2] != 0 {
				t.Fatal("expected", 0, "got", tpx[Y2])
			}

			if ovr != tc.ovr {
				t.Fatal("expected", string(tc.ovr), "got", string(ovr))
			}
		})
	}
}

func Test_Matrix_Target_Overflow_right(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		mil float64
		ovr byte
	}{
		// Case 000, move +1920 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			ovr: byte('o'),
		},
		// Case 001, move -4 along x and +3 along y
		{
			obc: Bucket{
				163, // x0
				150, // y0
				163, // x1
				107, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('r'),
		},
		// Case 002,
		{
			obc: Bucket{
				179, // x0, out of bounds
				150, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(2), // quadrant 2
				byte(2), // 0.71° from 90°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('r'),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != 0 {
				t.Fatal("expected", 0, "got", tbc[X0])
			}
			if tbc[Y0] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y0])
			}

			if tbc[X1] != 0 {
				t.Fatal("expected", 0, "got", tbc[X1])
			}
			if tbc[Y1] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y1])
			}

			if tpx[X2] != 0 {
				t.Fatal("expected", 0, "got", tpx[X2])
			}
			if tpx[Y2] != 0 {
				t.Fatal("expected", 0, "got", tpx[Y2])
			}

			if ovr != tc.ovr {
				t.Fatal("expected", string(tc.ovr), "got", string(ovr))
			}
		})
	}
}

func Test_Matrix_Target_Overflow_bottom(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		mil float64
		ovr byte
	}{
		// Case 000, move +1920 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			ovr: byte('o'),
		},
		// Case 001, move -4 along x and +3 along y
		{
			obc: Bucket{
				150, // x0
				100, // y0
				107, // x1
				100, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('b'),
		},
		// Case 002,
		{
			obc: Bucket{
				150, // x0
				27,  // y0, out of bounds
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(2),   // quadrant 2
				byte(253), // 89.29° from 90°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('b'),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != 0 {
				t.Fatal("expected", 0, "got", tbc[X0])
			}
			if tbc[Y0] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y0])
			}

			if tbc[X1] != 0 {
				t.Fatal("expected", 0, "got", tbc[X1])
			}
			if tbc[Y1] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y1])
			}

			if tpx[X2] != 0 {
				t.Fatal("expected", 0, "got", tpx[X2])
			}
			if tpx[Y2] != 0 {
				t.Fatal("expected", 0, "got", tpx[Y2])
			}

			if ovr != tc.ovr {
				t.Fatal("expected", string(tc.ovr), "got", string(ovr))
			}
		})
	}
}

func Test_Matrix_Target_Overflow_left(t *testing.T) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
		mil float64
		ovr byte
	}{
		// Case 000, move +1920 along y
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			ovr: byte('o'),
		},
		// Case 001, move -4 along x and +3 along y
		{
			obc: Bucket{
				100, // x0
				150, // y0
				100, // x1
				107, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('l'),
		},
		// Case 002,
		{
			obc: Bucket{
				27,  // x0, out of bounds
				150, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(4), // quadrant 4
				byte(2), // 0.71° from 270°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
			ovr: byte('l'),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tbc, tpx, ovr := Target(tc.obc, tc.opx, tc.spc, tc.tim)

			if tbc[X0] != 0 {
				t.Fatal("expected", 0, "got", tbc[X0])
			}
			if tbc[Y0] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y0])
			}

			if tbc[X1] != 0 {
				t.Fatal("expected", 0, "got", tbc[X1])
			}
			if tbc[Y1] != 0 {
				t.Fatal("expected", 0, "got", tbc[Y1])
			}

			if tpx[X2] != 0 {
				t.Fatal("expected", 0, "got", tpx[X2])
			}
			if tpx[Y2] != 0 {
				t.Fatal("expected", 0, "got", tpx[Y2])
			}

			if ovr != tc.ovr {
				t.Fatal("expected", string(tc.ovr), "got", string(ovr))
			}
		})
	}
}

func Benchmark_Matrix_Target(b *testing.B) {
	testCases := []struct {
		obc Bucket
		opx Pixel
		spc Space
		tim [2]byte
	}{
		// Case 000, ~3.00 ns/op
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1),   // quadrant 1
				byte(108), // 38.12° from 0°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
		},
		// Case 001, ~3.00 ns/op
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(2),   // quadrant 2
				byte(253), // 89.29° from 90°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
		},
		// Case 002, ~3.20 ns/op
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
		},
		// Case 003, ~3.00 ns/op
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(4),   // quadrant 4
				byte(108), // 38.12° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
		},
		// Case 004, ~1.80 ns/op
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 270°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 005, ~2.80 ns/op
		{
			obc: Bucket{
				150, // x0
				179, // y0, out of bounds
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1), // quadrant 1
				byte(2), // 0.71° from 0°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
		},
		// Case 006, ~2.70 ns/op
		{
			obc: Bucket{
				163, // x0
				150, // y0
				163, // x1
				107, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
		},
		// Case 007, ~1.80 ns/op
		{
			obc: Bucket{
				150, // x0
				130, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 008, ~2.80 ns/op
		{
			obc: Bucket{
				150, // x0
				100, // y0
				107, // x1
				100, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
		},
		// Case 009, ~2.80 ns/op
		{
			obc: Bucket{
				100, // x0
				150, // y0
				100, // x1
				107, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
		},
		// Case 010, ~2.80 ns/op
		{
			obc: Bucket{
				27,  // x0, out of bounds
				150, // y0
				107, // x1
				101, // y1
			},
			opx: Pixel{
				111, // x2
				129, // y2
			},
			spc: Space{
				byte(4), // quadrant 4
				byte(2), // 0.71° from 270°
			},
			tim: [2]byte{
				byte(100), // 4 standard frames
				byte(4),   // 400% speed
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Target(tc.obc, tc.opx, tc.spc, tc.tim)
			}
		})
	}
}
