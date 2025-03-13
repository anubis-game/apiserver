package engine

import (
	"fmt"
	"slices"
	"testing"
)

func Benchmark_Engine_allpt8_map(b *testing.B) {
	testCases := []struct {
		lis []byte
	}{
		// Case 000, len 10, ~113 ns/op, 1 allocs/op
		{
			lis: []byte{
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
			},
		},
		// Case 001, len 100, ~785 ns/op, 1 allocs/op
		{
			lis: []byte{
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				all := map[byte]struct{}{}
				for _, x := range tc.lis {
					all[x] = struct{}{}
				}
			}
		})
	}
}

func Benchmark_Engine_allpt8_slice(b *testing.B) {
	testCases := []struct {
		lis []byte
	}{
		// Case 000, len 10, ~51 ns/op, 1 allocs/op
		{
			lis: []byte{
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
			},
		},
		// Case 001, len 100, ~315 ns/op, 1 allocs/op
		{
			lis: []byte{
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
				1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				all := []byte{}
				for _, x := range tc.lis {
					if !slices.Contains(all, x) {
						all = append(all, x)
					}
				}
			}
		})
	}
}
