package vector

import (
	"fmt"
	"testing"
)

func Test_Vector_Window_Has_True(t *testing.T) {
	testCases := []struct {
		obj Object
	}{
		// Case 000
		//
		//                  200
		//         +----------+ 200
		//         | x        |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{100, 200},
		},
		// Case 001
		//
		//                  200
		//         +----------+ 200
		//         |    x     |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{150, 200},
		},
		// Case 002
		//
		//                  200
		//         +----------+ 200
		//         |        x |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{200, 200},
		},
		// Case 003
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |        x |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{200, 150},
		},
		// Case 004
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |        x |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{200, 100},
		},
		// Case 005
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |    x     |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{150, 100},
		},
		// Case 006
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         | x        |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{100, 100},
		},
		// Case 007
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         | x        |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{100, 150},
		},
		// Case 008
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |    x     |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: Object{137, 143},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := Window{
				OBL: Object{100, 100},
				OTR: Object{200, 200},
			}

			has := win.Has(tc.obj)

			if has != true {
				t.Fatalf("expected %#v got %#v", true, has)
			}
		})
	}
}

func Test_Vector_Window_Has_False(t *testing.T) {
	testCases := []struct {
		obj []Object
	}{
		// Case 000
		//
		//       x          200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []Object{
				{86, 206},
				{99, 201},
				{99, 200},
				{100, 201},
			},
		},
		// Case 001
		//
		//              x   200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []Object{
				{150, 201},
				{149, 201},
				{151, 236},
			},
		},
		// Case 002
		//
		//                  200 x
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []Object{
				{208, 299},
				{201, 201},
				{201, 200},
				{200, 201},
			},
		},
		// Case 003
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          | x
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []Object{
				{201, 151},
				{201, 150},
				{264, 149},
			},
		},
		// Case 004
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100          x
		//
		{
			obj: []Object{
				{256, 87},
				{201, 99},
				{200, 99},
				{201, 100},
			},
		},
		// Case 005
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//         100  x
		//
		{
			obj: []Object{
				{149, 99},
				{150, 99},
				{151, 85},
			},
		},
		// Case 006
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//         |          |
		//         |          |
		//     100 +----------+
		//       x 100
		//
		{
			obj: []Object{
				{96, 92},
				{99, 99},
				{99, 100},
				{100, 99},
			},
		},
		// Case 007
		//
		//                  200
		//         +----------+ 200
		//         |          |
		//       x |          |
		//         |          |
		//     100 +----------+
		//         100
		//
		{
			obj: []Object{
				{99, 149},
				{99, 150},
				{92, 151},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			win := Window{
				OBL: Object{100, 100},
				OTR: Object{200, 200},
			}

			for _, x := range tc.obj {
				has := win.Has(x)

				if has != false {
					t.Fatalf("expected %#v got %#v", false, has)
				}
			}
		})
	}
}

func Benchmark_Vector_Window_Has(b *testing.B) {
	testCases := []struct {
		obj Object
	}{
		// Case 000, 0.30 ns/op
		{
			obj: Object{100, 200},
		},
		// Case 001, 0.30 ns/op
		{
			obj: Object{150, 200},
		},
		// Case 002, 0.30 ns/op
		{
			obj: Object{200, 200},
		},
		// Case 003, 0.30 ns/op
		{
			obj: Object{200, 150},
		},
		// Case 004, 0.30 ns/op
		{
			obj: Object{200, 100},
		},
		// Case 005, 0.30 ns/op
		{
			obj: Object{150, 100},
		},
		// Case 006, 0.30 ns/op
		{
			obj: Object{100, 100},
		},
		// Case 007, 0.30 ns/op
		{
			obj: Object{100, 150},
		},
		// Case 008, 0.30 ns/op
		{
			obj: Object{137, 143},
		},
		// Case 009, 0.30 ns/op
		{
			obj: Object{86, 206},
		},
		// Case 010, 0.30 ns/op
		{
			obj: Object{149, 201},
		},
		// Case 011, 0.30 ns/op
		{
			obj: Object{201, 200},
		},
		// Case 012, 0.30 ns/op
		{
			obj: Object{264, 149},
		},
		// Case 013, 0.30 ns/op
		{
			obj: Object{201, 99},
		},
		// Case 014, 0.30 ns/op
		{
			obj: Object{149, 99},
		},
		// Case 015, 0.30 ns/op
		{
			obj: Object{100, 99},
		},
		// Case 016, 0.30 ns/op
		{
			obj: Object{99, 150},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			win := Window{
				OBL: Object{100, 100},
				OTR: Object{200, 200},
			}

			b.ResetTimer()
			for range b.N {
				win.Has(tc.obj)
			}
		})
	}
}
