package matrix

import (
	"fmt"
	"slices"
	"testing"
)

func Test_Matrix_Bucket_Scale(t *testing.T) {
	testCases := []struct {
		bck Bucket
		sca Bucket
	}{
		// Case 000
		{
			bck: Bucket{100, 100, 100, 100},
			sca: Bucket{100, 100, 121, 121},
		},
		// Case 001
		{
			bck: Bucket{115, 123, 107, 101},
			sca: Bucket{115, 123, 128, 122},
		},
		// Case 002
		{
			bck: Bucket{115, 122, 107, 163},
			sca: Bucket{115, 123, 128, 120},
		},
		// Case 003
		{
			bck: Bucket{115, 123, 148, 159},
			sca: Bucket{116, 124, 105, 116},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			sca := tc.bck.Scale(21)

			if !slices.Equal(sca[:], tc.sca[:]) {
				t.Fatal("expected", tc.sca, "got", sca)
			}
		})
	}
}

func Benchmark_Matrix_Bucket_scaByt(b *testing.B) {
	testCases := []struct {
		bck Bucket
	}{
		// Case 000, ~0.30 ns/op
		{
			bck: Bucket{100, 100, 100, 100},
		},
		// Case 001, ~0.30 ns/op
		{
			bck: Bucket{115, 123, 107, 101},
		},
		// Case 002, ~0.30 ns/op
		{
			bck: Bucket{115, 122, 107, 163},
		},
		// Case 003, ~0.30 ns/op
		{
			bck: Bucket{115, 123, 148, 159},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tc.bck.Scale(20)
			}
		})
	}
}
