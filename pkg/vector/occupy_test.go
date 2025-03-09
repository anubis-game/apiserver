package vector

import (
	"testing"
)

func Test_Vector_Occupy(t *testing.T) {
	var vec *Vector
	{
		vec = tesVec()
	}

	var atp int
	var arg int
	var abt int
	var alf int
	{
		atp = vec.otp
		arg = vec.org
		abt = vec.obt
		alf = vec.olf
	}

	// The initial test Vector stretches across 1 partition.
	//
	//       |
	//       |
	//       |   p
	//       |
	//     x +-------------
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
		atp = vec.otp
		arg = vec.org
		abt = vec.obt
		alf = vec.olf
	}

	// The updated test Vector stretches across 4x2 partitions.
	//
	//       |
	//       |   p p p p
	//       |   p p p p
	//       |
	//     x +-------------
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
}
