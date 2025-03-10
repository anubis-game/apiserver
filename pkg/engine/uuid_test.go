package engine

import (
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

		act := eng.act[u]
		if act == true {
			t.Fatalf("expected %#v got %#v", false, true)
		}

		{
			eng.uuid(router.Uuid{Uid: u, Jod: router.Join, Wal: wal})
		}

		act = eng.act[u]
		if act == false {
			t.Fatalf("expected %#v got %#v", true, false)
		}

		vec, _ := eng.mem.vec.Load(u)
		if vec == nil {
			t.Fatalf("expected %T got %#v", &vector.Vector{}, nil)
		}

		// if !bytes.Contains(eng.fuw[u], wal.Bytes()) {
		// 	panic(fmt.Sprintf("expected %#v got %#v", wal, eng.fuw[u]))
		// }
	}
}

// ~2,265 ns/op, 28 allocs/op
func Benchmark_Engine_uuid(b *testing.B) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	var uid byte
	{
		uid = 0x5
	}

	for b.Loop() {
		eng.uuid(router.Uuid{Uid: uid, Jod: router.Join, Wal: tesWal(uid)})
	}
}
