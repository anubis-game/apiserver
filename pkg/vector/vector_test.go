package vector

import (
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_New(t *testing.T) {
	var hea matrix.Coordinate
	{
		hea = matrix.Coordinate{
			X: 3000,
			Y: 4000,
		}
	}

	var mot Motion
	{
		mot = Motion{
			Qdr: 0x1,
			Agl: 0x5,
		}
	}

	var vec *Vector
	{
		vec = New(Config{
			Hea: hea,
			Mot: mot,
		})
	}

	if vec.len != 1 {
		t.Fatalf("expected %#v got %#v", 1, vec.len)
	}
	if vec.hea.crd != hea {
		t.Fatalf("expected %#v got %#v", hea, vec.hea.crd)
	}
	if vec.tai.crd != hea {
		t.Fatalf("expected %#v got %#v", hea, vec.tai.crd)
	}
	if vec.mot.Qdr != mot.Qdr {
		t.Fatalf("expected %#v got %#v", mot.Qdr, vec.mot.Qdr)
	}
	if vec.mot.Agl != mot.Agl {
		t.Fatalf("expected %#v got %#v", mot.Agl, vec.mot.Agl)
	}
}

func tesVec() *Vector {
	return New(Config{
		Hea: matrix.Coordinate{
			X: 1000,
			Y: 1000,
		},
		Mot: Motion{
			Qdr: 0x1,
			Agl: 0x0,
		},
	})
}

// testUpd creates a Vector that is 78 segments long at a size of 785 points,
// given the test *Vector returned by tesVec().
func tesUpd(vec *Vector) {
	//
	//     +---------HR
	//     │
	//     │
	//     |
	//     T
	//
	//     {X: 4004, Y: 4003}, // H
	//     {X: 3999, Y: 4003},
	//     ...
	//     {X: 1014, Y: 4003},
	//     {X: 1009, Y: 4002},
	//     {X: 1005, Y: 3099}, // +
	//     {X: 1002, Y: 3095},
	//     {X: 1000, Y: 3090},
	//     ...
	//     {X: 1000, Y: 1010},
	//     {X: 1000, Y: 1005}, // T
	//

	{
		vec.Update(190, 0x1, 0x0, Nrm)
	}

	for range 57 {
		vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
	}

	for range 62 {
		vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
	}
}
