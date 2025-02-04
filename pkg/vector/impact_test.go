package vector

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Impact_True(t *testing.T) {
	testCases := []struct {
		oxy object.Object
		osz byte
		txy object.Object
		tsz byte
	}{
		// Case 000
		//
		//      o |
		//     ---+---
		//        | e
		//
		{
			oxy: object.Object{X: 477_999, Y: 510_401}, osz: 80,
			txy: object.Object{X: 478_109, Y: 510_326}, tsz: 60,
		},
		// Case 001
		//
		//        | o
		//     ---+---
		//      e |
		//
		{
			oxy: object.Object{X: 478_087, Y: 510_381}, osz: 50,
			txy: object.Object{X: 478_047, Y: 510_324}, tsz: 20,
		},
		// Case 002
		//
		//      e | o
		//     ---+---
		//        |
		//
		{
			oxy: object.Object{X: 478_222, Y: 510_389}, osz: 50,
			txy: object.Object{X: 478_131, Y: 510_389}, tsz: 50,
		},
		// Case 003
		//
		//      o |
		//     ---+---
		//      e |
		//
		{
			oxy: object.Object{X: 478_305, Y: 510_389}, osz: 30,
			txy: object.Object{X: 478_305, Y: 510_334}, tsz: 40,
		},
		// Case 004
		//
		//     oe |
		//     ---+---
		//        |
		//
		{
			oxy: object.Object{X: 478_305, Y: 510_312}, osz: 10,
			txy: object.Object{X: 478_305, Y: 510_312}, tsz: 25,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			imp := Impact(tc.oxy, tc.osz, tc.txy, tc.tsz)

			if imp != true {
				t.Fatalf("expected %#v got %#v", true, imp)
			}
		})
	}
}

func Test_Vector_Impact_False(t *testing.T) {
	testCases := []struct {
		oxy object.Object
		osz byte
		txy object.Object
		tsz byte
	}{
		// Case 000
		//
		//      o |
		//     ---+---
		//        | e
		//
		{
			oxy: object.Object{X: 477_999, Y: 510_401}, osz: 80,
			txy: object.Object{X: 478_109, Y: 510_326}, tsz: 20,
		},
		// Case 001
		//
		//        | o
		//     ---+---
		//      e |
		//
		{
			oxy: object.Object{X: 478_087, Y: 510_381}, osz: 40,
			txy: object.Object{X: 478_047, Y: 510_324}, tsz: 20,
		},
		// Case 002
		//
		//      e | o
		//     ---+---
		//        |
		//
		{
			oxy: object.Object{X: 478_222, Y: 510_389}, osz: 40,
			txy: object.Object{X: 478_131, Y: 510_389}, tsz: 20,
		},
		// Case 003
		//
		//      o |
		//     ---+---
		//      e |
		//
		{
			oxy: object.Object{X: 478_305, Y: 510_389}, osz: 10,
			txy: object.Object{X: 478_305, Y: 510_334}, tsz: 40,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			imp := Impact(tc.oxy, tc.osz, tc.txy, tc.tsz)

			if imp != false {
				t.Fatalf("expected %#v got %#v", false, imp)
			}
		})
	}
}

func Benchmark_Vector_Impact(b *testing.B) {
	testCases := []struct {
		oxy object.Object
		osz byte
		txy object.Object
		tsz byte
	}{
		// Case 000, ~0.30 ns/op
		{
			oxy: object.Object{X: 477_999, Y: 510_401}, osz: 80,
			txy: object.Object{X: 478_109, Y: 510_326}, tsz: 60,
		},
		// Case 001, ~0.30 ns/op
		{
			oxy: object.Object{X: 478_087, Y: 510_381}, osz: 50,
			txy: object.Object{X: 478_047, Y: 510_324}, tsz: 20,
		},
		// Case 002, ~0.30 ns/op
		{
			oxy: object.Object{X: 478_222, Y: 510_389}, osz: 50,
			txy: object.Object{X: 478_131, Y: 510_389}, tsz: 50,
		},
		// Case 003, ~0.30 ns/op
		{
			oxy: object.Object{X: 478_305, Y: 510_389}, osz: 30,
			txy: object.Object{X: 478_305, Y: 510_334}, tsz: 40,
		},
		// Case 004, ~0.30 ns/op
		{
			oxy: object.Object{X: 478_305, Y: 510_312}, osz: 10,
			txy: object.Object{X: 478_305, Y: 510_312}, tsz: 25,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for range b.N {
				Impact(tc.oxy, tc.osz, tc.txy, tc.tsz)
			}
		})
	}
}
