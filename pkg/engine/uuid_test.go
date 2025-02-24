package engine

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/address"
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/envvar"
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
			Don: make(<-chan struct{}),
			Env: envvar.Env{
				EngineCapacity: 50,
			},
			Log: logger.Fake(),
		})
	}

	{
		go fil.Daemon()
	}

	var eng *Engine
	{
		eng = &Engine{
			buf: xsync.NewMapOf[byte, []byte](),
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
		eng.uuid(router.Packet{
			Cli: client.New(client.Config{
				Wal: wal,
			}),
			Uid: uid,
		})
	}

	if eng.mem.ply.Size() != 1 {
		t.Fatalf("expected %#v got %#v", 1, eng.mem.ply.Size())
	}

	ply, _ := eng.mem.ply.Load(uid)
	if !address.Equal(ply.Cli.Wallet(), wal) {
		t.Fatalf("expected %#v got %#v", wal, ply.Cli.Wallet())
	}
}

// ~8,000 ns/op, 61 allocs/op
func Benchmark_Engine_uuid(b *testing.B) {
	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		var fil *filler.Filler
		{
			fil = filler.New(filler.Config{
				Don: make(<-chan struct{}),
				Env: envvar.Env{
					EngineCapacity: 50,
				},
				Log: logger.Fake(),
			})
		}

		{
			go fil.Daemon()
		}

		var eng *Engine
		{
			eng = &Engine{
				buf: xsync.NewMapOf[byte, []byte](),
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
			eng.uuid(router.Packet{
				Cli: client.New(client.Config{
					Wal: wal,
				}),
				Uid: uid,
			})
		}
	})
}
