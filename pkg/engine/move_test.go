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

func Test_Engine_move_quadrant_inside_range(t *testing.T) {
	testCases := []struct {
		b byte
	}{
		// Case 000
		{
			b: 1,
		},
		// Case 001
		{
			b: 2,
		},
		// Case 002
		{
			b: 3,
		},
		// Case 003
		{
			b: 4,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			r := tc.b-1 > 3
			if r {
				t.Fatalf("expected %#v got %#v", false, r)
			}
		})
	}
}

func Test_Engine_move_quadrant_outside_range(t *testing.T) {
	testCases := []struct {
		b byte
	}{
		// Case 000
		{
			b: 0,
		},
		// Case 001
		{
			b: 5,
		},
		// Case 002
		{
			b: 6,
		},
		// Case 003
		{
			b: 188,
		},
		// Case 004
		{
			b: 255,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			r := tc.b-1 > 3
			if !r {
				t.Fatalf("expected %#v got %#v", true, r)
			}
		})
	}
}

func Test_Engine_move(t *testing.T) {
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

	if ply.Vec.Motion().Get().Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, ply.Vec.Motion().Get().Qdr)
	}
	if ply.Vec.Motion().Get().Agl != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, ply.Vec.Motion().Get().Agl)
	}

	{
		eng.move(router.Packet{Uid: uid, Byt: []byte{0x1, 0x76}})
	}

	if ply.Vec.Motion().Get().Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, ply.Vec.Motion().Get().Qdr)
	}
	if ply.Vec.Motion().Get().Agl != 0x76 {
		t.Fatalf("expected %#v got %#v", 0x76, ply.Vec.Motion().Get().Agl)
	}

	{
		eng.move(router.Packet{Uid: uid, Byt: []byte{0x3, 0xa}})
	}

	if ply.Vec.Motion().Get().Qdr != 0x3 {
		t.Fatalf("expected %#v got %#v", 0x3, ply.Vec.Motion().Get().Qdr)
	}
	if ply.Vec.Motion().Get().Agl != 0xa {
		t.Fatalf("expected %#v got %#v", 0xa, ply.Vec.Motion().Get().Agl)
	}

	{
		eng.move(router.Packet{Uid: uid, Byt: []byte{0x1, 0x80}})
	}

	if ply.Vec.Motion().Get().Qdr != 0x1 {
		t.Fatalf("expected %#v got %#v", 0x1, ply.Vec.Motion().Get().Qdr)
	}
	if ply.Vec.Motion().Get().Agl != 0x80 {
		t.Fatalf("expected %#v got %#v", 0x80, ply.Vec.Motion().Get().Agl)
	}
}

// ~26.10 ns/op
func Benchmark_Engine_move(b *testing.B) {
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
			eng.move(router.Packet{Uid: uid, Byt: []byte{0x1, 0x76}})
		}
	})
}

func Benchmark_Engine_move_quadrant_one_condition(b *testing.B) {
	testCases := []struct {
		b byte
	}{
		// Case 000, ~1.00 ns/op
		{
			b: 0,
		},
		// Case 001, ~1.00 ns/op
		{
			b: 1,
		},
		// Case 002, ~1.00 ns/op
		{
			b: 2,
		},
		// Case 003, ~1.00 ns/op
		{
			b: 3,
		},
		// Case 004, ~1.00 ns/op
		{
			b: 4,
		},
		// Case 005, ~1.00 ns/op
		{
			b: 5,
		},
		// Case 006, ~1.00 ns/op
		{
			b: 6,
		},
		// Case 007, ~1.00 ns/op
		{
			b: 188,
		},
		// Case 008, ~1.00 ns/op
		{
			b: 255,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				_ = tc.b-1 > 3
			}
		})
	}
}

func Benchmark_Engine_move_quadrant_two_conditions(b *testing.B) {
	testCases := []struct {
		b byte
	}{
		// Case 000, ~1.00 ns/op
		{
			b: 0,
		},
		// Case 001, ~1.00 ns/op
		{
			b: 1,
		},
		// Case 002, ~1.00 ns/op
		{
			b: 2,
		},
		// Case 003, ~1.00 ns/op
		{
			b: 3,
		},
		// Case 004, ~1.00 ns/op
		{
			b: 4,
		},
		// Case 005, ~1.00 ns/op
		{
			b: 5,
		},
		// Case 006, ~1.00 ns/op
		{
			b: 6,
		},
		// Case 007, ~1.00 ns/op
		{
			b: 188,
		},
		// Case 008, ~1.00 ns/op
		{
			b: 255,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				_ = tc.b < 1 || tc.b > 4
			}
		})
	}
}
