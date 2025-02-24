package vector

import (
	"fmt"
	"slices"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

// https://www.desmos.com/calculator/cjcyg1ohwz
var vectSmtObj = [][]object.Object{
	// New()
	{
		{X: 100, Y: 100}, // T
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
		{X: 100, Y: 300},
		{X: 100, Y: 350},
		{X: 100, Y: 400},
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
		{X: 300, Y: 400}, // H
	},
	// Case 000
	{
		{X: 100, Y: 100}, // T
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
		{X: 100, Y: 300},
		{X: 100, Y: 350},
		{X: 110, Y: 390}, // x+10 y-10
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
		{X: 300, Y: 400}, // H
	},
	// Case 001
	{
		{X: 100, Y: 100}, // T
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
		{X: 100, Y: 300},
		{X: 102, Y: 348}, // x+2 y-2
		{X: 116, Y: 384}, // x+6 y-6
		{X: 152, Y: 398}, // x+2 y-2
		{X: 200, Y: 400},
		{X: 250, Y: 400},
		{X: 300, Y: 400}, // H
	},
	// Case 002
	{
		{X: 100, Y: 100},
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
		{X: 100, Y: 300},
		{X: 104, Y: 346}, // x+2 y-2
		{X: 120, Y: 380}, // x+4 y-4
		{X: 154, Y: 396}, // x+2 y-2
		{X: 200, Y: 400},
		{X: 250, Y: 400},
		{X: 300, Y: 400},
	},
	// Case 003
	{
		{X: 100, Y: 100},
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
		{X: 101, Y: 299},
		{X: 106, Y: 344},
		{X: 124, Y: 376},
		{X: 156, Y: 394},
		{X: 201, Y: 399},
		{X: 250, Y: 400},
		{X: 300, Y: 400},
	},
	// Case 004
	{
		{X: 100, Y: 100},
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
		{X: 102, Y: 298},
		{X: 109, Y: 341},
		{X: 127, Y: 373},
		{X: 159, Y: 391},
		{X: 202, Y: 398},
		{X: 250, Y: 400},
		{X: 300, Y: 400},
	},
	// Case 005, after 100 Smooth() calls
	{
		{X: 100, Y: 100},
		{X: 111, Y: 139},
		{X: 124, Y: 176},
		{X: 139, Y: 211},
		{X: 156, Y: 244},
		{X: 175, Y: 275},
		{X: 196, Y: 304},
		{X: 219, Y: 331},
		{X: 244, Y: 356},
		{X: 271, Y: 379},
		{X: 300, Y: 400},
	},
}

func Test_Vector_Smooth(t *testing.T) {
	//
	//     +---------HR
	//     │
	//     │
	//     |
	//     T
	//
	var vec *Vector
	{
		vec = New(Config{
			Obj: vectSmtObj[0],
		})
	}

	vec.Ranger(func(obj object.Object) {
		fmt.Printf("(%d,%d),", obj.X, obj.Y)
	})
	fmt.Printf("0\n")

	for i, x := range vectSmtObj[1:6] {
		{
			vec.Smooth()
		}

		var act []object.Object
		vec.Ranger(func(obj object.Object) {
			act = append(act, obj)
			fmt.Printf("(%d,%d),", obj.X, obj.Y)
		})
		fmt.Printf("%#v\n", i+1)

		if len(act) != 11 {
			t.Fatalf("expected %#v got %#v", 11, len(act))
		}
		if !slices.Equal(x, act) {
			t.Fatalf("(%03d) expected %#v got %#v", i, x, act)
		}
	}

	for range 95 {
		vec.Smooth()
	}

	var act []object.Object
	vec.Ranger(func(obj object.Object) {
		act = append(act, obj)
		fmt.Printf("(%d,%d),", obj.X, obj.Y)
	})
	fmt.Printf("6\n")

	if len(act) != 11 {
		t.Fatalf("expected %#v got %#v", 11, len(act))
	}
	if !slices.Equal(vectSmtObj[6], act) {
		t.Fatalf("(%03d) expected %#v got %#v", 5, vectSmtObj[6], act)
	}
}

func Test_Vector_smooth(t *testing.T) {
	testCases := []struct {
		lef object.Object
		mid object.Object
		rig object.Object
		smx int
		smy int
	}{
		// Case 000
		{
			lef: object.Object{X: 00, Y: 00},
			mid: object.Object{X: 05, Y: 10}, // y-4
			rig: object.Object{X: 10, Y: 00},
			smx: 5,
			smy: 6,
		},
		// Case 001
		{
			lef: object.Object{X: 00, Y: 20},
			mid: object.Object{X: 05, Y: 10}, // y+4
			rig: object.Object{X: 10, Y: 20},
			smx: 5,
			smy: 14,
		},
		// Case 002
		{
			lef: object.Object{X: 00, Y: 00},
			mid: object.Object{X: 05, Y: 00},
			rig: object.Object{X: 10, Y: 00},
			smx: 5,
			smy: 0,
		},
		// Case 003
		{
			lef: object.Object{X: 04, Y: 12},
			mid: object.Object{X: 05, Y: 12}, // x+1
			rig: object.Object{X: 11, Y: 12},
			smx: 6,
			smy: 12,
		},
		// Case 004
		{
			lef: object.Object{X: 04, Y: 02},
			mid: object.Object{X: 05, Y: 12}, // x+1 y-4
			rig: object.Object{X: 11, Y: 02},
			smx: 6,
			smy: 8,
		},
		// Case 005, https://www.desmos.com/calculator/ebtgon5qgm
		{
			lef: object.Object{X: 5, Y: 63},
			mid: object.Object{X: 30, Y: 96}, // x-2 y-8
			rig: object.Object{X: 45, Y: 87},
			smx: 28,
			smy: 88,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			smx, smy := smooth(tc.lef, tc.mid, tc.rig)

			if smx != tc.smx {
				t.Fatalf("expected %#v got %#v", tc.smx, smx)
			}
			if smy != tc.smy {
				t.Fatalf("expected %#v got %#v", tc.smy, smy)
			}
		})
	}
}

// ~15.00 ns/op
func Benchmark_Vector_Smooth(b *testing.B) {
	var vec *Vector
	{
		vec = New(Config{
			Obj: []object.Object{
				{X: 100, Y: 100}, // T
				{X: 100, Y: 150},
				{X: 100, Y: 200},
				{X: 100, Y: 250},
				{X: 100, Y: 300},
				{X: 100, Y: 350},
				{X: 100, Y: 400},
				{X: 150, Y: 400},
				{X: 200, Y: 400},
				{X: 250, Y: 400},
				{X: 300, Y: 400}, // H
			},
		})
	}

	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		for b.Loop() {
			vec.Smooth()
		}
	})
}

func Benchmark_Vector_smooth(b *testing.B) {
	testCases := []struct {
		lef object.Object
		mid object.Object
		rig object.Object
		smx int
		smy int
	}{
		// Case 000, 1.90 ns/op
		{
			lef: object.Object{X: 00, Y: 00},
			mid: object.Object{X: 05, Y: 10},
			rig: object.Object{X: 10, Y: 00},
		},
		// Case 001, 1.90 ns/op
		{
			lef: object.Object{X: 00, Y: 20},
			mid: object.Object{X: 05, Y: 10},
			rig: object.Object{X: 10, Y: 20},
		},
		// Case 002, 1.90 ns/op
		{
			lef: object.Object{X: 00, Y: 00},
			mid: object.Object{X: 05, Y: 00},
			rig: object.Object{X: 10, Y: 00},
		},
		// Case 003, 1.90 ns/op
		{
			lef: object.Object{X: 04, Y: 12},
			mid: object.Object{X: 05, Y: 12},
			rig: object.Object{X: 11, Y: 12},
		},
		// Case 004, 1.90 ns/op
		{
			lef: object.Object{X: 04, Y: 02},
			mid: object.Object{X: 05, Y: 12},
			rig: object.Object{X: 11, Y: 02},
		},
		// Case 005, 1.90 ns/op
		{
			lef: object.Object{X: 5, Y: 63},
			mid: object.Object{X: 30, Y: 96},
			rig: object.Object{X: 45, Y: 87},
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				smooth(tc.lef, tc.mid, tc.rig)
			}
		})
	}
}
