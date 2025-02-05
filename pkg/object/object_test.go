package object

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Object(t *testing.T) {
	testCases := []struct {
		o Object
		b Bucket
	}{
		// Case 000
		{
			o: Object{X: 0, Y: 0},
			b: Bucket{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001
		{
			o: Object{X: 5, Y: 0},
			b: Bucket{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002
		{
			o: Object{X: 255, Y: 256},
			b: Bucket{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003
		{
			o: Object{X: 4_096, Y: 11_623},
			b: Bucket{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := tc.o.Bucket()
			o := b.Object()

			if !reflect.DeepEqual(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
			if !reflect.DeepEqual(o, tc.o) {
				t.Fatalf("expected %#v got %#v", tc.o, o)
			}
		})
	}
}

var bucSnk Bucket

func Benchmark_Object_Bucket(b *testing.B) {
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
				bucSnk = tc.o.Bucket()
			}
		})
	}
}

var objSnk Object

func Benchmark_Object_Bucket_Object(b *testing.B) {
	testCases := []struct {
		b Bucket
	}{
		// Case 000, ~1.00 ns/op
		{
			b: Bucket{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001, ~1.00 ns/op
		{
			b: Bucket{0x0, 0x0, 0x0, 0x0, 0x5, 0x0},
		},
		// Case 002, ~1.00 ns/op
		{
			b: Bucket{0x0, 0x0, 0x3, 0x4, 0x3f, 0x0},
		},
		// Case 003, ~1.00 ns/op
		{
			b: Bucket{0x1, 0x2, 0x0, 0x35, 0x0, 0x27},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				objSnk = tc.b.Object()
			}
		})
	}
}
