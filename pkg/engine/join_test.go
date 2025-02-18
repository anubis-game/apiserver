package engine

import (
	"fmt"
	"testing"
	"time"
)

func Test_Engine_timCap(t *testing.T) {
	testCases := []struct {
		ply int
		cpu int
		tim time.Duration
	}{
		//
		// 1 CPU
		//

		// Case 000
		{
			ply: 0,
			cpu: 1,
			tim: 24 * time.Millisecond,
		},
		// Case 001
		{
			ply: 1,
			cpu: 1,
			tim: 24 * time.Millisecond,
		},
		// Case 002
		{
			ply: 5,
			cpu: 1,
			tim: 4800 * time.Microsecond,
		},
		// Case 003
		{
			ply: 80,
			cpu: 1,
			tim: 300 * time.Microsecond,
		},
		// Case 004
		{
			ply: 500,
			cpu: 1,
			tim: 48 * time.Microsecond,
		},

		//
		// 2 CPUs
		//

		// Case 005
		{
			ply: 0,
			cpu: 2,
			tim: 12 * time.Millisecond,
		},
		// Case 006
		{
			ply: 1,
			cpu: 2,
			tim: 12 * time.Millisecond,
		},
		// Case 007
		{
			ply: 5,
			cpu: 2,
			tim: 2400 * time.Microsecond,
		},
		// Case 008
		{
			ply: 80,
			cpu: 2,
			tim: 150 * time.Microsecond,
		},
		// Case 009
		{
			ply: 500,
			cpu: 2,
			tim: 24 * time.Microsecond,
		},

		//
		// 4 CPUs
		//

		// Case 010
		{
			ply: 0,
			cpu: 4,
			tim: 6 * time.Millisecond,
		},
		// Case 011
		{
			ply: 1,
			cpu: 4,
			tim: 6 * time.Millisecond,
		},
		// Case 012
		{
			ply: 5,
			cpu: 4,
			tim: 1200 * time.Microsecond,
		},
		// Case 013
		{
			ply: 80,
			cpu: 4,
			tim: 75 * time.Microsecond,
		},
		// Case 014
		{
			ply: 500,
			cpu: 4,
			tim: 12 * time.Microsecond,
		},

		//
		// 8 CPUs
		//

		// Case 015
		{
			ply: 0,
			cpu: 8,
			tim: 3 * time.Millisecond,
		},
		// Case 016
		{
			ply: 1,
			cpu: 8,
			tim: 3 * time.Millisecond,
		},
		// Case 017
		{
			ply: 5,
			cpu: 8,
			tim: 600 * time.Microsecond,
		},
		// Case 018
		{
			ply: 80,
			cpu: 8,
			tim: 37500 * time.Nanosecond,
		},
		// Case 019
		{
			ply: 500,
			cpu: 8,
			tim: 6 * time.Microsecond,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tim := timCap(tc.ply, tc.cpu)

			if tim != tc.tim {
				t.Fatalf("expected %s got %s", tc.tim, tim)
			}
		})
	}
}

func Benchmark_Engine_timCap(b *testing.B) {
	testCases := []struct {
		ply int
		cpu int
	}{
		//
		// 1 CPU
		//

		// Case 000, ~0.40 ns/op
		{
			ply: 0,
			cpu: 1,
		},
		// Case 001, ~0.40 ns/op
		{
			ply: 1,
			cpu: 1,
		},
		// Case 002, ~0.50 ns/op
		{
			ply: 5,
			cpu: 1,
		},
		// Case 003, ~0.50 ns/op
		{
			ply: 80,
			cpu: 1,
		},
		// Case 004, ~0.50 ns/op
		{
			ply: 500,
			cpu: 1,
		},

		//
		// 2 CPUs
		//

		// Case 005, ~0.40 ns/op
		{
			ply: 0,
			cpu: 2,
		},
		// Case 006, ~0.40 ns/op
		{
			ply: 1,
			cpu: 2,
		},
		// Case 007, ~0.60 ns/op
		{
			ply: 5,
			cpu: 2,
		},
		// Case 008, ~0.60 ns/op
		{
			ply: 80,
			cpu: 2,
		},
		// Case 009, ~0.60 ns/op
		{
			ply: 500,
			cpu: 2,
		},

		//
		// 4 CPUs
		//

		// Case 010, ~0.40 ns/op
		{
			ply: 0,
			cpu: 4,
		},
		// Case 011, ~0.40 ns/op
		{
			ply: 1,
			cpu: 4,
		},
		// Case 012, ~0.60 ns/op
		{
			ply: 5,
			cpu: 4,
		},
		// Case 013, ~0.60 ns/op
		{
			ply: 80,
			cpu: 4,
		},
		// Case 014, ~0.60 ns/op
		{
			ply: 500,
			cpu: 4,
		},

		//
		// 8 CPUs
		//

		// Case 015, ~0.40 ns/op
		{
			ply: 0,
			cpu: 8,
		},
		// Case 016, ~0.40 ns/op
		{
			ply: 1,
			cpu: 8,
		},
		// Case 017, ~0.60 ns/op
		{
			ply: 5,
			cpu: 8,
		},
		// Case 018, ~0.60 ns/op
		{
			ply: 80,
			cpu: 8,
		},
		// Case 019, ~0.60 ns/op
		{
			ply: 500,
			cpu: 8,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				timCap(tc.ply, tc.cpu)
			}
		})
	}
}
