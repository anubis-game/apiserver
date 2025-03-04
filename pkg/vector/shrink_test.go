package vector

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_shrink_closer(t *testing.T) {
	testCases := []struct {
		lef matrix.Coordinate
		rig matrix.Coordinate
		red matrix.Coordinate
	}{
		// Case 000
		{
			lef: matrix.Coordinate{X: 0, Y: 0},
			rig: matrix.Coordinate{X: 10, Y: 0},
			red: matrix.Coordinate{X: 5, Y: 0},
		},
		// Case 001
		{
			lef: matrix.Coordinate{X: 0, Y: 0},
			rig: matrix.Coordinate{X: 0, Y: 10},
			red: matrix.Coordinate{X: 0, Y: 5},
		},
		// Case 002
		{
			lef: matrix.Coordinate{X: 25, Y: 0},
			rig: matrix.Coordinate{X: 0, Y: 0},
			red: matrix.Coordinate{X: 20, Y: 0},
		},
		// Case 003
		{
			lef: matrix.Coordinate{X: 0, Y: 25},
			rig: matrix.Coordinate{X: 0, Y: 0},
			red: matrix.Coordinate{X: 0, Y: 20},
		},
		// Case 004, https://www.desmos.com/calculator/ekcwxmd4tg
		{
			lef: matrix.Coordinate{X: 1002, Y: 1295},
			rig: matrix.Coordinate{X: 1014, Y: 1303},
			red: matrix.Coordinate{X: 1006, Y: 1298},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			rdx, rdy := closer(tc.lef, tc.rig)
			red := matrix.Coordinate{X: rdx, Y: rdy}

			if red != tc.red {
				t.Fatalf("expected %#v got %#v", tc.red, red)
			}
		})
	}
}
