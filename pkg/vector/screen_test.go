package vector

import (
	"slices"
	"sort"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

// TODO:test move screen top, bottom, left

func Test_Vector_Screen_move_right(t *testing.T) {
	var vec *Vector
	{
		vec = New(Config{
			Hea: matrix.Coordinate{
				X: 300,
				Y: 400,
			},
			Mot: Motion{
				Qdr: 0x1,
				Agl: 0x0,
			},
		})
	}

	for range 4 {
		vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
	}

	var act []matrix.Partition
	{
		act = slices.Clone(vec.scr) // don't change the screen partition order
	}

	var exp []matrix.Partition
	{
		exp = []matrix.Partition{
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
		sort.Sort(matrix.Partitions(exp))
		sort.Sort(matrix.Partitions(act))
	}

	if len(act) != 25 {
		t.Fatalf("expected %#v got %#v", 25, len(act))
	}
	if !slices.Equal(act, exp) {
		t.Fatalf("expected %#v got %#v", exp, act)
	}

	// After 18 more update cycles by rotation, we explore a new set of partition
	// coordinates to the east.

	for range 18 {
		vec.Update(0, 0x2, 0x0, Nrm)
	}

	{
		act = slices.Clone(vec.scr) // don't change the screen partition order
	}

	{
		exp = []matrix.Partition{
			{X: 128, Y: 640}, {X: 256, Y: 640} /*    */, {X: 384, Y: 640} /*    */, {X: 512, Y: 640}, {X: 640, Y: 640},
			/*                                                                                                       */
			{X: 128, Y: 512}, {X: 256, Y: 512} /*    */, {X: 384, Y: 512} /*    */, {X: 512, Y: 512}, {X: 640, Y: 512},
			/*                                                 HEAD                                                  */
			{X: 128, Y: 384}, {X: 256, Y: 384} /*HEAD*/, {X: 384, Y: 384} /*HEAD*/, {X: 512, Y: 384}, {X: 640, Y: 384},
			/*                                                 HEAD                                                  */
			{X: 128, Y: 256}, {X: 256, Y: 256} /*    */, {X: 384, Y: 256} /*    */, {X: 512, Y: 256}, {X: 640, Y: 256},
			/*                                                                                                       */
			{X: 128, Y: 128}, {X: 256, Y: 128} /*    */, {X: 384, Y: 128} /*    */, {X: 512, Y: 128}, {X: 640, Y: 128},
		}
	}

	{
		sort.Sort(matrix.Partitions(exp))
		sort.Sort(matrix.Partitions(act))
	}

	if len(act) != 25 {
		t.Fatalf("expected %#v got %#v", 25, len(act))
	}
	if !slices.Equal(act, exp) {
		t.Fatalf("expected %#v got %#v", exp, act)
	}
}
