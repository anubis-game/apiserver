package engine

import (
	"bytes"
	"testing"

	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/ethereum/go-ethereum/common"
)

func Test_Engine_uuid(t *testing.T) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	if eng.uni.Length() != 249 {
		t.Fatalf("expected %d got %d", 249, eng.uni.Length())
	}

	for u := range eng.uni.Length() {
		var vec *vector.Vector
		{
			vec = tesVec(u)
		}

		var wal common.Address
		{
			wal = tesWal(u)
		}

		cli := eng.ply.cli[u]
		if cli != nil {
			t.Fatalf("expected %#v got %#v", nil, cli)
		}

		{
			eng.uuid(router.Uuid{Uid: u, Jod: router.Join, Wal: wal, Cli: make(chan<- []byte), Vec: vec})
		}

		cli = eng.ply.cli[u]
		if cli == nil {
			t.Fatalf("expected %T got %#v", make(chan<- []byte), nil)
		}

		act := eng.ply.buf[u]
		if act[22] != byte(schema.Body) {
			t.Fatalf("expected %#v got %#v", byte(schema.Body), act[22])
		}

		len := eng.ply.buf[u]
		if len[24] != 2 {
			t.Fatalf("expected %#v got %#v", 2, len[24])
		}

		vec = eng.mem.vec[u]
		if vec == nil {
			t.Fatalf("expected %T got %#v", &vector.Vector{}, nil)
		}

		if !bytes.Contains(eng.ply.buf[u], wal.Bytes()) {
			t.Fatalf("expected %#v got %#v", wal, eng.ply.buf[u])
		}
	}
}

// ~798 ns/op, 9 allocs/op
func Benchmark_Engine_uuid(b *testing.B) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	var uid byte
	{
		uid = 0x5
	}

	var vec *vector.Vector
	{
		vec = tesVec(uid)
	}

	{
		eng.uuid(router.Uuid{Uid: 0x0, Jod: router.Join, Wal: tesWal(0x0), Cli: make(chan<- []byte), Vec: tesVec(0x0)})
		eng.uuid(router.Uuid{Uid: 0x1, Jod: router.Join, Wal: tesWal(0x1), Cli: make(chan<- []byte), Vec: tesVec(0x1)})
		eng.uuid(router.Uuid{Uid: 0x2, Jod: router.Join, Wal: tesWal(0x2), Cli: make(chan<- []byte), Vec: tesVec(0x2)})
		eng.uuid(router.Uuid{Uid: 0x3, Jod: router.Join, Wal: tesWal(0x3), Cli: make(chan<- []byte), Vec: tesVec(0x3)})
	}

	for b.Loop() {
		eng.uuid(router.Uuid{Uid: uid, Jod: router.Join, Wal: tesWal(uid), Cli: make(chan<- []byte), Vec: vec})
	}
}
