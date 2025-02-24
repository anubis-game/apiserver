package engine

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/puzpuzpuz/xsync/v3"
)

func Test_Engine_race(t *testing.T) {
	var eng *Engine
	{
		eng = &Engine{
			mem: &memory{
				ply: xsync.NewMapOf[byte, *player.Player](),
			},
		}
	}

	var uid byte
	{
		uid = 0x5
	}

	var ply *player.Player
	{
		ply = &player.Player{
			Vec: vector.New(vector.Config{
				Mot: vector.Motion{
					Qdr: 0x1,
					Agl: 0x80,
					Vlc: vector.Nrm,
				},
				Obj: []object.Object{
					{X: 100, Y: 100}, // 0
					{X: 103, Y: 103}, // 1
					{X: 106, Y: 106}, // 2
					{X: 109, Y: 109}, // 3
					{X: 112, Y: 112}, // 4
				},
				Uid: uid,
			}),
		}
	}

	{
		eng.mem.ply.Store(uid, ply)
	}

	if ply.Vec.Motion().Get().Vlc != vector.Nrm {
		t.Fatalf("expected %#v got %#v", vector.Nrm, ply.Vec.Motion().Get().Vlc)
	}

	{
		eng.race(router.Packet{Uid: uid})
	}

	if ply.Vec.Motion().Get().Vlc != vector.Rcn {
		t.Fatalf("expected %#v got %#v", vector.Rcn, ply.Vec.Motion().Get().Vlc)
	}

	{
		eng.race(router.Packet{Uid: uid})
	}

	if ply.Vec.Motion().Get().Vlc != vector.Nrm {
		t.Fatalf("expected %#v got %#v", vector.Nrm, ply.Vec.Motion().Get().Vlc)
	}

	{
		eng.race(router.Packet{Uid: uid})
	}

	if ply.Vec.Motion().Get().Vlc != vector.Rcn {
		t.Fatalf("expected %#v got %#v", vector.Rcn, ply.Vec.Motion().Get().Vlc)
	}
}

// ~24.40 ns/op
func Benchmark_Engine_race(b *testing.B) {
	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		var eng *Engine
		{
			eng = &Engine{
				mem: &memory{
					ply: xsync.NewMapOf[byte, *player.Player](),
				},
			}
		}

		var uid byte
		{
			uid = 0x5
		}

		var ply *player.Player
		{
			ply = &player.Player{
				Vec: vector.New(vector.Config{
					Mot: vector.Motion{
						Qdr: 0x1,
						Agl: 0x80,
						Vlc: vector.Nrm,
					},
					Obj: []object.Object{
						{X: 100, Y: 100}, // 0
						{X: 103, Y: 103}, // 1
						{X: 106, Y: 106}, // 2
						{X: 109, Y: 109}, // 3
						{X: 112, Y: 112}, // 4
					},
					Uid: uid,
				}),
			}
		}

		{
			eng.mem.ply.Store(uid, ply)
		}

		for b.Loop() {
			eng.race(router.Packet{Uid: uid})
		}
	})
}
