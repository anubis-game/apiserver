package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Decode_Int64(t *testing.T) {
	testCases := []struct {
		b []byte
		d int64
	}{
		// Case 000
		{
			b: []byte("1"),
			d: 1,
		},
		// Case 001
		{
			b: []byte("1234"),
			d: 1234,
		},
		// Case 002
		{
			b: []byte("98732347652171"),
			d: 98732347652171,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			d, err := Int64(tc.b)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(d, tc.d) {
				t.Fatal("expected", tc.d, "got", d)
			}
		})
	}
}

func Benchmark_Decode_Int64(b *testing.B) {
	testCases := []struct {
		b []byte
	}{
		// Case 000 ~57 ns/op
		{
			b: []byte("1"),
		},
		// Case 001 ~100 ns/op
		{
			b: []byte("1234"),
		},
		// Case 002 ~100 ns/op
		{
			b: []byte("98732347652171"),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Int64(tc.b)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
