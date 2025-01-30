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

func Test_Matrix_Bucket_Ngh_Quadrant_1(t *testing.T) {
	testCases := []struct {
		bck Bucket
		ngh []Bucket
	}{
		// Case 000
		{
			bck: Bucket{100, 100, 100, 100},
			ngh: []Bucket{
				{100, 100, 100, 101}, {100, 100, 101, 101},
				{100, 100, 100, 100}, {100, 100, 101, 100},
				{}, {},
			},
		},
		// Case 001
		{
			bck: Bucket{100, 100, 101, 100},
			ngh: []Bucket{
				{100, 100, 101, 101}, {100, 100, 102, 101},
				{100, 100, 101, 100}, {100, 100, 102, 100},
				{}, {},
			},
		},
		// Case 002
		{
			bck: Bucket{115, 123, 107, 101},
			ngh: []Bucket{
				{115, 123, 107, 102}, {115, 123, 108, 102},
				{115, 123, 107, 101}, {115, 123, 108, 101},
				{}, {},
			},
		},
		// Case 003
		{
			bck: Bucket{115, 122, 107, 163},
			ngh: []Bucket{
				{115, 123, 107, 100}, {115, 123, 108, 100},
				{115, 122, 107, 163}, {115, 122, 108, 163},
				{}, {},
			},
		},
		// Case 004
		{
			bck: Bucket{115, 123, 107, 100},
			ngh: []Bucket{
				{115, 123, 107, 101}, {115, 123, 108, 101},
				{115, 123, 107, 100}, {115, 123, 108, 100},
				{}, {},
			},
		},
		// Case 005
		{
			bck: Bucket{115, 123, 163, 163},
			ngh: []Bucket{
				{115, 124, 163, 100}, {116, 124, 100, 100},
				{115, 123, 163, 163}, {116, 123, 100, 163},
				{}, {},
			},
		},
		// Case 006
		{
			bck: Bucket{115, 123, 100, 100},
			ngh: []Bucket{
				{115, 123, 100, 101}, {115, 123, 101, 101},
				{115, 123, 100, 100}, {115, 123, 101, 100},
				{}, {},
			},
		},
		// Case 007
		{
			bck: Bucket{114, 122, 163, 163},
			ngh: []Bucket{
				{114, 123, 163, 100}, {115, 123, 100, 100},
				{114, 122, 163, 163}, {115, 122, 100, 163},
				{}, {},
			},
		},
		// Case 008
		{
			bck: Bucket{162, 162, 163, 163},
			ngh: []Bucket{
				{162, 163, 163, 100}, {163, 163, 100, 100},
				{162, 162, 163, 163}, {163, 162, 100, 163},
				{}, {},
			},
		},
		// Case 009
		{
			bck: Bucket{163, 163, 163, 163},
			ngh: []Bucket{
				{163, 164, 163, 100}, {164, 164, 100, 100},
				{163, 163, 163, 163}, {164, 163, 100, 163},
				{}, {},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ngh := tc.bck.Ngh(0x01)

			if !slices.Equal(ngh[:], tc.ngh) {
				t.Fatal("expected", tc.ngh, "got", ngh)
			}
		})
	}
}

func Test_Matrix_Bucket_Ngh_Quadrant_2(t *testing.T) {
	testCases := []struct {
		bck Bucket
		ngh []Bucket
	}{
		// Case 000
		{
			bck: Bucket{100, 100, 100, 100},
			ngh: []Bucket{
				{100, 100, 100, 100}, {100, 100, 101, 100},
				{100, 99, 100, 163}, {100, 99, 101, 163},
				{}, {},
			},
		},
		// Case 001
		{
			bck: Bucket{100, 100, 101, 100},
			ngh: []Bucket{
				{100, 100, 101, 100}, {100, 100, 102, 100},
				{100, 99, 101, 163}, {100, 99, 102, 163},
				{}, {},
			},
		},
		// Case 002
		{
			bck: Bucket{115, 123, 107, 101},
			ngh: []Bucket{
				{115, 123, 107, 101}, {115, 123, 108, 101},
				{115, 123, 107, 100}, {115, 123, 108, 100},
				{}, {},
			},
		},
		// Case 003
		{
			bck: Bucket{115, 122, 107, 163},
			ngh: []Bucket{
				{115, 122, 107, 163}, {115, 122, 108, 163},
				{115, 122, 107, 162}, {115, 122, 108, 162},
				{}, {},
			},
		},
		// Case 004
		{
			bck: Bucket{115, 123, 107, 100},
			ngh: []Bucket{
				{115, 123, 107, 100}, {115, 123, 108, 100},
				{115, 122, 107, 163}, {115, 122, 108, 163},
				{}, {},
			},
		},
		// Case 005
		{
			bck: Bucket{115, 123, 163, 163},
			ngh: []Bucket{
				{115, 123, 163, 163}, {116, 123, 100, 163},
				{115, 123, 163, 162}, {116, 123, 100, 162},
				{}, {},
			},
		},
		// Case 006
		{
			bck: Bucket{115, 123, 100, 100},
			ngh: []Bucket{
				{115, 123, 100, 100}, {115, 123, 101, 100},
				{115, 122, 100, 163}, {115, 122, 101, 163},
				{}, {},
			},
		},
		// Case 007
		{
			bck: Bucket{115, 123, 100, 100},
			ngh: []Bucket{
				{115, 123, 100, 100}, {115, 123, 101, 100},
				{115, 122, 100, 163}, {115, 122, 101, 163},
				{}, {},
			},
		},
		// Case 008
		{
			bck: Bucket{114, 122, 163, 163},
			ngh: []Bucket{
				{114, 122, 163, 163}, {115, 122, 100, 163},
				{114, 122, 163, 162}, {115, 122, 100, 162},
				{}, {},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ngh := tc.bck.Ngh(0x02)

			if !slices.Equal(ngh[:], tc.ngh) {
				t.Fatal("expected", tc.ngh, "got", ngh)
			}
		})
	}
}

func Test_Matrix_Bucket_Ngh_Quadrant_3(t *testing.T) {
	testCases := []struct {
		bck Bucket
		ngh []Bucket
	}{
		// Case 000
		{
			bck: Bucket{100, 100, 100, 100},
			ngh: []Bucket{
				{99, 100, 163, 100}, {100, 100, 100, 100},
				{99, 99, 163, 163}, {100, 99, 100, 163},
				{}, {},
			},
		},
		// Case 001
		{
			bck: Bucket{100, 100, 101, 100},
			ngh: []Bucket{
				{100, 100, 100, 100}, {100, 100, 101, 100},
				{100, 99, 100, 163}, {100, 99, 101, 163},
				{}, {},
			},
		},
		// Case 002
		{
			bck: Bucket{115, 123, 107, 101},
			ngh: []Bucket{
				{115, 123, 106, 101}, {115, 123, 107, 101},
				{115, 123, 106, 100}, {115, 123, 107, 100},
				{}, {},
			},
		},
		// Case 003
		{
			bck: Bucket{115, 122, 107, 163},
			ngh: []Bucket{
				{115, 122, 106, 163}, {115, 122, 107, 163},
				{115, 122, 106, 162}, {115, 122, 107, 162},
				{}, {},
			},
		},
		// Case 004
		{
			bck: Bucket{115, 123, 107, 100},
			ngh: []Bucket{
				{115, 123, 106, 100}, {115, 123, 107, 100},
				{115, 122, 106, 163}, {115, 122, 107, 163},
				{}, {},
			},
		},
		// Case 005
		{
			bck: Bucket{115, 123, 163, 163},
			ngh: []Bucket{
				{115, 123, 162, 163}, {115, 123, 163, 163},
				{115, 123, 162, 162}, {115, 123, 163, 162},
				{}, {},
			},
		},
		// Case 006
		{
			bck: Bucket{115, 123, 100, 100},
			ngh: []Bucket{
				{114, 123, 163, 100}, {115, 123, 100, 100},
				{114, 122, 163, 163}, {115, 122, 100, 163},
				{}, {},
			},
		},
		// Case 007
		{
			bck: Bucket{115, 123, 100, 100},
			ngh: []Bucket{
				{114, 123, 163, 100}, {115, 123, 100, 100},
				{114, 122, 163, 163}, {115, 122, 100, 163},
				{}, {},
			},
		},
		// Case 008
		{
			bck: Bucket{114, 122, 163, 163},
			ngh: []Bucket{
				{114, 122, 162, 163}, {114, 122, 163, 163},
				{114, 122, 162, 162}, {114, 122, 163, 162},
				{}, {},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ngh := tc.bck.Ngh(0x03)

			if !slices.Equal(ngh[:], tc.ngh) {
				t.Fatal("expected", tc.ngh, "got", ngh)
			}
		})
	}
}

func Test_Matrix_Bucket_Ngh_Quadrant_4(t *testing.T) {
	testCases := []struct {
		bck Bucket
		ngh []Bucket
	}{
		// Case 000
		{
			bck: Bucket{100, 100, 100, 100},
			ngh: []Bucket{
				{99, 100, 163, 101}, {100, 100, 100, 101},
				{99, 100, 163, 100}, {100, 100, 100, 100},
				{}, {},
			},
		},
		// Case 001
		{
			bck: Bucket{100, 100, 101, 100},
			ngh: []Bucket{
				{100, 100, 100, 101}, {100, 100, 101, 101},
				{100, 100, 100, 100}, {100, 100, 101, 100},
				{}, {},
			},
		},
		// Case 002
		{
			bck: Bucket{115, 123, 107, 101},
			ngh: []Bucket{
				{115, 123, 106, 102}, {115, 123, 107, 102},
				{115, 123, 106, 101}, {115, 123, 107, 101},
				{}, {},
			},
		},
		// Case 003
		{
			bck: Bucket{115, 122, 107, 163},
			ngh: []Bucket{
				{115, 123, 106, 100}, {115, 123, 107, 100},
				{115, 122, 106, 163}, {115, 122, 107, 163},
				{}, {},
			},
		},
		// Case 004
		{
			bck: Bucket{115, 123, 107, 100},
			ngh: []Bucket{
				{115, 123, 106, 101}, {115, 123, 107, 101},
				{115, 123, 106, 100}, {115, 123, 107, 100},
				{}, {},
			},
		},
		// Case 005
		{
			bck: Bucket{115, 123, 163, 163},
			ngh: []Bucket{
				{115, 124, 162, 100}, {115, 124, 163, 100},
				{115, 123, 162, 163}, {115, 123, 163, 163},
				{}, {},
			},
		},
		// Case 006
		{
			bck: Bucket{115, 123, 100, 100},
			ngh: []Bucket{
				{114, 123, 163, 101}, {115, 123, 100, 101},
				{114, 123, 163, 100}, {115, 123, 100, 100},
				{}, {},
			},
		},
		// Case 007
		{
			bck: Bucket{115, 123, 100, 100},
			ngh: []Bucket{
				{114, 123, 163, 101}, {115, 123, 100, 101},
				{114, 123, 163, 100}, {115, 123, 100, 100},
				{}, {},
			},
		},
		// Case 008
		{
			bck: Bucket{114, 122, 163, 163},
			ngh: []Bucket{
				{114, 123, 162, 100}, {114, 123, 163, 100},
				{114, 122, 162, 163}, {114, 122, 163, 163},
				{}, {},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ngh := tc.bck.Ngh(0x04)

			if !slices.Equal(ngh[:], tc.ngh) {
				t.Fatal("expected", tc.ngh, "got", ngh)
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

func Benchmark_Matrix_Bucket_Ngh(b *testing.B) {
	testCases := []struct {
		bck Bucket
		qdr byte
	}{
		// Case 000, ~4.90 ns/op
		{
			bck: Bucket{100, 100, 100, 100},
			qdr: 0x01,
		},
		// Case 001, ~4.90 ns/op
		{
			bck: Bucket{100, 100, 101, 100},
			qdr: 0x02,
		},
		// Case 002, ~4.90 ns/op
		{
			bck: Bucket{115, 123, 107, 101},
			qdr: 0x03,
		},
		// Case 003, ~4.90 ns/op
		{
			bck: Bucket{115, 122, 107, 163},
			qdr: 0x04,
		},
		// Case 004, ~4.90 ns/op
		{
			bck: Bucket{115, 123, 107, 100},
			qdr: 0x01,
		},
		// Case 005, ~4.90 ns/op
		{
			bck: Bucket{115, 123, 163, 163},
			qdr: 0x02,
		},
		// Case 006, ~5.00 ns/op
		{
			bck: Bucket{115, 123, 100, 100},
			qdr: 0x03,
		},
		// Case 007, ~5.00 ns/op
		{
			bck: Bucket{115, 123, 100, 100},
			qdr: 0x04,
		},
		// Case 008, ~4.90 ns/op
		{
			bck: Bucket{114, 122, 163, 163},
			qdr: 0x01,
		},
		// Case 009, ~5.00 ns/op
		{
			bck: Bucket{114, 122, 163, 163},
			qdr: 0x02,
		},
		// Case 010, ~4.90 ns/op
		{
			bck: Bucket{162, 162, 163, 163},
			qdr: 0x03,
		},
		// Case 011, ~4.90 ns/op
		{
			bck: Bucket{163, 163, 163, 163},
			qdr: 0x04,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tc.bck.Ngh(tc.qdr)
			}
		})
	}
}
