package vector

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

// TODO:test move screen top, bottom, left

func Test_Vector_Bounds_move_right(t *testing.T) {
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

	var atp int
	var arg int
	var abt int
	var alf int
	{
		atp, arg, abt, alf = vec.Bounds()
	}

	//
	//     {X: 000, Y: 640}, {X: 128, Y: 640} /*    */, {X: 256, Y: 640} /*    */, {X: 384, Y: 640}, {X: 512, Y: 640},
	//     /*                                                                                                       */
	//     {X: 000, Y: 512}, {X: 128, Y: 512} /*    */, {X: 256, Y: 512} /*    */, {X: 384, Y: 512}, {X: 512, Y: 512},
	//     /*                                                 HEAD                                                  */
	//     {X: 000, Y: 384}, {X: 128, Y: 384} /*HEAD*/, {X: 256, Y: 384} /*HEAD*/, {X: 384, Y: 384}, {X: 512, Y: 384},
	//     /*                                                 HEAD                                                  */
	//     {X: 000, Y: 256}, {X: 128, Y: 256} /*    */, {X: 256, Y: 256} /*    */, {X: 384, Y: 256}, {X: 512, Y: 256},
	//     /*                                                                                                       */
	//     {X: 000, Y: 128}, {X: 128, Y: 128} /*    */, {X: 256, Y: 128} /*    */, {X: 384, Y: 128}, {X: 512, Y: 128},
	//
	var etp int
	var erg int
	var ebt int
	var elf int
	{
		etp = 640
		erg = 512
		ebt = 128
		elf = 0
	}

	if atp != etp {
		t.Fatalf("expected %#v got %#v", etp, atp)
	}
	if arg != erg {
		t.Fatalf("expected %#v got %#v", erg, arg)
	}
	if abt != ebt {
		t.Fatalf("expected %#v got %#v", ebt, abt)
	}
	if alf != elf {
		t.Fatalf("expected %#v got %#v", elf, alf)
	}

	// After 18 more update cycles by rotation, we explore a new set of partition
	// coordinates to the east.

	for range 18 {
		vec.Update(0, 0x2, 0x0, Nrm)
	}

	{
		atp, arg, abt, alf = vec.Bounds()
	}

	//
	//     {X: 128, Y: 640}, {X: 256, Y: 640} /*    */, {X: 384, Y: 640} /*    */, {X: 512, Y: 640}, {X: 640, Y: 640},
	//     /*                                                                                                       */
	//     {X: 128, Y: 512}, {X: 256, Y: 512} /*    */, {X: 384, Y: 512} /*    */, {X: 512, Y: 512}, {X: 640, Y: 512},
	//     /*                                                 HEAD                                                  */
	//     {X: 128, Y: 384}, {X: 256, Y: 384} /*HEAD*/, {X: 384, Y: 384} /*HEAD*/, {X: 512, Y: 384}, {X: 640, Y: 384},
	//     /*                                                 HEAD                                                  */
	//     {X: 128, Y: 256}, {X: 256, Y: 256} /*    */, {X: 384, Y: 256} /*    */, {X: 512, Y: 256}, {X: 640, Y: 256},
	//     /*                                                                                                       */
	//     {X: 128, Y: 128}, {X: 256, Y: 128} /*    */, {X: 384, Y: 128} /*    */, {X: 512, Y: 128}, {X: 640, Y: 128},
	//
	{
		etp = 640
		erg = 640 // move right
		ebt = 128
		elf = 128 // move right
	}

	if atp != etp {
		t.Fatalf("expected %#v got %#v", etp, atp)
	}
	if arg != erg {
		t.Fatalf("expected %#v got %#v", erg, arg)
	}
	if abt != ebt {
		t.Fatalf("expected %#v got %#v", ebt, abt)
	}
	if alf != elf {
		t.Fatalf("expected %#v got %#v", elf, alf)
	}
}

func Test_Vector_Bounds_fos_2_to_7(t *testing.T) {
	testCases := []struct {
		siz int
		fos int
		top int
		rig int
		bot int
		lef int
	}{
		// Case 000
		{
			siz: 0,
			fos: 2,
			top: 4_224,
			rig: 3_200,
			bot: 3_712,
			lef: 2_688,
		},
		// Case 001
		{
			siz: 250,
			fos: 3,
			top: 4_352,
			rig: 3_328,
			bot: 3_584,
			lef: 2_560,
		},
		// Case 002
		{
			siz: 1_000,
			fos: 4,
			top: 4_480,
			rig: 3_456,
			bot: 3_456,
			lef: 2_432,
		},
		// Case 003
		{
			siz: 5_000,
			fos: 5,
			top: 4_608,
			rig: 3_584,
			bot: 3_328,
			lef: 2_304,
		},
		// Case 004
		{
			siz: 10_000,
			fos: 6,
			top: 4_736,
			rig: 3_712,
			bot: 3_200,
			lef: 2_176,
		},
		// Case 005
		{
			siz: 25_000,
			fos: 7,
			top: 4_864,
			rig: 3_840,
			bot: 3_072,
			lef: 2_048,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = New(Config{
					Hea: matrix.Coordinate{
						X: 3000,
						Y: 4000,
					},
					Mot: Motion{
						Qdr: 0x1,
						Agl: 0x0,
					},
				})
			}

			{
				vec.Update(tc.siz, 0x1, 0x0, Nrm)
			}

			var top int
			var rig int
			var bot int
			var lef int
			{
				top, rig, bot, lef = vec.Bounds()
			}

			if vec.crx.Fos != tc.fos {
				t.Fatalf("expected %#v got %#v", tc.fos, vec.crx.Fos)
			}
			if top != tc.top {
				t.Fatalf("expected %#v got %#v", tc.top, top)
			}
			if rig != tc.rig {
				t.Fatalf("expected %#v got %#v", tc.rig, rig)
			}
			if bot != tc.bot {
				t.Fatalf("expected %#v got %#v", tc.bot, bot)
			}
			if lef != tc.lef {
				t.Fatalf("expected %#v got %#v", tc.lef, lef)
			}
		})
	}
}

// ~2 ns/op
func Benchmark_Vector_Bounds_default(b *testing.B) {
	var vec *Vector
	{
		vec = tesVec()
	}

	{
		tesUpd(vec)
	}

	for b.Loop() {
		vec.Bounds()
	}
}

// ~4 ns/op
func Benchmark_Vector_Bounds_argument(b *testing.B) {
	var vec *Vector
	{
		vec = tesVec()
	}

	{
		tesUpd(vec)
	}

	for b.Loop() {
		vec.Bounds(vec.Charax().Fos)
	}
}

type screen struct{ top, rig, bot, lef int }
