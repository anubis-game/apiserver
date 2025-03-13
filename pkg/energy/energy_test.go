package energy

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func Test_Energy_New(t *testing.T) {
	testCases := []struct {
		c Config
		b []byte
	}{
		// Case 000
		{
			c: Config{
				Crd: matrix.Coordinate{X: 12_547, Y: 512},
				Siz: 0x7c,
				Typ: 0x3,
			},
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
			b := New(tc.c)
			c := decode(b)

			if !reflect.DeepEqual(c, tc.c) {
				t.Fatalf("expected %#v got %#v", tc.c, c)
			}
			if !reflect.DeepEqual(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
		})
	}
}

// ~11 ns/op, 1 allocs/op
func Benchmark_Energy_New(b *testing.B) {
	for b.Loop() {
		New(Config{
			Crd: matrix.Coordinate{X: 12_547, Y: 512},
			Siz: 0x7c,
			Typ: 0x3,
		})
	}
}
