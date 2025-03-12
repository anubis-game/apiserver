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

	if eng.ply.qdr[uid] != 0x0 {
		t.Fatalf("expected %#v got %#v", 0x0, eng.ply.qdr[uid])
	}
	if eng.ply.agl[uid] != 0x0 {
		t.Fatalf("expected %#v got %#v", 0x0, eng.ply.agl[uid])
	}

	eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x80})

	if eng.ply.qdr[uid] != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.ply.qdr[uid])
	}
	if eng.ply.agl[uid] != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, eng.ply.agl[uid])
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x76})
	}

	if eng.ply.qdr[uid] != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.ply.qdr[uid])
	}
	if eng.ply.agl[uid] != 0x76 {
		t.Fatalf("expected %#v got %#v", 0x76, eng.ply.agl[uid])
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x3, Agl: 0xa})
	}

	if eng.ply.qdr[uid] != 0x3 {
		t.Fatalf("expected %#v got %#v", 0x3, eng.ply.qdr[uid])
	}
	if eng.ply.agl[uid] != 0xa {
		t.Fatalf("expected %#v got %#v", 0xa, eng.ply.agl[uid])
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x80})
	}

	if eng.ply.qdr[uid] != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.ply.qdr[uid])
	}
	if eng.ply.agl[uid] != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, eng.ply.agl[uid])
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
