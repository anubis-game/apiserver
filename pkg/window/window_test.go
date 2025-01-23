package window

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Window_Bytes(t *testing.T) {
	testCases := []struct {
		w *Window
		b []byte
	}{
		// Case 000
		{
			w: New(Config{
				Bck: matrix.Bucket{115, 123, 107, 119},
				Pxl: matrix.Pixel{124, 125},
				Spc: matrix.Space{2, 207},
			}),
			b: []byte{0x73, 0x7b, 0x6b, 0x77, 0x73, 0x7b, 0x70, 0x7c, 0x73, 0x7b, 0x75, 0x81, 0x7c, 0x7d, 0x2, 0xcf},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.w.Bytes()
			w := FromBytes(b)

			if *tc.w != *w {
				t.Fatal("expected", tc.w, "got", w)
			}

			if !bytes.Equal(b, tc.b) {
				t.Fatal("expected", tc.b, "got", b)
			}
		})
	}
}

func Test_Window_Centre(t *testing.T) {
	testCases := []struct {
		w *Window
		c matrix.Bucket
	}{
		// Case 000
		{
			w: New(Config{
				Bck: matrix.Bucket{100, 100, 100, 100},
				Pxl: matrix.Pixel{124, 125},
				Spc: matrix.Space{2, 207},
			}),
			c: matrix.Bucket{100, 100, 105, 105},
		},
		// Case 000
		{
			w: New(Config{
				Bck: matrix.Bucket{115, 123, 107, 119},
				Pxl: matrix.Pixel{124, 125},
				Spc: matrix.Space{2, 207},
			}),
			c: matrix.Bucket{115, 123, 112, 124},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			c := tc.w.wcn

			if tc.c != c {
				t.Fatal("expected", tc.c, "got", c)
			}
		})
	}
}

func Benchmark_Window_Bytes(b *testing.B) {
	testCases := []struct {
		w *Window
	}{
		// Case 000 ~1.30 ns/op
		{
			w: New(Config{
				Bck: matrix.Bucket{115, 123, 107, 119},
				Pxl: matrix.Pixel{124, 125},
				Spc: matrix.Space{2, 207},
			}),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tc.w.Bytes()
			}
		})
	}
}
