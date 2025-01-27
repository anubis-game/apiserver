package matrix

import (
	"fmt"
	"slices"
	"testing"
)

func Test_Matrix_Bucket_Dec(t *testing.T) {
	testCases := []struct {
		bck Bucket
		sca Bucket
	}{
		// Case 000
		{
			bck: Bucket{100, 100, 100, 100},
			sca: Bucket{99, 99, 143, 143},
		},
		// Case 001
		{
			bck: Bucket{115, 123, 107, 101},
			sca: Bucket{114, 122, 150, 144},
		},
		// Case 002
		{
			bck: Bucket{115, 122, 107, 163},
			sca: Bucket{114, 122, 150, 142},
		},
		// Case 003
		{
			bck: Bucket{115, 123, 148, 159},
			sca: Bucket{115, 123, 127, 138},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			sca := tc.bck.Dec(21)

			if !slices.Equal(sca[:], tc.sca[:]) {
				t.Fatal("expected", tc.sca, "got", sca)
			}
		})
	}
}

func Test_Matrix_Bucket_Inc(t *testing.T) {
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
			sca := tc.bck.Inc(21)

			if !slices.Equal(sca[:], tc.sca[:]) {
				t.Fatal("expected", tc.sca, "got", sca)
			}
		})
	}
}

func Benchmark_Matrix_Bucket_Dec(b *testing.B) {
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
				tc.bck.Dec(21)
			}
		})
	}
}

func Benchmark_Matrix_Bucket_Inc(b *testing.B) {
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
				tc.bck.Inc(21)
			}
		})
	}
}
