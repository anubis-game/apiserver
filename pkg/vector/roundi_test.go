package vector

import (
	"fmt"
	"testing"
)

func Test_Vector_roundI(t *testing.T) {
	testCases := []struct {
		f float64
		i int
	}{
		// Case 000
		{
			f: 0.4,
			i: 0,
		},
		// Case 001
		{
			f: 0.5,
			i: 1,
		},
		// Case 002
		{
			f: 0.6,
			i: 1,
		},
		// Case 003
		{
			f: -0.4,
			i: 0,
		},
		// Case 004
		{
			f: -0.5,
			i: -1,
		},
		// Case 005
		{
			f: -0.6,
			i: -1,
		},
		// Case 006
		{
			f: 1.0,
			i: 1,
		},
		// Case 007
		{
			f: -1.0,
			i: -1,
		},
		// Case 008
		{
			f: 100.0,
			i: 100,
		},
		// Case 009
		{
			f: -100.0,
			i: -100,
		},
		// Case 010
		{
			f: 123456.49,
			i: 123456,
		},
		// Case 011
		{
			f: 123456.50,
			i: 123457,
		},
		// Case 012
		{
			f: -123456.49,
			i: -123456,
		},
		// Case 013
		{
			f: -123456.50,
			i: -123457,
		},
		// Case 014
		{
			f: 0.0,
			i: 0,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			i := roundI(tc.f)

			if i != tc.i {
				t.Fatalf("expected %#v, got: %#v", tc.i, i)
			}
		})
	}
}
