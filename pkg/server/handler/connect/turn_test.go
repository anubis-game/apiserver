package connect

import (
	"fmt"
	"testing"
)

func Test_Handler_Connect_turn_quadrant_inside_range(t *testing.T) {
	testCases := []struct {
		b byte
	}{
		// Case 000
		{
			b: 1,
		},
		// Case 001
		{
			b: 2,
		},
		// Case 002
		{
			b: 3,
		},
		// Case 003
		{
			b: 4,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			r := tc.b-1 > 3
			if r {
				t.Fatalf("expected %#v got %#v", false, r)
			}
		})
	}
}

func Test_Handler_Connect_turn_quadrant_outside_range(t *testing.T) {
	testCases := []struct {
		b byte
	}{
		// Case 000
		{
			b: 0,
		},
		// Case 001
		{
			b: 5,
		},
		// Case 002
		{
			b: 6,
		},
		// Case 003
		{
			b: 188,
		},
		// Case 004
		{
			b: 255,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			r := tc.b-1 > 3
			if !r {
				t.Fatalf("expected %#v got %#v", true, r)
			}
		})
	}
}

func Benchmark_Handler_Connect_turn_quadrant_one_condition(b *testing.B) {
	testCases := []struct {
		b byte
	}{
		// Case 000, ~1.00 ns/op
		{
			b: 0,
		},
		// Case 001, ~1.00 ns/op
		{
			b: 1,
		},
		// Case 002, ~1.00 ns/op
		{
			b: 2,
		},
		// Case 003, ~1.00 ns/op
		{
			b: 3,
		},
		// Case 004, ~1.00 ns/op
		{
			b: 4,
		},
		// Case 005, ~1.00 ns/op
		{
			b: 5,
		},
		// Case 006, ~1.00 ns/op
		{
			b: 6,
		},
		// Case 007, ~1.00 ns/op
		{
			b: 188,
		},
		// Case 008, ~1.00 ns/op
		{
			b: 255,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				_ = tc.b-1 > 3
			}
		})
	}
}

func Benchmark_Handler_Connect_turn_quadrant_two_conditions(b *testing.B) {
	testCases := []struct {
		b byte
	}{
		// Case 000, ~1.00 ns/op
		{
			b: 0,
		},
		// Case 001, ~1.00 ns/op
		{
			b: 1,
		},
		// Case 002, ~1.00 ns/op
		{
			b: 2,
		},
		// Case 003, ~1.00 ns/op
		{
			b: 3,
		},
		// Case 004, ~1.00 ns/op
		{
			b: 4,
		},
		// Case 005, ~1.00 ns/op
		{
			b: 5,
		},
		// Case 006, ~1.00 ns/op
		{
			b: 6,
		},
		// Case 007, ~1.00 ns/op
		{
			b: 188,
		},
		// Case 008, ~1.00 ns/op
		{
			b: 255,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				_ = tc.b < 1 || tc.b > 4
			}
		})
	}
}
