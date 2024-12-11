package decode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Decode_Decode(t *testing.T) {
	testCases := []struct {
		b []byte
		d [][]byte
	}{
		// Case 000
		{
			b: []byte("1,2,3,4"),
			d: [][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
			},
		},
		// Case 001
		{
			b: []byte("1234,abcd,5678,2347652171,jsdgwfxwqjdwb, 23ftgdbhed282  ,10"),
			d: [][]byte{
				[]byte("1234"),
				[]byte("abcd"),
				[]byte("5678"),
				[]byte("2347652171"),
				[]byte("jsdgwfxwqjdwb"),
				[]byte(" 23ftgdbhed282  "),
				[]byte("10"),
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			d := Decode(tc.b)
			if !reflect.DeepEqual(d, tc.d) {
				t.Fatal("expected", tc.d, "got", d)
			}
		})
	}
}

func Benchmark_Decode(b *testing.B) {
	testCases := []struct {
		b []byte
	}{
		// Case 000 ~57.5 ns/op
		{
			b: []byte("1,2,3,4"),
		},
		// Case 001 ~100.0 ns/op
		{
			b: []byte("1234,abcd,5678,2347652171,jsdgwfxwqjdwb, 23ftgdbhed282  ,10"),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Decode(tc.b)
			}
		})
	}
}
