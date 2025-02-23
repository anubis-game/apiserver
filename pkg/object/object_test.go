package object

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Object(t *testing.T) {
	testCases := []struct {
		o Object
		b [Len]byte
	}{
		// Case 000
		{
			o: Object{X: 0, Y: 0},
			b: [Len]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001
		{
			o: Object{X: 5, Y: 0},
			b: [Len]byte{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002
		{
			o: Object{X: 255, Y: 256},
			b: [Len]byte{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003
		{
			o: Object{X: 4_096, Y: 11_623},
			b: [Len]byte{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
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

func Benchmark_Object_Byt(b *testing.B) {
	testCases := []struct {
		o Object
	}{
		// Case 000, ~2.00 ns/op
		{
			o: Object{X: 0, Y: 0},
		},
		// Case 001, ~2.00 ns/op
		{
			o: Object{X: 5, Y: 0},
		},
		// Case 002, ~2.00 ns/op
		{
			o: Object{X: 255, Y: 256},
		},
		// Case 003, ~2.00 ns/op
		{
			o: Object{X: 4_096, Y: 11_623},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				tc.o.Byt()
			}
		})
	}
}

func Benchmark_Object_New(b *testing.B) {
	testCases := []struct {
		b [Len]byte
	}{
		// Case 000, ~2.00 ns/op
		{
			b: [Len]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001, ~2.00 ns/op
		{
			b: [Len]byte{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002, ~2.00 ns/op
		{
			b: [Len]byte{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003, ~2.00 ns/op
		{
			b: [Len]byte{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				New(tc.b[:])
			}
		})
	}
}
