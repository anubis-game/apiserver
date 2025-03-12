package vector

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_Impact_True(t *testing.T) {
	testCases := []struct {
		vhc matrix.Coordinate // vector head coordinate
		vnr byte              // vector node radius
		inc matrix.Coordinate // impact node coordinate
		inr byte              // impact node radius
	}{
		// Case 000
		//
		//      o |
		//     ---+---
		//        | e
		//
		{
			vhc: matrix.Coordinate{X: 477_999, Y: 510_401}, vnr: 80,
			inc: matrix.Coordinate{X: 478_109, Y: 510_326}, inr: 60,
		},
		// Case 001
		//
		//        | o
		//     ---+---
		//      e |
		//
		{
			vhc: matrix.Coordinate{X: 478_087, Y: 510_381}, vnr: 50,
			inc: matrix.Coordinate{X: 478_047, Y: 510_324}, inr: 20,
		},
		// Case 002
		//
		//      e | o
		//     ---+---
		//        |
		//
		{
			vhc: matrix.Coordinate{X: 478_222, Y: 510_389}, vnr: 50,
			inc: matrix.Coordinate{X: 478_131, Y: 510_389}, inr: 50,
		},
		// Case 003
		//
		//      o |
		//     ---+---
		//      e |
		//
		{
			vhc: matrix.Coordinate{X: 478_305, Y: 510_389}, vnr: 30,
			inc: matrix.Coordinate{X: 478_305, Y: 510_334}, inr: 40,
		},
		// Case 004
		//
		//     oe |
		//     ---+---
		//        |
		//
		{
			vhc: matrix.Coordinate{X: 478_305, Y: 510_312}, vnr: 10,
			inc: matrix.Coordinate{X: 478_305, Y: 510_312}, inr: 25,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = &Vector{
					crx: Charax{Rad: tc.vnr},
					hea: &Linker{crd: tc.vhc},
				}
			}

			imp := vec.Impact(tc.inc, tc.inr)

			if imp != true {
				t.Fatalf("expected %#v got %#v", true, imp)
			}
		})
	}
}

func Test_Vector_Impact_False(t *testing.T) {
	testCases := []struct {
		vhc matrix.Coordinate
		vnr byte
		inc matrix.Coordinate
		inr byte
	}{
		// Case 000
		//
		//      o |
		//     ---+---
		//        | e
		//
		{
			vhc: matrix.Coordinate{X: 477_999, Y: 510_401}, vnr: 80,
			inc: matrix.Coordinate{X: 478_109, Y: 510_326}, inr: 20,
		},
		// Case 001
		//
		//        | o
		//     ---+---
		//      e |
		//
		{
			vhc: matrix.Coordinate{X: 478_087, Y: 510_381}, vnr: 40,
			inc: matrix.Coordinate{X: 478_047, Y: 510_324}, inr: 20,
		},
		// Case 002
		//
		//      e | o
		//     ---+---
		//        |
		//
		{
			vhc: matrix.Coordinate{X: 478_222, Y: 510_389}, vnr: 40,
			inc: matrix.Coordinate{X: 478_131, Y: 510_389}, inr: 20,
		},
		// Case 003
		//
		//      o |
		//     ---+---
		//      e |
		//
		{
			vhc: matrix.Coordinate{X: 478_305, Y: 510_389}, vnr: 10,
			inc: matrix.Coordinate{X: 478_305, Y: 510_334}, inr: 40,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = &Vector{
					crx: Charax{Rad: tc.vnr},
					hea: &Linker{crd: tc.vhc},
				}
			}

			imp := vec.Impact(tc.inc, tc.inr)

			if imp != false {
				t.Fatalf("expected %#v got %#v", false, imp)
			}
		})
	}
}

func Benchmark_Vector_Impact(b *testing.B) {
	testCases := []struct {
		vhc matrix.Coordinate
		vnr byte
		inc matrix.Coordinate
		inr byte
	}{
		// Case 000, ~2 ns/op
		{
			vhc: matrix.Coordinate{X: 477_999, Y: 510_401}, vnr: 80,
			inc: matrix.Coordinate{X: 478_109, Y: 510_326}, inr: 60,
		},
		// Case 001, ~2 ns/op
		{
			vhc: matrix.Coordinate{X: 478_087, Y: 510_381}, vnr: 50,
			inc: matrix.Coordinate{X: 478_047, Y: 510_324}, inr: 20,
		},
		// Case 002, ~2 ns/op
		{
			vhc: matrix.Coordinate{X: 478_222, Y: 510_389}, vnr: 50,
			inc: matrix.Coordinate{X: 478_131, Y: 510_389}, inr: 50,
		},
		// Case 003, ~2 ns/op
		{
			vhc: matrix.Coordinate{X: 478_305, Y: 510_389}, vnr: 30,
			inc: matrix.Coordinate{X: 478_305, Y: 510_334}, inr: 40,
		},
		// Case 004, ~2 ns/op
		{
			vhc: matrix.Coordinate{X: 478_305, Y: 510_312}, vnr: 10,
			inc: matrix.Coordinate{X: 478_305, Y: 510_312}, inr: 25,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			var vec *Vector
			{
				vec = &Vector{
					crx: Charax{Rad: tc.vnr},
					hea: &Linker{crd: tc.vhc},
				}
			}

			for b.Loop() {
				vec.Impact(tc.inc, tc.inr)
			}
		})
	}
}
