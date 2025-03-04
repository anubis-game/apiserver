package vector

import (
	"slices"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_Occupy_initial(t *testing.T) {
	var vec *Vector
	{
		vec = tesVec()
	}

	var act []matrix.Coordinate
	{
		act = vec.Oclist(matrix.Partition{X: 896, Y: 896})
	}

	var exp []matrix.Coordinate
	{
		exp = []matrix.Coordinate{
			{X: 1000, Y: 1000},
		}
	}

	if len(vec.ocl) != 1 {
		t.Fatalf("expected %#v got %#v", 1, len(vec.ocl))
	}
	if !slices.Equal(act, exp) {
		t.Fatalf("expected %#v got %#v", exp, act)
	}

	for range 4 {
		vec.Update(int(Si/Li), vec.mot.Qdr, vec.mot.Agl, Nrm)
	}

	{
		act = vec.Oclist(matrix.Partition{X: 896, Y: 896})
	}

	{
		exp = []matrix.Coordinate{
			{X: 1000, Y: 1005},
			{X: 1000, Y: 1010},
			{X: 1000, Y: 1015},
			{X: 1000, Y: 1020},
		}
	}

	if len(vec.ocl) != 1 {
		t.Fatalf("expected %#v got %#v", 1, len(vec.ocl))
	}
	if !slices.Equal(act, exp) {
		t.Fatalf("expected %#v got %#v", exp, act)
	}
}
