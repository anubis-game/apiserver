package vector

import (
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_Occupy_pt1(t *testing.T) {
	var vec *Vector
	{
		vec = tesVec()
	}

	var atp int
	var arg int
	var abt int
	var alf int
	{
		atp, arg, abt, alf = vec.Occupy()
	}

	// The initial test Vector stretches across 1 partition.
	//
	//       |
	//       |
	//       |
	//       |   p
	//       |
	//     x +----------------
	//       y
	//

	var etp int
	var erg int
	var ebt int
	var elf int
	{
		etp = 896
		erg = 896
		ebt = 896
		elf = 896
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

	{
		tesUpd(vec)
	}

	{
		atp, arg, abt, alf = vec.Occupy()
	}

	// The updated test Vector stretches across 4x2 partitions, because it grew
	// and moved towards the east.
	//
	//       |
	//       |
	//       |   p  p  p  p
	//       |   p  p  p  p
	//       |
	//     x +----------------
	//       y
	//

	{
		etp = 1280
		erg = 1280
		ebt = 1152
		elf = 896
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

	// Move the test Vector further out towards the east without growing.

	for range 30 {
		vec.Update(0, 0x2, 0x0, Nrm)
	}

	{
		atp, arg, abt, alf = vec.Occupy()
	}

	// The updated test Vector stretches across 4x1 partitions, because the tail
	// came up as the Vector moved out further east.
	//
	//       |
	//       |
	//       |
	//       |   p  p  p  p
	//       |
	//     x +----------------
	//       y
	//

	{
		etp = 1280
		erg = 1408
		ebt = 1280
		elf = 1024
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

func Test_Vector_Occupy_pt8(t *testing.T) {
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
			Uid: 0x5,
		})
	}

	var atp int
	var arg int
	var abt int
	var alf int
	{
		atp, arg, abt, alf = vec.Occupy(matrix.Pt8)
	}

	// The initial test Vector stretches across 9 partitions.
	//
	//       |
	//       |   p  p  p
	//       |   p  p  p
	//       |   p  p  p
	//       |
	//     x +----------------
	//       y
	//

	var etp int
	var erg int
	var ebt int
	var elf int
	{
		etp = 4096
		erg = 3072
		ebt = 2048
		elf = 1024
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

	{
		tesUpd(vec)
	}

	{
		atp, arg, abt, alf = vec.Occupy(matrix.Pt8)
	}

	// The updated test Vector stretches across 4x3 partitions.
	//
	//       |
	//       |   p  p  p  p
	//       |   p  p  p  p
	//       |   p  p  p  p
	//       |
	//     x +----------------
	//       y
	//

	{
		etp = 5120
		erg = 4096
		ebt = 3072
		elf = 1024
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

	// Move the test Vector further out towards the east without growing.

	for range 100 {
		vec.Update(0, 0x2, 0x0, Nrm)
	}

	{
		atp, arg, abt, alf = vec.Occupy(matrix.Pt8)
	}

	// The updated test Vector stretches across 4x1 partitions, because the tail
	// came up as the Vector moved out further east.
	//
	//       |
	//       |   p  p  p
	//       |   p  p  p
	//       |   p  p  p
	//       |
	//     x +----------------
	//       y
	//

	{
		etp = 5120
		erg = 4096
		ebt = 3072
		elf = 2048
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
