package schema

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	tok = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
)

func Test_Schema_Encode(t *testing.T) {
	testCases := []struct {
		f func(Action, []byte) []byte
	}{
		// Case 000
		{
			f: func(act Action, mes []byte) []byte {
				return bytes.Join(append([][]byte{{byte(act)}}, mes), nil)
			},
		},
		// Case 001
		{
			f: func(act Action, mes []byte) []byte {
				return Encode(act, mes)
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			e := tc.f(Auth, []byte(tok))
			if len(e) != 37 {
				t.Fatal("expected", 37, "got", len(e))
			}
		})
	}
}

func Benchmark_Schema_Encode(b *testing.B) {
	testCases := []struct {
		f func(Action, []byte) []byte
	}{
		// Case 000, ~60.40 ns/op
		{
			f: func(act Action, mes []byte) []byte {
				return bytes.Join(append([][]byte{{byte(act)}}, mes), nil)
			},
		},
		// Case 001, ~26.00 ns/op
		{
			f: func(act Action, mes []byte) []byte {
				return Encode(act, mes)
			},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				tc.f(Auth, []byte(tok))
			}
		})
	}
}
