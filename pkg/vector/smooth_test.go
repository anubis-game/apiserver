package vector

import (
	"fmt"
	"strings"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

// https://www.desmos.com/calculator/gewoldajer
func Test_Vector_Smooth(t *testing.T) {
	var vec *Vector
	{
		vec = tesVec()
	}

	{
		tesUpd(vec)
	}

	for i := range 100 {
		if i <= 5 || i == 99 {
			var lis []string
			vec.ranger(func(obj matrix.Coordinate) {
				lis = append(lis, fmt.Sprintf("(%d,%d)", obj.X, obj.Y))
			})
			fmt.Printf("%03d\n", i)
			fmt.Printf("\n")
			fmt.Printf("%s\n", strings.Join(lis, ",")) // those coordinates can be put into Desmos
			fmt.Printf("\n")
		}

		{
			vec.smooth()
		}
	}

	// TODO:test ensure we reach some point on the curve.
}

func Test_Vector_smooth(t *testing.T) {
	testCases := []struct {
		lef matrix.Coordinate
		mid matrix.Coordinate
		rig matrix.Coordinate
		apx matrix.Coordinate
	}{
		// Case 000
		{
			lef: matrix.Coordinate{X: 0, Y: 0},
			mid: matrix.Coordinate{X: 5, Y: 10}, // y-5
			rig: matrix.Coordinate{X: 10, Y: 0},
			apx: matrix.Coordinate{X: 5, Y: 5},
		},
		// Case 001
		{
			lef: matrix.Coordinate{X: 0, Y: 0},
			mid: matrix.Coordinate{X: 5, Y: 5}, // y-3
			rig: matrix.Coordinate{X: 10, Y: 0},
			apx: matrix.Coordinate{X: 5, Y: 2},
		},
		// Case 002
		{
			lef: matrix.Coordinate{X: 0, Y: 0},
			mid: matrix.Coordinate{X: 5, Y: 2}, // y-1
			rig: matrix.Coordinate{X: 10, Y: 0},
			apx: matrix.Coordinate{X: 5, Y: 1},
		},
		// Case 003
		{
			lef: matrix.Coordinate{X: 0, Y: 0},
			mid: matrix.Coordinate{X: 5, Y: 1}, // y-1
			rig: matrix.Coordinate{X: 10, Y: 0},
			apx: matrix.Coordinate{X: 5, Y: 0},
		},
		// Case 004, can't reduce below zero
		{
			lef: matrix.Coordinate{X: 0, Y: 0},
			mid: matrix.Coordinate{X: 5, Y: 0},
			rig: matrix.Coordinate{X: 10, Y: 0},
			apx: matrix.Coordinate{X: 5, Y: 0},
		},
		// Case 005
		{
			lef: matrix.Coordinate{X: 0, Y: 20},
			mid: matrix.Coordinate{X: 5, Y: 10}, // y+5
			rig: matrix.Coordinate{X: 10, Y: 20},
			apx: matrix.Coordinate{X: 5, Y: 15},
		},
		// Case 006
		{
			lef: matrix.Coordinate{X: 0, Y: 25},
			mid: matrix.Coordinate{X: 5, Y: 25},
			rig: matrix.Coordinate{X: 10, Y: 25},
			apx: matrix.Coordinate{X: 5, Y: 25},
		},
		// Case 007
		{
			lef: matrix.Coordinate{X: 04, Y: 12},
			mid: matrix.Coordinate{X: 05, Y: 12}, // x+1
			rig: matrix.Coordinate{X: 11, Y: 12},
			apx: matrix.Coordinate{X: 06, Y: 12},
		},
		// Case 008
		{
			lef: matrix.Coordinate{X: 4, Y: 2},
			mid: matrix.Coordinate{X: 5, Y: 12}, // x+1 y-5
			rig: matrix.Coordinate{X: 11, Y: 2},
			apx: matrix.Coordinate{X: 6, Y: 7},
		},
		// Case 009, https://www.desmos.com/calculator/bg2iwrsgpx
		{
			lef: matrix.Coordinate{X: 1005, Y: 1163},
			mid: matrix.Coordinate{X: 1230, Y: 1296}, // x-78 y-86
			rig: matrix.Coordinate{X: 1145, Y: 1087},
			apx: matrix.Coordinate{X: 1152, Y: 1210},
		},
		// Case 010, https://www.desmos.com/calculator/sktglbjmcd
		{
			lef: matrix.Coordinate{X: 1005, Y: 1163},
			mid: matrix.Coordinate{X: 1145, Y: 1087}, // x-14 y+71
			rig: matrix.Coordinate{X: 1230, Y: 1296},
			apx: matrix.Coordinate{X: 1131, Y: 1158},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			smx, smy := smooth(tc.lef, tc.mid, tc.rig)
			apx := matrix.Coordinate{X: smx, Y: smy}

			if apx != tc.apx {
				t.Fatalf("expected %#v got %#v", tc.apx, apx)
			}
		})
	}
}

// ~840 ns/op
func Benchmark_Vector_Smooth(b *testing.B) {
	var vec *Vector
	{
		vec = tesVec()
	}

	{
		tesUpd(vec)
	}

	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		for b.Loop() {
			vec.smooth()
		}
	})
}

func Benchmark_Vector_smooth(b *testing.B) {
	testCases := []struct {
		lef matrix.Coordinate
		mid matrix.Coordinate
		rig matrix.Coordinate
		smx int
		smy int
	}{
		// Case 000, ~39 ns/op
		{
			lef: matrix.Coordinate{X: 00, Y: 00},
			mid: matrix.Coordinate{X: 05, Y: 10},
			rig: matrix.Coordinate{X: 10, Y: 00},
		},
		// Case 001, ~39 ns/op
		{
			lef: matrix.Coordinate{X: 00, Y: 20},
			mid: matrix.Coordinate{X: 05, Y: 10},
			rig: matrix.Coordinate{X: 10, Y: 20},
		},
		// Case 002, ~39 ns/op
		{
			lef: matrix.Coordinate{X: 00, Y: 00},
			mid: matrix.Coordinate{X: 05, Y: 00},
			rig: matrix.Coordinate{X: 10, Y: 00},
		},
		// Case 003, ~39 ns/op
		{
			lef: matrix.Coordinate{X: 04, Y: 12},
			mid: matrix.Coordinate{X: 05, Y: 12},
			rig: matrix.Coordinate{X: 11, Y: 12},
		},
		// Case 004, ~39 ns/op
		{
			lef: matrix.Coordinate{X: 04, Y: 02},
			mid: matrix.Coordinate{X: 05, Y: 12},
			rig: matrix.Coordinate{X: 11, Y: 02},
		},
		// Case 005, ~39 ns/op
		{
			lef: matrix.Coordinate{X: 5, Y: 63},
			mid: matrix.Coordinate{X: 30, Y: 96},
			rig: matrix.Coordinate{X: 45, Y: 87},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				smooth(tc.lef, tc.mid, tc.rig)
			}
		})
	}
}
