package engine

import (
	"testing"

	"github.com/anubis-game/apiserver/pkg/vector"
)

func Test_Engine_race(t *testing.T) {
	var eng *Engine
	{
		eng = &Engine{
			rac: make([]byte, 6),
		}
	}

	var uid byte
	{
		uid = 0x5
	}

	// A new *Engine type starts out without any racing information. Therefore the
	// zero bytes.

	if eng.rac[uid] != 0x0 {
		t.Fatalf("expected %#v got %#v", 0x0, eng.rac[uid])
	}

	{
		eng.race(uid)
	}

	// The first call to Engine.race() must switch to racing mode.

	if eng.rac[uid] != vector.Rcn {
		t.Fatalf("expected %#v got %#v", vector.Rcn, eng.rac[uid])
	}

	{
		eng.race(uid)
	}

	// Further calls to Engine.race() must alternative between normal and racing
	// speed.

	if eng.rac[uid] != vector.Nrm {
		t.Fatalf("expected %#v got %#v", vector.Nrm, eng.rac[uid])
	}

	{
		eng.race(uid)
	}

	if eng.rac[uid] != vector.Rcn {
		t.Fatalf("expected %#v got %#v", vector.Rcn, eng.rac[uid])
	}

	{
		eng.race(uid)
	}

	if eng.rac[uid] != vector.Nrm {
		t.Fatalf("expected %#v got %#v", vector.Nrm, eng.rac[uid])
	}
}

// ~2 ns/op
func Benchmark_Engine_race(b *testing.B) {
	var eng *Engine
	{
		eng = &Engine{
			rac: make([]byte, 6),
		}
	}

	var uid byte
	{
		uid = 0x5
	}

	for b.Loop() {
		eng.race(uid)
	}
}
