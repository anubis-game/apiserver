package player

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/vector"
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
				Uid: uuid.MustParse("f47ac10b-58cc-4372-a567-0e02b2c3d479"),
				Vec: vector.New(vector.Config{
					Mot: vector.Motion{
						Qdr: 0x1,
						Agl: 0x7f,
						Vlc: 0x1,
					},
					Obj: []object.Object{
						{X: 100, Y: 100},
						{X: 103, Y: 103},
						{X: 106, Y: 106},
						{X: 109, Y: 109},
						{X: 112, Y: 112},
					},
				}),
			},
			b: []byte{
				// uid
				0xf4, 0x7a, 0xc1, 0xb, 0x58, 0xcc, 0x43, 0x72, 0xa5, 0x67, 0xe, 0x2, 0xb2, 0xc3, 0xd4, 0x79,
				// crx
				0xa, 0x32, 0x0,
				// mot
				0x1, 0x7f, 0x1,
				// vec
				0x0, 0x0, 0x1, 0x1, 0x24, 0x24, 0x0, 0x0, 0x1, 0x1, 0x27, 0x27, 0x0, 0x0, 0x1, 0x1, 0x2a, 0x2a, 0x0, 0x0, 0x1, 0x1, 0x2d, 0x2d, 0x0, 0x0, 0x1, 0x1, 0x30, 0x30,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.p.Bytes()
			p := FromBytes(b)

			if !reflect.DeepEqual(p.Uid, tc.p.Uid) {
				t.Fatalf("expected %#v got %#v", tc.p.Uid, p.Uid)
			}
			if !reflect.DeepEqual(p.Vec.Bytes(), tc.p.Vec.Bytes()) {
				t.Fatalf("expected %#v got %#v", tc.p.Vec.Bytes(), p.Vec.Bytes())
			}
			if !reflect.DeepEqual(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
		})
	}
}
