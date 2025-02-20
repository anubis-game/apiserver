package engine

import (
	"fmt"
	"testing"
	"time"

	"github.com/anubis-game/apiserver/pkg/address"
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
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
			buf: xsync.NewMapOf[[2]byte, []byte](),
			fil: fil,
			lkp: &lookup{
				nrg: xsync.NewMapOf[object.Object, map[object.Object]struct{}](),
				ply: xsync.NewMapOf[object.Object, map[[2]byte]struct{}](),
			},
			mem: &memory{
				nrg: xsync.NewMapOf[object.Object, *energy.Energy](),
				ply: xsync.NewMapOf[[2]byte, *player.Player](),
			},
		}
	}

	var uid [2]byte
	{
		uid = [2]byte{0x0, 0x5}
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

func Test_Engine_timCap(t *testing.T) {
	testCases := []struct {
		ply int
		cpu int
		tim time.Duration
	}{
		//
		// 1 CPU
		//

		// Case 000
		{
			ply: 0,
			cpu: 1,
			tim: 24 * time.Millisecond,
		},
		// Case 001
		{
			ply: 1,
			cpu: 1,
			tim: 24 * time.Millisecond,
		},
		// Case 002
		{
			ply: 5,
			cpu: 1,
			tim: 4800 * time.Microsecond,
		},
		// Case 003
		{
			ply: 80,
			cpu: 1,
			tim: 300 * time.Microsecond,
		},
		// Case 004
		{
			ply: 100,
			cpu: 1,
			tim: 240 * time.Microsecond,
		},
		// Case 005
		{
			ply: 250,
			cpu: 1,
			tim: 96 * time.Microsecond,
		},
		// Case 006
		{
			ply: 500,
			cpu: 1,
			tim: 48 * time.Microsecond,
		},

		//
		// 2 CPUs
		//

		// Case 007
		{
			ply: 0,
			cpu: 2,
			tim: 12 * time.Millisecond,
		},
		// Case 008
		{
			ply: 1,
			cpu: 2,
			tim: 12 * time.Millisecond,
		},
		// Case 009
		{
			ply: 5,
			cpu: 2,
			tim: 2400 * time.Microsecond,
		},
		// Case 010
		{
			ply: 80,
			cpu: 2,
			tim: 150 * time.Microsecond,
		},
		// Case 011
		{
			ply: 100,
			cpu: 2,
			tim: 120 * time.Microsecond,
		},
		// Case 012
		{
			ply: 250,
			cpu: 2,
			tim: 48 * time.Microsecond,
		},
		// Case 013
		{
			ply: 500,
			cpu: 2,
			tim: 24 * time.Microsecond,
		},

		//
		// 4 CPUs
		//

		// Case 014
		{
			ply: 0,
			cpu: 4,
			tim: 6 * time.Millisecond,
		},
		// Case 015
		{
			ply: 1,
			cpu: 4,
			tim: 6 * time.Millisecond,
		},
		// Case 016
		{
			ply: 5,
			cpu: 4,
			tim: 1200 * time.Microsecond,
		},
		// Case 017
		{
			ply: 80,
			cpu: 4,
			tim: 75 * time.Microsecond,
		},
		// Case 018
		{
			ply: 100,
			cpu: 4,
			tim: 60 * time.Microsecond,
		},
		// Case 019
		{
			ply: 250,
			cpu: 4,
			tim: 24 * time.Microsecond,
		},
		// Case 020
		{
			ply: 500,
			cpu: 4,
			tim: 12 * time.Microsecond,
		},

		//
		// 8 CPUs
		//

		// Case 021
		{
			ply: 0,
			cpu: 8,
			tim: 3 * time.Millisecond,
		},
		// Case 022
		{
			ply: 1,
			cpu: 8,
			tim: 3 * time.Millisecond,
		},
		// Case 023
		{
			ply: 5,
			cpu: 8,
			tim: 600 * time.Microsecond,
		},
		// Case 024
		{
			ply: 80,
			cpu: 8,
			tim: 37500 * time.Nanosecond,
		},
		// Case 025
		{
			ply: 100,
			cpu: 8,
			tim: 30 * time.Microsecond,
		},
		// Case 026
		{
			ply: 250,
			cpu: 8,
			tim: 12 * time.Microsecond,
		},
		// Case 027
		{
			ply: 500,
			cpu: 8,
			tim: 6 * time.Microsecond,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tim := timCap(tc.ply, tc.cpu)

			if tim != tc.tim {
				t.Fatalf("expected %s got %s", tc.tim, tim)
			}
		})
	}
}

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
				buf: xsync.NewMapOf[[2]byte, []byte](),
				fil: fil,
				lkp: &lookup{
					nrg: xsync.NewMapOf[object.Object, map[object.Object]struct{}](),
					ply: xsync.NewMapOf[object.Object, map[[2]byte]struct{}](),
				},
				mem: &memory{
					nrg: xsync.NewMapOf[object.Object, *energy.Energy](),
					ply: xsync.NewMapOf[[2]byte, *player.Player](),
				},
			}
		}

		var uid [2]byte
		{
			uid = [2]byte{0x0, 0x5}
		}

		var wal common.Address
		{
			wal = common.HexToAddress("0x0000000000000000000000000000000000000005")
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// ~4,400 ns/op
			eng.uuid(router.Packet{
				Cli: client.New(client.Config{
					Wal: wal,
				}),
				Uid: uid,
			})
		}
	})
}

func Benchmark_Engine_timCap(b *testing.B) {
	testCases := []struct {
		ply int
		cpu int
	}{
		//
		// 1 CPU
		//

		// Case 000, ~0.40 ns/op
		{
			ply: 0,
			cpu: 1,
		},
		// Case 001, ~0.40 ns/op
		{
			ply: 1,
			cpu: 1,
		},
		// Case 002, ~0.50 ns/op
		{
			ply: 5,
			cpu: 1,
		},
		// Case 003, ~0.50 ns/op
		{
			ply: 80,
			cpu: 1,
		},
		// Case 004, ~0.50 ns/op
		{
			ply: 500,
			cpu: 1,
		},

		//
		// 2 CPUs
		//

		// Case 005, ~0.40 ns/op
		{
			ply: 0,
			cpu: 2,
		},
		// Case 006, ~0.40 ns/op
		{
			ply: 1,
			cpu: 2,
		},
		// Case 007, ~0.60 ns/op
		{
			ply: 5,
			cpu: 2,
		},
		// Case 008, ~0.60 ns/op
		{
			ply: 80,
			cpu: 2,
		},
		// Case 009, ~0.60 ns/op
		{
			ply: 500,
			cpu: 2,
		},

		//
		// 4 CPUs
		//

		// Case 010, ~0.40 ns/op
		{
			ply: 0,
			cpu: 4,
		},
		// Case 011, ~0.40 ns/op
		{
			ply: 1,
			cpu: 4,
		},
		// Case 012, ~0.60 ns/op
		{
			ply: 5,
			cpu: 4,
		},
		// Case 013, ~0.60 ns/op
		{
			ply: 80,
			cpu: 4,
		},
		// Case 014, ~0.60 ns/op
		{
			ply: 500,
			cpu: 4,
		},

		//
		// 8 CPUs
		//

		// Case 015, ~0.40 ns/op
		{
			ply: 0,
			cpu: 8,
		},
		// Case 016, ~0.40 ns/op
		{
			ply: 1,
			cpu: 8,
		},
		// Case 017, ~0.60 ns/op
		{
			ply: 5,
			cpu: 8,
		},
		// Case 018, ~0.60 ns/op
		{
			ply: 80,
			cpu: 8,
		},
		// Case 019, ~0.60 ns/op
		{
			ply: 500,
			cpu: 8,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				timCap(tc.ply, tc.cpu)
			}
		})
	}
}
