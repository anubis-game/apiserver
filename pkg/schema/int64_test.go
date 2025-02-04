package schema

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_Schema_BytesToInt64(t *testing.T) {
	testCases := []struct {
		b []byte
		i int64
	}{
		// Case 000
		{
			b: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			i: 0,
		},
		// Case 001
		{
			b: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xd2},
			i: 1234,
		},
		// Case 002
		{
			b: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			i: 9223372036854775807,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			i := BytesToInt64(tc.b)

			if i != tc.i {
				t.Fatalf("expected %#v got %#v", tc.i, i)
			}
		})
	}
}

func Test_Schema_Int64ToBytes(t *testing.T) {
	testCases := []struct {
		i int64
		b []byte
	}{
		// Case 000
		{
			i: 0,
			b: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001
		{
			i: 1234,
			b: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xd2},
		},
		// Case 002
		{
			i: 9223372036854775807,
			b: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			b := Int64ToBytes(tc.i)

			if !slices.Equal(b, tc.b) {
				t.Fatalf("expected %#v got %#v", tc.b, b)
			}
		})
	}
}

func Benchmark_Schema_BytesToInt64(b *testing.B) {
	testCases := []struct {
		b []byte
	}{
		// Case 000 ~0.30 ns/op
		{
			b: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		// Case 001 ~0.30 ns/op
		{
			b: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xd2},
		},
		// Case 002 ~0.30 ns/op
		{
			b: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BytesToInt64(tc.b)
			}
		})
	}
}

func Benchmark_Schema_Int64ToBytes(b *testing.B) {
	testCases := []struct {
		i int64
	}{
		// Case 000 ~0.30 ns/op
		{
			i: 0,
		},
		// Case 001 ~0.30 ns/op
		{
			i: 1234,
		},
		// Case 002 ~0.30 ns/op
		{
			i: 9223372036854775807,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Int64ToBytes(tc.i)
			}
		})
	}
}
