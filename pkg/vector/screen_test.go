package vector

import (
	"slices"
	"sort"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Screen(t *testing.T) {
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
			Mot: Motion{
				Qdr: 0x1,
				Agl: 0x80,
				Vlc: Rcn,
			},
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

	var exp []object.Object
	{
		exp = []object.Object{
			{X: 000, Y: 640}, {X: 128, Y: 640} /*    */, {X: 256, Y: 640} /*    */, {X: 384, Y: 640}, {X: 512, Y: 640},
			/*                                                                                                       */
			{X: 000, Y: 512}, {X: 128, Y: 512} /*    */, {X: 256, Y: 512} /*    */, {X: 384, Y: 512}, {X: 512, Y: 512},
			/*                                                 HEAD                                                  */
			{X: 000, Y: 384}, {X: 128, Y: 384} /*HEAD*/, {X: 256, Y: 384} /*HEAD*/, {X: 384, Y: 384}, {X: 512, Y: 384},
			/*                                                 HEAD                                                  */
			{X: 000, Y: 256}, {X: 128, Y: 256} /*    */, {X: 256, Y: 256} /*    */, {X: 384, Y: 256}, {X: 512, Y: 256},
			/*                                                                                                       */
			{X: 000, Y: 128}, {X: 128, Y: 128} /*    */, {X: 256, Y: 128} /*    */, {X: 384, Y: 128}, {X: 512, Y: 128},
		}
	}

	{
		sort.Sort(object.Sorter(exp))
		sort.Sort(object.Sorter(vec.scr.Prt))
	}

	if len(vec.scr.Prt) != 25 {
		t.Fatalf("expected %#v got %#v", 25, len(vec.scr.Prt))
	}
	if !slices.Equal(exp, vec.scr.Prt) {
		t.Fatalf("expected %#v got %#v", exp, vec.scr.Prt)
	}

	// After 6 iterations of adjusting our view by rotation, we explore a new set
	// of boundary partitions to the east.

	for range 6 {
		vec.Adjust(0, vec.mot.Get())
	}

	{
		exp = []object.Object{
			{X: 640, Y: 640},
			{X: 640, Y: 512},
			{X: 640, Y: 384},
			{X: 640, Y: 256},
			{X: 640, Y: 128},
		}
	}

	{
		sort.Sort(object.Sorter(exp))
		sort.Sort(object.Sorter(vec.scr.Prt))
	}

	if len(vec.scr.Prt) != 5 {
		t.Fatalf("expected %#v got %#v", 5, len(vec.scr.Prt))
	}
	if !slices.Equal(exp, vec.scr.Prt) {
		t.Fatalf("expected %#v got %#v", exp, vec.scr.Prt)
	}

	// After 2 more iterations of adjusting our view by rotation, we explore a new
	// set of boundary partitions to the north.

	for range 2 {
		vec.Adjust(0, vec.mot.Get())
	}

	{
		exp = []object.Object{
			{X: 128, Y: 768}, {X: 256, Y: 768}, {X: 384, Y: 768}, {X: 512, Y: 768}, {X: 640, Y: 768},
		}
	}

	{
		sort.Sort(object.Sorter(exp))
		sort.Sort(object.Sorter(vec.scr.Prt))
	}

	if len(vec.scr.Prt) != 5 {
		t.Fatalf("expected %#v got %#v", 5, len(vec.scr.Prt))
	}
	if !slices.Equal(exp, vec.scr.Prt) {
		t.Fatalf("expected %#v got %#v", exp, vec.scr.Prt)
	}
}
