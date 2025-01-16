package coordinate

import (
	"fmt"
	"testing"
)

func Test_Coordinate_Next_Quadrant_1(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		nxt [6]byte
	}{
		// Case 000, move +3 along x and +4 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant 1
				byte(108), // 38.12° from 0°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(2),  // y1
				byte(14), // x2
				byte(1),  // y2
			},
		},
		// Case 001, move +19 along x and 0 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(30), // x2
				byte(29), // y2
			},
		},
		// Case 002, move +1920 along x and +24 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(17), // x0
				byte(23), // y0
				byte(3),  // x1
				byte(2),  // y1
				byte(11), // x2
				byte(21), // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Coordinate_Next_Quadrant_2(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
		nxt [6]byte
	}{
		// Case 000, move +4 along x and -3 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(2),   // quadrant 2
				byte(108), // 38.12° from 90°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(15), // x2
				byte(26), // y2
			},
		},
		// Case 001, move 0 along x and -19 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(2),   // quadrant 2
				byte(253), // 89.29° from 90°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(10), // y2
			},
		},
		// Case 002, move +1920 along x and -24 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(2), // quadrant 2
				byte(2), // 0.71° from 90°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(17), // x0
				byte(23), // y0
				byte(3),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(5),  // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Coordinate_Next_Quadrant_3(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
		nxt [6]byte
	}{
		// Case 000, move -3 along x and -4 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(3),   // quadrant 3
				byte(108), // 38.12° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(8),  // x2
				byte(25), // y2
			},
		},
		// Case 001, move -19 along x and 0 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(6),  // x1
				byte(1),  // y1
				byte(24), // x2
				byte(29), // y2
			},
		},
		// Case 002, move -22 along x and -1766 along y. This test covers an edge
		// case where y2 may not lead to the necessary boundary violation check,
		// even if y0 and y1 are at 0 already.
		{
			cur: [6]byte{
				byte(23), // x0
				byte(1),  // y0
				byte(1),  // x1
				byte(23), // y1
				byte(29), // x2
				byte(20), // y2
			},
			spc: [2]byte{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(46),  // 4600% speed
			},
			nxt: [6]byte{
				byte(23), // x0
				byte(0),  // y0
				byte(1),  // x1
				byte(0),  // y1
				byte(7),  // x2
				byte(14), // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Coordinate_Next_Quadrant_4(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
		nxt [6]byte
	}{
		// Case 000, move -4 along x and +3 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(4),   // quadrant 4
				byte(108), // 38.12° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(2),  // y1
				byte(7),  // x2
				byte(0),  // y2
			},
		},
		// Case 001, move 0 along x and +19 along y
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
			nxt: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(2),  // y1
				byte(11), // x2
				byte(16), // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != 0x00 {
				t.Fatal("expected", 0, "got", string(ovr))
			}
		})
	}
}

func Test_Coordinate_Next_Overflow_top(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
		nxt [6]byte
	}{
		// Case 000, move +24 along x and +1920 along y
		{
			cur: [6]byte{
				byte(23), // x0
				byte(30), // y0
				byte(1),  // x1
				byte(7),  // y1
				byte(29), // x2
				byte(11), // y2
			},
			spc: [2]byte{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 270°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 001, move +22 along x and +1805 along y. This test covers an edge
		// case where y2 may not lead to the necessary boundary violation check,
		// even if y0 and y1 are at 31 already.
		{
			cur: [6]byte{
				byte(23), // x0
				byte(30), // y0
				byte(1),  // x1
				byte(7),  // y1
				byte(29), // x2
				byte(11), // y2
			},
			spc: [2]byte{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 270°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(47),  // 4700% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 002, move +24 along x and +1920 along y
		{
			cur: [6]byte{
				byte(23), // x0
				byte(30), // y0
				byte(1),  // x1
				byte(7),  // y1
				byte(29), // x2
				byte(11), // y2
			},
			spc: [2]byte{
				byte(1), // quadrant 1
				byte(2), // 0.71° from 0°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != byte('t') {
				t.Fatal("expected", "t", "got", string(ovr))
			}
		})
	}
}

func Test_Coordinate_Next_Overflow_right(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
		nxt [6]byte
	}{
		// Case 000, move +1920 along x and +24 along y
		{
			cur: [6]byte{
				byte(30), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 001, move +1805 along x and +22 along y. This test covers an edge
		// case where x2 may not lead to the necessary boundary violation check,
		// even if x0 and x1 are at 31 already.
		{
			cur: [6]byte{
				byte(30), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(47),  // 4700% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 002, move +1920 along x and -24 along y
		{
			cur: [6]byte{
				byte(30), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(2), // quadrant 2
				byte(2), // 0.71° from 90°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != byte('r') {
				t.Fatal("expected", "r", "got", string(ovr))
			}
		})
	}
}

func Test_Coordinate_Next_Overflow_bottom(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
		nxt [6]byte
	}{
		// Case 000, move -24 along x and -1920 along y
		{
			cur: [6]byte{
				byte(23), // x0
				byte(1),  // y0
				byte(1),  // x1
				byte(23), // y1
				byte(29), // x2
				byte(20), // y2
			},
			spc: [2]byte{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 001, move -22 along x and -1805 along y. This test covers an edge
		// case where y2 may not lead to the necessary boundary violation check,
		// even if y0 and y1 are at 0 already.
		{
			cur: [6]byte{
				byte(23), // x0
				byte(1),  // y0
				byte(1),  // x1
				byte(23), // y1
				byte(29), // x2
				byte(20), // y2
			},
			spc: [2]byte{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(47),  // 4700% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 002, move +24 along x and +1920 along y
		{
			cur: [6]byte{
				byte(23), // x0
				byte(1),  // y0
				byte(1),  // x1
				byte(23), // y1
				byte(29), // x2
				byte(20), // y2
			},
			spc: [2]byte{
				byte(2),   // quadrant 2
				byte(253), // 89.29° from 90°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != byte('b') {
				t.Fatal("expected", "b", "got", string(ovr))
			}
		})
	}
}

func Test_Coordinate_Next_Overflow_left(t *testing.T) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
		nxt [6]byte
	}{
		// Case 000, move -1920 along x and -24 along y
		{
			cur: [6]byte{
				byte(1),  // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(20), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 001, move -1805 along x and -22 along y. This test covers an edge
		// case where x2 may not lead to the necessary boundary violation check,
		// even if x0 and x1 are at 0 already.
		{
			cur: [6]byte{
				byte(1),  // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(20), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(47),  // 4700% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
		// Case 002, move -1920 along x and +24 along y
		{
			cur: [6]byte{
				byte(1),  // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(20), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(4), // quadrant 4
				byte(2), // 0.71° from 270°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
			nxt: [6]byte{
				byte(0), // x0
				byte(0), // y0
				byte(0), // x1
				byte(0), // y1
				byte(0), // x2
				byte(0), // y2
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			nxt, ovr := Next(tc.cur, tc.spc, tc.tim)

			if nxt[0] != tc.nxt[0] {
				t.Fatal("expected", tc.nxt[0], "got", nxt[0]) // x0
			}
			if nxt[1] != tc.nxt[1] {
				t.Fatal("expected", tc.nxt[1], "got", nxt[1]) // y0
			}

			if nxt[2] != tc.nxt[2] {
				t.Fatal("expected", tc.nxt[2], "got", nxt[2]) // x1
			}
			if nxt[3] != tc.nxt[3] {
				t.Fatal("expected", tc.nxt[3], "got", nxt[3]) // y1
			}

			if nxt[4] != tc.nxt[4] {
				t.Fatal("expected", tc.nxt[4], "got", nxt[4]) // x2
			}
			if nxt[5] != tc.nxt[5] {
				t.Fatal("expected", tc.nxt[5], "got", nxt[5]) // y2
			}

			if ovr != byte('l') {
				t.Fatal("expected", "l", "got", string(ovr))
			}
		})
	}
}

func Benchmark_Coordinate_Next(b *testing.B) {
	testCases := []struct {
		cur [6]byte
		spc [2]byte
		tim [2]byte
		mil float64
	}{
		// Case 000, ~4.50 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant
				byte(108), // 38.12° from 0°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
		},
		// Case 001, ~3.90 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
		},
		// Case 002, ~3.90 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(2),   // quadrant 2
				byte(108), // 38.12° from 90°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
		},
		// Case 003, ~3.90 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(2),   // quadrant 2
				byte(253), // 89.29° from 90°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
		},
		// Case 004, ~3.95 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(3),   // quadrant 3
				byte(108), // 38.12° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
		},
		// Case 005, ~4.50 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
		},
		// Case 006, ~4.50 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(4),   // quadrant 4
				byte(108), // 38.12° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(1),  // 100% speed
			},
		},
		// Case 007, ~4.60 ns/op
		{
			cur: [6]byte{
				byte(15), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(25), // standard frame
				byte(4),  // 400% speed
			},
		},
		// Case 008, ~57.00 ns/op
		{
			cur: [6]byte{
				byte(23), // x0
				byte(30), // y0
				byte(1),  // x1
				byte(7),  // y1
				byte(29), // x2
				byte(11), // y2
			},
			spc: [2]byte{
				byte(4),   // quadrant 4
				byte(253), // 89.29° from 270°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 009, ~57.80 ns/op
		{
			cur: [6]byte{
				byte(23), // x0
				byte(30), // y0
				byte(1),  // x1
				byte(7),  // y1
				byte(29), // x2
				byte(11), // y2
			},
			spc: [2]byte{
				byte(1), // quadrant 1
				byte(2), // 0.71° from 0°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 010, ~56.80 ns/op
		{
			cur: [6]byte{
				byte(30), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(1),   // quadrant 1
				byte(253), // 89.29° from 0°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 011, ~57.10 ns/op
		{
			cur: [6]byte{
				byte(30), // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(11), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(2), // quadrant 2
				byte(2), // 0.71° from 90°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 012, ~57.20 ns/op
		{
			cur: [6]byte{
				byte(23), // x0
				byte(1),  // y0
				byte(1),  // x1
				byte(23), // y1
				byte(29), // x2
				byte(20), // y2
			},
			spc: [2]byte{
				byte(3), // quadrant 3
				byte(2), // 0.71° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 013, ~57.40 ns/op
		{
			cur: [6]byte{
				byte(23), // x0
				byte(1),  // y0
				byte(1),  // x1
				byte(23), // y1
				byte(29), // x2
				byte(20), // y2
			},
			spc: [2]byte{
				byte(2),   // quadrant 2
				byte(253), // 89.29° from 90°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 014, ~42.80 ns/op
		{
			cur: [6]byte{
				byte(1),  // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(20), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(3),   // quadrant 3
				byte(253), // 89.29° from 180°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
		// Case 015, ~42.40 ns/op
		{
			cur: [6]byte{
				byte(1),  // x0
				byte(23), // y0
				byte(7),  // x1
				byte(1),  // y1
				byte(20), // x2
				byte(29), // y2
			},
			spc: [2]byte{
				byte(4), // quadrant 4
				byte(2), // 0.71° from 270°
			},
			tim: [2]byte{
				byte(200), // 8 standard frames
				byte(50),  // 5000% speed
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Next(tc.cur, tc.spc, tc.tim)
			}
		})
	}
}
