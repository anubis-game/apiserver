package engine

import (
	"bytes"
	"testing"

	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/ethereum/go-ethereum/common"
)

func Test_Engine_uuid(t *testing.T) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	if eng.uni.Length() != 250 {
		t.Fatalf("expected %#v got %#v", 250, eng.uni.Length())
	}

	for u := range eng.uni.Length() {
		var wal common.Address
		{
			wal = tesWal(u)
		}

		cli := eng.ply.cli[u]
		if cli != nil {
			t.Fatalf("expected %#v got %#v", nil, cli)
		}

		{
			eng.uuid(router.Uuid{Uid: u, Jod: router.Join, Wal: wal, Cli: make(chan<- []byte)})
		}

		cli = eng.ply.cli[u]
		if cli == nil {
			t.Fatalf("expected %T got %#v", make(chan<- []byte), nil)
		}

		vec := eng.mem.vec[u]
		if vec == nil {
			t.Fatalf("expected %T got %#v", &vector.Vector{}, nil)
		}

		if !bytes.Contains(eng.ply.buf[u], wal.Bytes()) {
			t.Fatalf("expected %#v got %#v", wal, eng.ply.buf[u])
		}
	}
}

// ~2,400 ns/op, 29 allocs/op
func Benchmark_Engine_uuid(b *testing.B) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	var uid byte
	{
		uid = 0x5
	}

	{
		eng.uuid(router.Uuid{Uid: 0x0, Jod: router.Join, Wal: tesWal(0x0), Cli: make(chan<- []byte)})
		eng.uuid(router.Uuid{Uid: 0x1, Jod: router.Join, Wal: tesWal(0x1), Cli: make(chan<- []byte)})
		eng.uuid(router.Uuid{Uid: 0x2, Jod: router.Join, Wal: tesWal(0x2), Cli: make(chan<- []byte)})
		eng.uuid(router.Uuid{Uid: 0x3, Jod: router.Join, Wal: tesWal(0x3), Cli: make(chan<- []byte)})
	}

	for b.Loop() {
		eng.uuid(router.Uuid{Uid: uid, Jod: router.Join, Wal: tesWal(uid), Cli: make(chan<- []byte)})
	}
}
