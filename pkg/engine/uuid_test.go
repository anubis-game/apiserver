package engine

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/xh3b4sd/logger"
)

func Test_Engine_uuid(t *testing.T) {
	var fil *filler.Filler
	{
		fil = filler.New(filler.Config{
			Cap: 50,
			Don: make(<-chan struct{}),
			Log: logger.Fake(),
		})
	}

	{
		go fil.Daemon()
	}

	var eng *Engine
	{
		eng = &Engine{
			fbf: xsync.NewMapOf[byte, []byte](),
			fcn: make([]chan<- []byte, 6),
			fil: fil,
			lkp: &lookup{
				nrg: xsync.NewMapOf[object.Object, map[object.Object]struct{}](),
				ply: xsync.NewMapOf[object.Object, map[byte]struct{}](),
			},
			mem: &memory{
				nrg: xsync.NewMapOf[object.Object, *energy.Energy](),
				ply: xsync.NewMapOf[byte, *player.Player](),
			},
			uni: unique.New[common.Address, byte](),
		}
	}

	var uid byte
	{
		uid = 0x5
	}

	var wal common.Address
	{
		wal = common.HexToAddress("0x0000000000000000000000000000000000000005")
	}

	{
		eng.uuid(router.Uuid{Uid: uid, Jod: router.Join, Wal: wal})
	}

	if eng.mem.ply.Size() != 1 {
		t.Fatalf("expected %#v got %#v", 1, eng.mem.ply.Size())
	}

	ply, _ := eng.mem.ply.Load(uid)

	if !bytes.Contains(ply.Wallet(), wal.Bytes()) {
		t.Fatalf("expected %#v got %#v", wal, ply.Wallet())
	}
}

// ~4,300 ns/op, 49 allocs/op
func Benchmark_Engine_uuid(b *testing.B) {
	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		var fil *filler.Filler
		{
			fil = filler.New(filler.Config{
				Cap: 50,
				Don: make(<-chan struct{}),
				Log: logger.Fake(),
			})
		}

		{
			go fil.Daemon()
		}

		var eng *Engine
		{
			eng = &Engine{
				fbf: xsync.NewMapOf[byte, []byte](),
				fcn: make([]chan<- []byte, 6),
				fil: fil,
				lkp: &lookup{
					nrg: xsync.NewMapOf[object.Object, map[object.Object]struct{}](),
					ply: xsync.NewMapOf[object.Object, map[byte]struct{}](),
				},
				mem: &memory{
					nrg: xsync.NewMapOf[object.Object, *energy.Energy](),
					ply: xsync.NewMapOf[byte, *player.Player](),
				},
				uni: unique.New[common.Address, byte](),
			}
		}

		var uid byte
		{
			uid = 0x5
		}

		var wal common.Address
		{
			wal = common.HexToAddress("0x0000000000000000000000000000000000000005")
		}

		for b.Loop() {
			eng.uuid(router.Uuid{Uid: uid, Jod: router.Join, Wal: wal})
		}
	})
}
