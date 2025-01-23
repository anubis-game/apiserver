package player

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/ethereum/go-ethereum/common"
)

func Test_Player_Bytes(t *testing.T) {
	testCases := []struct {
		p Player
		b []byte
	}{
		// Case 000
		{
			p: Player{
				Bck: matrix.Bucket{115, 123, 107, 119},
				Pxl: matrix.Pixel{124, 125},
				Siz: byte(15),
				Wal: common.HexToAddress("0x9632185d3851Fd06304C09BA6F1c1308189BE12b"),
			},
			b: []byte{0x96, 0x32, 0x18, 0x5d, 0x38, 0x51, 0xfd, 0x6, 0x30, 0x4c, 0x9, 0xba, 0x6f, 0x1c, 0x13, 0x8, 0x18, 0x9b, 0xe1, 0x2b, 0x73, 0x7b, 0x6b, 0x77, 0x7c, 0x7d, 0xf},
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
				Bck: matrix.Bucket{115, 123, 107, 119},
				Pxl: matrix.Pixel{124, 125},
				Siz: byte(15),
				Wal: common.HexToAddress("0x9632185d3851Fd06304C09BA6F1c1308189BE12b"),
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
