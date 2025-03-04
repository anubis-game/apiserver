package matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Matrix_Coordinate(t *testing.T) {
	testCases := []struct {
		c Coordinate
		b [CoordinateBytes]byte
	}{
		// Case 000
		{
			c: Coordinate{X: 0, Y: 0},
			b: [CoordinateBytes]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001
		{
			c: Coordinate{X: 5, Y: 0},
			b: [CoordinateBytes]byte{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002
		{
			c: Coordinate{X: 255, Y: 256},
			b: [CoordinateBytes]byte{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003
		{
			c: Coordinate{X: 4_096, Y: 11_623},
			b: [CoordinateBytes]byte{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.c.Byt()
			c := NewCoordinate(b[:])

			if !reflect.DeepEqual(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
			if !reflect.DeepEqual(c, tc.c) {
				t.Fatalf("expected %#v got %#v", tc.c, c)
			}
		})
	}
}

func Benchmark_Matrix_Coordinate_Byt(b *testing.B) {
	testCases := []struct {
		c Coordinate
	}{
		// Case 000, ~19 ns/op
		{
			c: Coordinate{X: 0, Y: 0},
		},
		// Case 001, ~19 ns/op
		{
			c: Coordinate{X: 5, Y: 0},
		},
		// Case 002, ~19 ns/op
		{
			c: Coordinate{X: 255, Y: 256},
		},
		// Case 003, ~19 ns/op
		{
			c: Coordinate{X: 4_096, Y: 11_623},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				tc.c.Byt()
			}
		})
	}
}

func Benchmark_Matrix_Coordinate_New(b *testing.B) {
	testCases := []struct {
		b [CoordinateBytes]byte
	}{
		// Case 000, ~65 ns/op
		{
			b: [CoordinateBytes]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001, ~65 ns/op
		{
			b: [CoordinateBytes]byte{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002, ~65 ns/op
		{
			b: [CoordinateBytes]byte{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003, ~65 ns/op
		{
			b: [CoordinateBytes]byte{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				NewCoordinate(tc.b[:])
			}
		})
	}
}
