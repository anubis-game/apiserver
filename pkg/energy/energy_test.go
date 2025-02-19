package energy

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func Test_Energy_Encode(t *testing.T) {
	testCases := []struct {
		e *Energy
		b []byte
	}{
		// Case 000
		{
			e: New(Config{
				Obj: object.Object{X: 12_547, Y: 512},
				Siz: 0x7c,
				Typ: 0x3,
			}),
			b: []byte{
				// action
				byte(schema.Food),
				// obj
				0x3, 0x0, 0x4, 0x8, 0x3, 0x0,
				// siz
				0x7c,
				// typ
				0x3,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.e.Encode()
			e := decode(b)

			if !reflect.DeepEqual(e, tc.e) {
				t.Fatalf("expected %#v got %#v", tc.e, e)
			}
			if !reflect.DeepEqual(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
		})
	}
}
