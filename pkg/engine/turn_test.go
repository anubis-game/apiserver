package engine

import (
	"testing"

	"github.com/anubis-game/apiserver/pkg/router"
)

func Test_Engine_turn(t *testing.T) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	var uid byte
	{
		uid = 0x5
	}

	if eng.ply.tur[uid].Qdr != 0x0 {
		t.Fatalf("expected %#v got %#v", 0x0, eng.ply.tur[uid].Qdr)
	}
	if eng.ply.tur[uid].Agl != 0x0 {
		t.Fatalf("expected %#v got %#v", 0x0, eng.ply.tur[uid].Agl)
	}

	eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x80})

	if eng.ply.tur[uid].Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.ply.tur[uid].Qdr)
	}
	if eng.ply.tur[uid].Agl != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, eng.ply.tur[uid].Agl)
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x76})
	}

	if eng.ply.tur[uid].Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.ply.tur[uid].Qdr)
	}
	if eng.ply.tur[uid].Agl != 0x76 {
		t.Fatalf("expected %#v got %#v", 0x76, eng.ply.tur[uid].Agl)
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x3, Agl: 0xa})
	}

	if eng.ply.tur[uid].Qdr != 0x3 {
		t.Fatalf("expected %#v got %#v", 0x3, eng.ply.tur[uid].Qdr)
	}
	if eng.ply.tur[uid].Agl != 0xa {
		t.Fatalf("expected %#v got %#v", 0xa, eng.ply.tur[uid].Agl)
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x80})
	}

	if eng.ply.tur[uid].Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.ply.tur[uid].Qdr)
	}
	if eng.ply.tur[uid].Agl != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, eng.ply.tur[uid].Agl)
	}
}

// ~2 ns/op
func Benchmark_Engine_turn(b *testing.B) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	var uid byte
	{
		uid = 0x5
	}

	for b.Loop() {
		eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x76})
	}
}
