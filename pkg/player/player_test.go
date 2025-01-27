package player

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/google/uuid"
)

func Test_Player_Bytes(t *testing.T) {
	testCases := []struct {
		p Player
		b []byte
	}{
		// Case 000
		{
			p: Player{
				Obj: matrix.Object{
					Bck: matrix.Bucket{115, 123, 107, 119},
					Pxl: matrix.Pixel{124, 125},
					Pro: matrix.Profile{15, 0},
					Uid: uuid.MustParse("f47ac10b-58cc-4372-a567-0e02b2c3d479"),
				},
				Spc: matrix.Space{1, 253},
			},
			b: []byte{
				// bucket
				0x73, 0x7b, 0x6b, 0x77,
				// pixel
				0x7c, 0x7d,
				// profile
				0xf, 0x0,
				// uuid
				0xf4, 0x7a, 0xc1, 0xb, 0x58, 0xcc, 0x43, 0x72, 0xa5, 0x67, 0xe, 0x2, 0xb2, 0xc3, 0xd4, 0x79,
				// space
				0x1, 0xfd,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.p.Bytes()
			p := FromBytes(b)

			if tc.p != p {
				t.Fatal("expected", tc.p, "got", p)
			}

			if !bytes.Equal(b, tc.b) {
				t.Fatal("expected", tc.b, "got", b)
			}
		})
	}
}

func Benchmark_Player_Bytes(b *testing.B) {
	testCases := []struct {
		p Player
	}{
		// Case 000 ~2.00 ns/op
		{
			p: Player{
				Obj: matrix.Object{
					Bck: matrix.Bucket{115, 123, 107, 119},
					Pxl: matrix.Pixel{124, 125},
					Pro: matrix.Profile{15, 0},
					Uid: uuid.MustParse("f47ac10b-58cc-4372-a567-0e02b2c3d479"),
				},
				Spc: matrix.Space{1, 253},
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tc.p.Bytes()
			}
		})
	}
}
