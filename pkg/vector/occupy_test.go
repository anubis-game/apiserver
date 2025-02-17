package vector

import (
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Test_Vector_Occupy(t *testing.T) {
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

	if len(vec.Occupy().Prt) != 12 {
		t.Fatalf("expected %#v got %#v", 12, len(vec.Occupy().Prt))
	}

	{
		vec.Occupy().Prt = nil
	}

	if len(vec.Occupy().Prt) != 0 {
		t.Fatalf("expected %#v got %#v", 0, len(vec.Occupy().Prt))
	}
}
