package player

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func Test_Player_EncodeVector(t *testing.T) {
	testCases := []struct {
		p *Player
		b []byte
	}{
		// Case 000
		{
			p: New(Config{
				Cli: client.New(client.Config{}),
				Uid: [2]byte{0x0, 0x5},
				Vec: vector.New(vector.Config{
					Mot: vector.Motion{
						Qdr: 0x1,
						Agl: 0x80,
						Vlc: vector.Nrm,
					},
					Obj: []object.Object{
						{X: 100, Y: 100}, // 0
						{X: 103, Y: 103}, // 1
						{X: 106, Y: 106}, // 2
						{X: 109, Y: 109}, // 3
						{X: 112, Y: 112}, // 4
					},
				}),
			}),
			b: []byte{
				// uid
				0x0, 0x5,
				// crx
				0xa, 0x32, 0x0,
				// mot
				0x1, 0x80, 0x1,
				// vec
				0x0, 0x0, 0x1, 0x1, 0x24, 0x24, // 0
				0x0, 0x0, 0x1, 0x1, 0x27, 0x27, // 1
				0x0, 0x0, 0x1, 0x1, 0x2a, 0x2a, // 2
				0x0, 0x0, 0x1, 0x1, 0x2d, 0x2d, // 3
				0x0, 0x0, 0x1, 0x1, 0x30, 0x30, // 4
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.p.Encode()
			p := Decode(b)

			if !reflect.DeepEqual(p.Uid, tc.p.Uid) {
				t.Fatalf("expected %#v got %#v", tc.p.Uid, p.Uid)
			}
			if !reflect.DeepEqual(p.Vec.Encode(), tc.p.Vec.Encode()) {
				t.Fatalf("expected %#v got %#v", tc.p.Vec.Encode(), p.Vec.Encode())
			}
			if !reflect.DeepEqual(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
		})
	}
}
