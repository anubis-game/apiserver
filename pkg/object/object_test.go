package object

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Object(t *testing.T) {
	testCases := []struct {
		o Object
		b [6]byte
	}{
		// Case 000
		{
			o: Object{X: 0, Y: 0},
			b: [6]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001
		{
			o: Object{X: 5, Y: 0},
			b: [6]byte{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002
		{
			o: Object{X: 255, Y: 256},
			b: [6]byte{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003
		{
			o: Object{X: 4_096, Y: 11_623},
			b: [6]byte{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.o.Byt()
			o := New(b[:])

			if !reflect.DeepEqual(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
			if !reflect.DeepEqual(o, tc.o) {
				t.Fatalf("expected %#v got %#v", tc.o, o)
			}
		})
	}
}

var bucSnk [6]byte

func Benchmark_Object_Byt(b *testing.B) {
	testCases := []struct {
		o Object
	}{
		// Case 000, ~1.60 ns/op
		{
			o: Object{X: 0, Y: 0},
		},
		// Case 001, ~1.60 ns/op
		{
			o: Object{X: 5, Y: 0},
		},
		// Case 002, ~1.60 ns/op
		{
			o: Object{X: 255, Y: 256},
		},
		// Case 003, ~1.60 ns/op
		{
			o: Object{X: 4_096, Y: 11_623},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bucSnk = tc.o.Byt()
			}
		})
	}
}

var objSnk Object

func Benchmark_Object_New(b *testing.B) {
	testCases := []struct {
		b [6]byte
	}{
		// Case 000, ~0.80 ns/op
		{
			b: [6]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001, ~0.80 ns/op
		{
			b: [6]byte{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002, ~0.80 ns/op
		{
			b: [6]byte{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003, ~0.80 ns/op
		{
			b: [6]byte{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				objSnk = New(tc.b[:])
			}
		})
	}
}
