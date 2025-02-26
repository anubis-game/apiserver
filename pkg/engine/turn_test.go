package engine

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/router"
)

func Test_Engine_turn(t *testing.T) {
	var eng *Engine
	{
		eng = &Engine{
			tur: make([]router.Turn, 6),
		}
	}

	var uid byte
	{
		uid = 0x5
	}

	if eng.tur[uid].Qdr != 0x0 {
		t.Fatalf("expected %#v got %#v", 0x0, eng.tur[uid].Qdr)
	}
	if eng.tur[uid].Agl != 0x0 {
		t.Fatalf("expected %#v got %#v", 0x0, eng.tur[uid].Agl)
	}

	eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x80})

	if eng.tur[uid].Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.tur[uid].Qdr)
	}
	if eng.tur[uid].Agl != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, eng.tur[uid].Agl)
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x76})
	}

	if eng.tur[uid].Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.tur[uid].Qdr)
	}
	if eng.tur[uid].Agl != 0x76 {
		t.Fatalf("expected %#v got %#v", 0x76, eng.tur[uid].Agl)
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x3, Agl: 0xa})
	}

	if eng.tur[uid].Qdr != 0x3 {
		t.Fatalf("expected %#v got %#v", 0x3, eng.tur[uid].Qdr)
	}
	if eng.tur[uid].Agl != 0xa {
		t.Fatalf("expected %#v got %#v", 0xa, eng.tur[uid].Agl)
	}

	{
		eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x80})
	}

	if eng.tur[uid].Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, eng.tur[uid].Qdr)
	}
	if eng.tur[uid].Agl != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, eng.tur[uid].Agl)
	}
}

// ~2.90 ns/op
func Benchmark_Engine_turn(b *testing.B) {
	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		var eng *Engine
		{
			eng = &Engine{
				tur: make([]router.Turn, 6),
			}
		}

		var uid byte
		{
			uid = 0x5
		}

		for b.Loop() {
			eng.turn(router.Turn{Uid: uid, Qdr: 0x1, Agl: 0x76})
		}
	})
}
