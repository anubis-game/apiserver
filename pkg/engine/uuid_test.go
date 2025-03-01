package engine

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/router"
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

		{
			eng.uuid(router.Uuid{Uid: u, Jod: router.Join, Wal: wal})
		}

		if eng.mvc[u].Load() == nil {
			panic(fmt.Sprintf("expected %#v got %#v", "pointer", nil))
		}

		if !bytes.Contains(eng.fuw[u], wal.Bytes()) {
			panic(fmt.Sprintf("expected %#v got %#v", wal, eng.fuw[u]))
		}
	}
}

// ~29,600 ns/op, 38 allocs/op
func Benchmark_Engine_uuid(b *testing.B) {
	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
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
	})
}
