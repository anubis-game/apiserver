package energy

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Energy_Bytes(t *testing.T) {
	testCases := []struct {
		e Energy
		b []byte
	}{
		// Case 000
		{
			e: Energy{
				Bck: matrix.Bucket{115, 123, 107, 119},
				Pxl: matrix.Pixel{124, 125},
				Pro: matrix.Profile{15, 0},
			},
			b: []byte{0x73, 0x7b, 0x6b, 0x77, 0x7c, 0x7d, 0xf, 0x0},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.e.Bytes()
			e := FromBytes(b)

			if tc.e != e {
				t.Fatal("expected", tc.e, "got", e)
			}

			if !bytes.Equal(b, tc.b) {
				t.Fatal("expected", tc.b, "got", b)
			}
		})
	}
}

func Benchmark_Energy_Bytes(b *testing.B) {
	testCases := []struct {
		e Energy
	}{
		// Case 000 ~0.30 ns/op
		{
			e: Energy{
				Bck: matrix.Bucket{115, 123, 107, 119},
				Pxl: matrix.Pixel{124, 125},
				Pro: matrix.Profile{15, 0},
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tc.e.Bytes()
			}
		})
	}
}
