package vector

import (
	"fmt"
	"slices"
	"testing"

	"github.com/anubis-game/apiserver/pkg/generic"
	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_Circle_true(t *testing.T) {
	testCases := []struct {
		crd matrix.Coordinate
		bud int
		unt int
		lis []matrix.Coordinate
	}{
		// Case 000,
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 5,
			unt: 2,
			lis: []matrix.Coordinate{
				{X: 1000, Y: 1000},
				{X: 1000, Y: 1005},
				{X: 1004, Y: 1004},
			},
		},
		// Case 001, https://www.desmos.com/calculator/jrg8juzbqr
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 25,
			unt: 2,
			lis: []matrix.Coordinate{
				{X: 1000, Y: 1000},
				{X: 1000, Y: 1005},
				{X: 1004, Y: 1004},
				{X: 1005, Y: 1000},
				{X: 1004, Y: 996},
				{X: 1000, Y: 995},
				{X: 996, Y: 996},
				{X: 995, Y: 1000},
				{X: 996, Y: 1004},
				{X: 1000, Y: 1010},
				{X: 1007, Y: 1007},
				{X: 1010, Y: 1000},
				{X: 1007, Y: 993},
			},
		},
		// Case 002, https://www.desmos.com/calculator/oj66lhchzc
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 480,
			unt: 10,
			lis: []matrix.Coordinate{
				{X: 1000, Y: 1000},
				{X: 1000, Y: 1005},
				{X: 1004, Y: 1004},
				{X: 1005, Y: 1000},
				{X: 1004, Y: 996},
				{X: 1000, Y: 995},
				{X: 996, Y: 996},
				{X: 995, Y: 1000},
				{X: 996, Y: 1004},
				{X: 1000, Y: 1010},
				{X: 1007, Y: 1007},
				{X: 1010, Y: 1000},
				{X: 1007, Y: 993},
				{X: 1000, Y: 990},
				{X: 993, Y: 993},
				{X: 990, Y: 1000},
				{X: 993, Y: 1007},
				{X: 1000, Y: 1015},
				{X: 1011, Y: 1011},
				{X: 1015, Y: 1000},
				{X: 1011, Y: 989},
				{X: 1000, Y: 985},
				{X: 989, Y: 989},
				{X: 985, Y: 1000},
				{X: 989, Y: 1011},
				{X: 1000, Y: 1020},
				{X: 1014, Y: 1014},
				{X: 1020, Y: 1000},
				{X: 1014, Y: 986},
				{X: 1000, Y: 980},
				{X: 986, Y: 986},
				{X: 980, Y: 1000},
				{X: 986, Y: 1014},
				{X: 1000, Y: 1025},
				{X: 1018, Y: 1018},
				{X: 1025, Y: 1000},
				{X: 1018, Y: 982},
				{X: 1000, Y: 975},
				{X: 982, Y: 982},
				{X: 975, Y: 1000},
				{X: 982, Y: 1018},
				{X: 1000, Y: 1030},
				{X: 1021, Y: 1021},
				{X: 1030, Y: 1000},
				{X: 1021, Y: 979},
				{X: 1000, Y: 970},
				{X: 979, Y: 979},
				{X: 970, Y: 1000},
				{X: 979, Y: 1021},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var lis []matrix.Coordinate

			Circle(tc.crd, tc.bud, tc.unt, func(c matrix.Coordinate) bool {
				lis = append(lis, c)
				// fmt.Printf("(%d,%d),", c.X, c.Y)
				return true
			})
			// fmt.Printf("\n")

			if !slices.Equal(lis, generic.Unique(lis)) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}

func Test_Vector_Circle_false(t *testing.T) {
	testCases := []struct {
		crd matrix.Coordinate
		bud int
		unt int
		lis []matrix.Coordinate
	}{
		// Case 000,
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 5,
			unt: 2,
			lis: []matrix.Coordinate{
				{X: 1000, Y: 1000}, // false
				{X: 1000, Y: 1005}, // false
				{X: 1004, Y: 1004}, // false
				{X: 1005, Y: 1000}, // extra
				{X: 1004, Y: 996},  // extra
				{X: 1000, Y: 995},  // extra
			},
		},
		// Case 001
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 25,
			unt: 2,
			lis: []matrix.Coordinate{
				{X: 1000, Y: 1000}, // false
				{X: 1000, Y: 1005}, // false
				{X: 1004, Y: 1004}, // false
				{X: 1005, Y: 1000},
				{X: 1004, Y: 996},
				{X: 1000, Y: 995},
				{X: 996, Y: 996},
				{X: 995, Y: 1000},
				{X: 996, Y: 1004},
				{X: 1000, Y: 1010},
				{X: 1007, Y: 1007},
				{X: 1010, Y: 1000},
				{X: 1007, Y: 993},
				{X: 1000, Y: 990}, // extra
				{X: 993, Y: 993},  // extra
				{X: 990, Y: 1000}, // extra
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var lis []matrix.Coordinate

			var i int
			Circle(tc.crd, tc.bud, tc.unt, func(c matrix.Coordinate) bool {
				lis = append(lis, c)

				if i < 3 {
					i++
					return false
				}

				return true
			})

			if !slices.Equal(lis, generic.Unique(lis)) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}

func Benchmark_Vector_Circle(b *testing.B) {
	testCases := []struct {
		crd matrix.Coordinate
		bud int
		unt int
	}{
		// Case 000, 3 calls, ~11 ns/op
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 5,
			unt: 2,
		},
		// Case 001, 13 calls, ~43 ns/op
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 25,
			unt: 2,
		},
		// Case 002, 49 calls, ~160 ns/op
		{
			crd: matrix.Coordinate{X: 1000, Y: 1000},
			bud: 480,
			unt: 10,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				Circle(tc.crd, tc.bud, tc.unt, func(c matrix.Coordinate) bool {
					return true
				})
			}
		})
	}
}
