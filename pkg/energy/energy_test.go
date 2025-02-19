package energy

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Energy_Encode(t *testing.T) {
	testCases := []struct {
		e *Energy
		b []byte
	}{
		// Case 000
		{
			e: &Energy{
				Obj: object.Object{X: 12_547, Y: 512},
				Siz: 0x7c,
				Typ: 0x3,
			},
			b: []byte{0x3, 0x0, 0x4, 0x8, 0x3, 0x0, 0x7c, 0x3},
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

func Benchmark_Energy_Encode(b *testing.B) {
	testCases := []struct {
		e *Energy
	}{
		// Case 000, ~1.50 ns/op
		{
			e: &Energy{
				Obj: object.Object{X: 12_547, Y: 512},
				Siz: 0x7c,
				Typ: 0x3,
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tc.e.Encode()
			}
		})
	}
}
