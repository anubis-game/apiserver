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

// testUpd creates a Vector that is 21 nodes long and contains 57 hidden nodes
// at a player size of 785 points, given the test *Vector returned by tesVec().
func tesUpd(vec *Vector) {
	//
	//     +---------HR
	//     │
	//     │
	//     |
	//     T
	//
	//     {X: 1304, Y: 1303}, // H
	//     {X: 1284, Y: 1303},
	//     {X: 1264, Y: 1303},
	//     {X: 1244, Y: 1303},
	//     {X: 1224, Y: 1303},
	//     {X: 1204, Y: 1303},
	//     {X: 1184, Y: 1303},
	//     {X: 1164, Y: 1303},
	//     {X: 1144, Y: 1303},
	//     {X: 1124, Y: 1303},
	//     {X: 1104, Y: 1303},
	//     {X: 1084, Y: 1303},
	//     {X: 1064, Y: 1303},
	//     {X: 1044, Y: 1303},
	//     {X: 1024, Y: 1303},
	//     {X: 1005, Y: 1299},
	//     {X: 1000, Y: 1280},
	//     {X: 1000, Y: 1260},
	//     {X: 1000, Y: 1240},
	//     {X: 1000, Y: 1220},
	//     {X: 1000, Y: 1215}, // T
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
