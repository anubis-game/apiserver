package generic

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Generic_Unique_string(t *testing.T) {
	testCases := []struct {
		lis []string
		uni []string
	}{
		// Case 000
		{
			lis: []string{},
			uni: []string{},
		},
		// Case 001
		{
			lis: []string{
				"55",
				"44",
			},
			uni: []string{
				"44",
				"55",
			},
		},
		// Case 002
		{
			lis: []string{
				"33",
				"44",
				"33",
				"33",
			},
			uni: []string{
				"33",
				"44",
			},
		},
		// Case 003
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"33",
				"55",
				"66",
				"55",
				"88",
			},
			uni: []string{
				"22",
				"33",
				"44",
				"55",
				"66",
				"88",
			},
		},
		// Case 004
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
			uni: []string{
				"22",
				"33",
				"44",
				"55",
				"66",
				"88",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Unique(tc.lis)

			slices.Sort(uni)
			slices.Sort(tc.uni)

			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}

func Test_Generic_Unique_int64(t *testing.T) {
	testCases := []struct {
		lis []int64
		uni []int64
	}{
		// Case 000
		{
			lis: []int64{},
			uni: []int64{},
		},
		// Case 001
		{
			lis: []int64{
				55,
				44,
			},
			uni: []int64{
				44,
				55,
			},
		},
		// Case 002
		{
			lis: []int64{
				33,
				44,
				33,
				33,
			},
			uni: []int64{
				33,
				44,
			},
		},
		// Case 003
		{
			lis: []int64{
				33,
				44,
				88,
				22,
				33,
				55,
				66,
				55,
				88,
			},
			uni: []int64{
				22,
				33,
				44,
				55,
				66,
				88,
			},
		},
		// Case 004
		{
			lis: []int64{
				33,
				44,
				88,
				22,
				55,
				66,
			},
			uni: []int64{
				22,
				33,
				44,
				55,
				66,
				88,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Unique(tc.lis)

			slices.Sort(uni)
			slices.Sort(tc.uni)

			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}

func Test_Generic_Unique_struct(t *testing.T) {
	testCases := []struct {
		lis []matrix.Coordinate
		uni []matrix.Coordinate
	}{
		// Case 000
		{
			lis: []matrix.Coordinate{},
			uni: []matrix.Coordinate{},
		},
		// Case 001
		{
			lis: []matrix.Coordinate{
				{X: 55, Y: 55},
				{X: 44, Y: 44},
			},
			uni: []matrix.Coordinate{
				{X: 44, Y: 44},
				{X: 55, Y: 55},
			},
		},
		// Case 002
		{
			lis: []matrix.Coordinate{
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 33, Y: 33},
				{X: 33, Y: 33},
			},
			uni: []matrix.Coordinate{
				{X: 33, Y: 33},
				{X: 44, Y: 44},
			},
		},
		// Case 003
		{
			lis: []matrix.Coordinate{
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 88, Y: 88},
				{X: 22, Y: 22},
				{X: 33, Y: 33},
				{X: 55, Y: 55},
				{X: 66, Y: 66},
				{X: 55, Y: 55},
				{X: 88, Y: 88},
			},
			uni: []matrix.Coordinate{
				{X: 22, Y: 22},
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 55, Y: 55},
				{X: 66, Y: 66},
				{X: 88, Y: 88},
			},
		},
		// Case 004
		{
			lis: []matrix.Coordinate{
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 88, Y: 88},
				{X: 22, Y: 22},
				{X: 55, Y: 55},
				{X: 66, Y: 66},
			},
			uni: []matrix.Coordinate{
				{X: 22, Y: 22},
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 55, Y: 55},
				{X: 66, Y: 66},
				{X: 88, Y: 88},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Unique(tc.lis)

			sort.Sort(matrix.Coordinates(tc.uni))
			sort.Sort(matrix.Coordinates(uni))

			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}

func Benchmark_Generic_Unique_string(b *testing.B) {
	testCases := []struct {
		lis []string
	}{
		// Case 000, ~2 ns/op
		{
			lis: []string{},
		},
		// Case 001, ~15 ns/op
		{
			lis: []string{
				"55",
				"44",
			},
		},
		// Case 002, ~33 ns/op
		{
			lis: []string{
				"33",
				"44",
				"33",
				"33",
			},
		},
		// Case 003, ~154 ns/op
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"33",
				"55",
				"66",
				"55",
				"88",
			},
		},
		// Case 004, ~109 ns/op
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				Unique(tc.lis)
			}
		})
	}
}

func Benchmark_Generic_Unique_int64(b *testing.B) {
	testCases := []struct {
		lis []int64
	}{
		// Case 000, ~2 ns/op
		{
			lis: []int64{},
		},
		// Case 001, ~6 ns/op
		{
			lis: []int64{
				55,
				44,
			},
		},
		// Case 002, ~10 ns/op
		{
			lis: []int64{
				33,
				44,
				33,
				33,
			},
		},
		// Case 003, ~32 ns/op
		{
			lis: []int64{
				33,
				44,
				88,
				22,
				33,
				55,
				66,
				55,
				88,
			},
		},
		// Case 004, ~22 ns/op
		{
			lis: []int64{
				33,
				44,
				88,
				22,
				55,
				66,
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				Unique(tc.lis)
			}
		})
	}
}

func Benchmark_Generic_Unique_struct(b *testing.B) {
	testCases := []struct {
		lis []matrix.Coordinate
	}{
		// Case 000, ~2 ns/op
		{
			lis: []matrix.Coordinate{},
		},
		// Case 001, ~7 ns/op
		{
			lis: []matrix.Coordinate{
				{X: 55, Y: 55},
				{X: 44, Y: 44},
			},
		},
		// Case 002, ~15 ns/op
		{
			lis: []matrix.Coordinate{
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 33, Y: 33},
				{X: 33, Y: 33},
			},
		},
		// Case 003, ~37 ns/op
		{
			lis: []matrix.Coordinate{
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 88, Y: 88},
				{X: 22, Y: 22},
				{X: 33, Y: 33},
				{X: 55, Y: 55},
				{X: 66, Y: 66},
				{X: 55, Y: 55},
				{X: 88, Y: 88},
			},
		},
		// Case 004, ~24 ns/op
		{
			lis: []matrix.Coordinate{
				{X: 33, Y: 33},
				{X: 44, Y: 44},
				{X: 88, Y: 88},
				{X: 22, Y: 22},
				{X: 55, Y: 55},
				{X: 66, Y: 66},
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				Unique(tc.lis)
			}
		})
	}
}
