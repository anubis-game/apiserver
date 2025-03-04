package vector

import (
	"fmt"
	"testing"
)

// TODO:test fix encoding once we know how to encode a Vector initially
// func Test_Vector_Encode(t *testing.T) {
// 	testCases := []struct {
// 		v *Vector
// 		b []byte
// 	}{
// 		// Case 000
// 		{
// 			v: New(Config{
// 				Hea: matrix.Coordinate{
// 					X: 100,
// 					Y: 100,
// 				},
// 				// Obj: []matrix.Coordinate{
// 				// 	{X: 100, Y: 100}, // 0
// 				// 	{X: 103, Y: 103}, // 1
// 				// 	{X: 106, Y: 106}, // 2
// 				// 	{X: 109, Y: 109}, // 3
// 				// 	{X: 112, Y: 112}, // 4
// 				// },
// 			}),
// 			b: []byte{
// 				// action
// 				byte(schema.Body),
// 				// uid
// 				0xa,
// 				// len
// 				0x5,
// 				// coordinates
// 				0x0, 0x0, 0x1, 0x1, 0x24, 0x24, // 0
// 				0x0, 0x0, 0x1, 0x1, 0x27, 0x27, // 1
// 				0x0, 0x0, 0x1, 0x1, 0x2a, 0x2a, // 2
// 				0x0, 0x0, 0x1, 0x1, 0x2d, 0x2d, // 3
// 				0x0, 0x0, 0x1, 0x1, 0x30, 0x30, // 4
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
// 			// b := tc.v.Encode()
// 			// v := decode(b)

// 			// if v.uid != tc.v.uid {
// 			// 	t.Fatalf("expected %#v got %#v", tc.v.uid, v.uid)
// 			// }
// 			// if !slices.Equal(v.Encode(), tc.v.Encode()) {
// 			// 	t.Fatalf("expected %#v got %#v", tc.v.Encode(), v.Encode())
// 			// }
// 			// if len(b) != 33 {
// 			// 	t.Fatalf("expected %#v got %#v", 33, len(b))
// 			// }
// 			// if !slices.Equal(b, tc.b) {
// 			// 	t.Fatalf("expected %#v got %#v", tc.b, b)
// 			// }
// 		})
// 	}
// }

func Benchmark_Vector_Encode(b *testing.B) {
	testCases := []struct {
		v *Vector
	}{
		// Case 000, ~24.50 ns/op
		{
			v: New(Config{
				// Obj: []matrix.Coordinate{
				// 	{X: 100, Y: 100}, // 0
				// 	{X: 103, Y: 103}, // 1
				// 	{X: 106, Y: 106}, // 2
				// 	{X: 109, Y: 109}, // 3
				// 	{X: 112, Y: 112}, // 4
				// },
			}),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				tc.v.Encode()
			}
		})
	}
}
