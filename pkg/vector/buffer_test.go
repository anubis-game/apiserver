package vector

import (
	"slices"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func Test_Vector_Buffer_Change(t *testing.T) {
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
			Uid: 0x5,
		})
	}

	var top int
	var rig int
	var bot int
	var lef int
	var old object.Object
	var new object.Object
	{
		top = 384
		rig = 256
		bot = 0
		lef = 0
		old = object.Object{}
		new = object.Object{}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 6 {
		t.Fatalf("expected %#v got %#v", 6, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 256}, []object.Object{
		{X: 100, Y: 300},
		{X: 100, Y: 350},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 384}, []object.Object{
		{X: 100, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 350, Y: 400}) // R
	}

	{
		top = 384
		rig = 256
		bot = 128
		lef = 0
		old = object.Object{}
		new = object.Object{}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 5 {
		t.Fatalf("expected %#v got %#v", 5, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150},
		{X: 100, Y: 200},
		{X: 100, Y: 250},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 256}, []object.Object{
		{X: 100, Y: 300},
		{X: 100, Y: 350},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 384}, []object.Object{
		{X: 100, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 400, Y: 400}) // R
	}

	{
		top = 384
		rig = 384
		bot = 128
		lef = 0
		old = object.Object{}
		new = object.Object{X: 384, Y: 384}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 6 {
		t.Fatalf("expected %#v got %#v", 6, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 200},
		{X: 100, Y: 250},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 256}, []object.Object{
		{X: 100, Y: 300},
		{X: 100, Y: 350},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 384}, []object.Object{
		{X: 100, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 384, Y: 384}, []object.Object{
		{X: 400, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 450, Y: 400}) // R
	}

	{
		top = 384
		rig = 384
		bot = 128
		lef = 0
		old = object.Object{}
		new = object.Object{}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 6 {
		t.Fatalf("expected %#v got %#v", 6, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 250},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 256}, []object.Object{
		{X: 100, Y: 300},
		{X: 100, Y: 350},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 384}, []object.Object{
		{X: 100, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 384, Y: 384}, []object.Object{
		{X: 400, Y: 400},
		{X: 450, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 500, Y: 400}) // R
	}

	{
		top = 384
		rig = 384
		bot = 256
		lef = 0
		old = object.Object{X: 0, Y: 128}
		new = object.Object{}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 5 {
		t.Fatalf("expected %#v got %#v", 5, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 256}, []object.Object{
		{X: 100, Y: 300},
		{X: 100, Y: 350},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 384}, []object.Object{
		{X: 100, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 384, Y: 384}, []object.Object{
		{X: 400, Y: 400},
		{X: 450, Y: 400},
		{X: 500, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 550, Y: 400}) // R
	}

	{
		top = 384
		rig = 512
		bot = 256
		lef = 0
		old = object.Object{}
		new = object.Object{X: 512, Y: 384}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 6 {
		t.Fatalf("expected %#v got %#v", 6, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 256}, []object.Object{
		{X: 100, Y: 350},
	})
	musBuf(t, vec, object.Object{X: 0, Y: 384}, []object.Object{
		{X: 100, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 384, Y: 384}, []object.Object{
		{X: 400, Y: 400},
		{X: 450, Y: 400},
		{X: 500, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 512, Y: 384}, []object.Object{
		{X: 550, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 600, Y: 400}) // R
	}

	{
		top = 384
		rig = 512
		bot = 384
		lef = 0
		old = object.Object{X: 0, Y: 256}
		new = object.Object{}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 5 {
		t.Fatalf("expected %#v got %#v", 5, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 384}, []object.Object{
		{X: 100, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 384, Y: 384}, []object.Object{
		{X: 400, Y: 400},
		{X: 450, Y: 400},
		{X: 500, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 512, Y: 384}, []object.Object{
		{X: 550, Y: 400},
		{X: 600, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 650, Y: 400}) // R
	}

	{
		top = 384
		rig = 640
		bot = 384
		lef = 128
		old = object.Object{X: 0, Y: 384}
		new = object.Object{X: 640, Y: 384}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 5 {
		t.Fatalf("expected %#v got %#v", 5, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 150, Y: 400},
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 384, Y: 384}, []object.Object{
		{X: 400, Y: 400},
		{X: 450, Y: 400},
		{X: 500, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 512, Y: 384}, []object.Object{
		{X: 550, Y: 400},
		{X: 600, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 640, Y: 384}, []object.Object{
		{X: 650, Y: 400},
	})

	{
		vec.Rotate(object.Object{X: 700, Y: 400}) // R
	}

	{
		top = 384
		rig = 640
		bot = 384
		lef = 128
		old = object.Object{}
		new = object.Object{}
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
	if vec.occ.Old != old {
		t.Fatalf("expected %#v got %#v", old, vec.occ.Old)
	}
	if vec.occ.New != new {
		t.Fatalf("expected %#v got %#v", new, vec.occ.New)
	}

	if len(vec.buf) != 5 {
		t.Fatalf("expected %#v got %#v", 5, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 128, Y: 384}, []object.Object{
		{X: 200, Y: 400},
		{X: 250, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 256, Y: 384}, []object.Object{
		{X: 300, Y: 400},
		{X: 350, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 384, Y: 384}, []object.Object{
		{X: 400, Y: 400},
		{X: 450, Y: 400},
		{X: 500, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 512, Y: 384}, []object.Object{
		{X: 550, Y: 400},
		{X: 600, Y: 400},
	})
	musBuf(t, vec, object.Object{X: 640, Y: 384}, []object.Object{
		{X: 650, Y: 400},
		{X: 700, Y: 400},
	})
}

func Test_Vector_Buffer_Duplicate_Coordinates(t *testing.T) {
	//
	//     +---------+
	//     R         │
	//     H         │
	//     |         │
	//     T---------+
	//
	var vec *Vector
	{
		vec = New(Config{
			Obj: []object.Object{
				{X: 100, Y: 100}, // T
				{X: 100, Y: 150},
				{X: 100, Y: 200},
				{X: 150, Y: 200},
				{X: 200, Y: 200},
				{X: 200, Y: 150},
				{X: 200, Y: 100},
				{X: 150, Y: 100},
				{X: 100, Y: 100},
				{X: 100, Y: 150}, // H
			},
			Uid: 0x5,
		})
	}

	var top int
	var rig int
	var bot int
	var lef int
	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 4 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 1
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150}, // 2
		{X: 100, Y: 200}, // 3
		{X: 100, Y: 150}, // 10
	})
	musBuf(t, vec, object.Object{X: 128, Y: 0}, []object.Object{
		{X: 200, Y: 100}, // 7
		{X: 150, Y: 100}, // 8
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 150, Y: 200}, // 4
		{X: 200, Y: 200}, // 5
		{X: 200, Y: 150}, // 6
	})

	{
		vec.Rotate(object.Object{X: 100, Y: 200}) // R
	}

	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 4 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150}, // 2
		{X: 100, Y: 200}, // 3
		{X: 100, Y: 150}, // 10
		{X: 100, Y: 200}, // 11
	})
	musBuf(t, vec, object.Object{X: 128, Y: 0}, []object.Object{
		{X: 200, Y: 100}, // 7
		{X: 150, Y: 100}, // 8
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 150, Y: 200}, // 4
		{X: 200, Y: 200}, // 5
		{X: 200, Y: 150}, // 6
	})

	{
		vec.Rotate(object.Object{X: 150, Y: 200}) // R
	}

	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 4 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 200}, // 3
		{X: 100, Y: 150}, // 10
		{X: 100, Y: 200}, // 11
	})
	musBuf(t, vec, object.Object{X: 128, Y: 0}, []object.Object{
		{X: 200, Y: 100}, // 7
		{X: 150, Y: 100}, // 8
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 150, Y: 200}, // 4
		{X: 200, Y: 200}, // 5
		{X: 200, Y: 150}, // 6
		{X: 150, Y: 200}, // 12
	})

	{
		vec.Rotate(object.Object{X: 200, Y: 200}) // R
	}

	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 4 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150}, // 10
		{X: 100, Y: 200}, // 11
	})
	musBuf(t, vec, object.Object{X: 128, Y: 0}, []object.Object{
		{X: 200, Y: 100}, // 7
		{X: 150, Y: 100}, // 8
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 150, Y: 200}, // 4
		{X: 200, Y: 200}, // 5
		{X: 200, Y: 150}, // 6
		{X: 150, Y: 200}, // 12
		{X: 200, Y: 200}, // 13
	})

	{
		vec.Rotate(object.Object{X: 200, Y: 150}) // R
	}

	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 4 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150}, // 10
		{X: 100, Y: 200}, // 11
	})
	musBuf(t, vec, object.Object{X: 128, Y: 0}, []object.Object{
		{X: 200, Y: 100}, // 7
		{X: 150, Y: 100}, // 8
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 200, Y: 200}, // 5
		{X: 200, Y: 150}, // 6
		{X: 150, Y: 200}, // 12
		{X: 200, Y: 200}, // 13
		{X: 200, Y: 150}, // 14
	})

	{
		vec.Shrink()
	}

	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 4 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150}, // 10
		{X: 100, Y: 200}, // 11
	})
	musBuf(t, vec, object.Object{X: 128, Y: 0}, []object.Object{
		{X: 200, Y: 100}, // 7
		{X: 150, Y: 100}, // 8
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 200, Y: 150}, // 6
		{X: 150, Y: 200}, // 12
		{X: 200, Y: 200}, // 13
		{X: 200, Y: 150}, // 14
	})

	{
		vec.Shrink()
	}

	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 4 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150}, // 10
		{X: 100, Y: 200}, // 11
	})
	musBuf(t, vec, object.Object{X: 128, Y: 0}, []object.Object{
		{X: 200, Y: 100}, // 7
		{X: 150, Y: 100}, // 8
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 150, Y: 200}, // 12
		{X: 200, Y: 200}, // 13
		{X: 200, Y: 150}, // 14
	})

	{
		vec.Shrink()
		vec.Shrink()
	}

	{
		top = 128
		rig = 128
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	if len(vec.buf) != 3 {
		t.Fatalf("expected %#v got %#v", 4, len(vec.buf))
	}
	musBuf(t, vec, object.Object{X: 0, Y: 0}, []object.Object{
		{X: 100, Y: 100}, // 9
	})
	musBuf(t, vec, object.Object{X: 0, Y: 128}, []object.Object{
		{X: 100, Y: 150}, // 10
		{X: 100, Y: 200}, // 11
	})
	musBuf(t, vec, object.Object{X: 128, Y: 128}, []object.Object{
		{X: 150, Y: 200}, // 12
		{X: 200, Y: 200}, // 13
		{X: 200, Y: 150}, // 14
	})
}

func Test_Vector_Buffer_No_Change(t *testing.T) {
	//
	//     +---------+
	//     │    R    │
	//     │    H    │
	//     |    │    │
	//     T    +----+
	//
	var vec *Vector
	{
		vec = New(Config{
			Obj: []object.Object{
				{X: 100, Y: 100}, // T
				{X: 100, Y: 105},
				{X: 100, Y: 110},
				{X: 100, Y: 115},
				{X: 100, Y: 120},
				{X: 105, Y: 120},
				{X: 110, Y: 120},
				{X: 115, Y: 120},
				{X: 120, Y: 120},
				{X: 120, Y: 115},
				{X: 120, Y: 110},
				{X: 120, Y: 105},
				{X: 120, Y: 100},
				{X: 115, Y: 100},
				{X: 110, Y: 100},
				{X: 110, Y: 105},
				{X: 110, Y: 110}, // H
			},
		})
	}

	var top int
	var rig int
	var bot int
	var lef int
	{
		top = 0
		rig = 0
		bot = 0
		lef = 0
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}

	{
		vec.Rotate(object.Object{X: 110, Y: 115}) // R
	}

	if vec.occ.Top != top {
		t.Fatalf("expected %#v got %#v", top, vec.occ.Top)
	}
	if vec.occ.Rig != rig {
		t.Fatalf("expected %#v got %#v", rig, vec.occ.Rig)
	}
	if vec.occ.Bot != bot {
		t.Fatalf("expected %#v got %#v", bot, vec.occ.Bot)
	}
	if vec.occ.Lef != lef {
		t.Fatalf("expected %#v got %#v", lef, vec.occ.Lef)
	}
}

func musBuf(t *testing.T, vec *Vector, prt object.Object, obj []object.Object) {
	act := vec.buf[prt]
	exp := objByt(vec.uid, obj)

	if act[1] == 0x0 {
		t.Fatalf("expected %#v got %#v", "non-zero ID bytes", act[1])
	}
	if !slices.Equal(act, exp) {
		t.Fatalf("expected %#v got %#v", exp, act)
	}
}

func objByt(uid byte, obj []object.Object) []byte {
	var buf []byte
	{
		buf = []byte{
			byte(schema.Body),
			uid,
			byte(len(obj)),
		}
	}

	for _, x := range obj {
		byt := x.Byt()
		buf = append(buf, byt[:]...)
	}

	return buf
}
