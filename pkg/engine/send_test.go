package engine

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
)

func Test_Engine_worker_read(t *testing.T) {
	var eng *Engine
	{
		eng = &Engine{
			buf: xsync.NewMapOf[byte, []byte](),
			mem: &memory{
				ply: xsync.NewMapOf[byte, *player.Player](),
			},
			uni: unique.New[common.Address, byte](),
		}
	}

	var uid byte
	{
		uid = eng.uni.Ensure(common.Address{})
	}

	var ply *player.Player
	{
		ply = &player.Player{
			Cli: client.New(client.Config{
				Con: tesCon("localhost:30001", "read", tesHan),
			}),
		}
	}

	{
		go ply.Cli.Daemon()
	}

	var buf []byte
	{
		buf = make([]byte, 32)
	}

	{
		eng.mem.ply.Store(uid, ply)
	}

	//

	for range 10 {
		{
			eng.buf.Store(uid, buf)
		}

		var dur time.Duration
		{
			sta := time.Now()
			eng.send(sta)
			dur = time.Since(sta)
		}

		// We are effectively only making sure that some message is actually being
		// send to a non-blocking channel. This should happen under 100
		// microseconds.

		if dur > 100*time.Microsecond {
			t.Fatalf("expected %s got %s", "under 100 microseconds", dur)
		}

		//

		err := ply.Cli.Stream([]byte("ping"))
		if err != nil {
			t.Fatalf("expected %#v got %#v", nil, err)
		}
	}
}

func Benchmark_Engine_send(b *testing.B) {
	testCases := []struct {
		buf []byte
	}{
		// Case 000, ~3,900 ns/op, 1 allocs/op
		{
			buf: make([]byte, 2),
		},
		// Case 001, ~3,900 ns/op, 2 allocs/op
		{
			buf: make([]byte, 32),
		},
		// Case 002, ~3,900 ns/op, 2 allocs/op
		{
			buf: make([]byte, 64),
		},
		// Case 003, ~3,900 ns/op, 2 allocs/op
		{
			buf: make([]byte, 128),
		},
		// Case 004, ~3,900 ns/op, 2 allocs/op
		{
			buf: make([]byte, 256),
		},
		// Case 005, ~4,100 ns/op, 2 allocs/op
		{
			buf: make([]byte, 512),
		},
		// Case 006, ~4,200 ns/op, 3 allocs/op
		{
			buf: make([]byte, 1024),
		},
		// Case 007, ~6,300 ns/op, 6 allocs/op
		{
			buf: make([]byte, 2048),
		},
		// Case 008, ~10,000 ns/op, 8 allocs/op
		{
			buf: make([]byte, 4096),
		},
		// Case 009, ~17,800 ns/op, 10 allocs/op
		{
			buf: make([]byte, 8192),
		},
	}

	var eng *Engine
	{
		eng = &Engine{
			buf: xsync.NewMapOf[byte, []byte](),
			mem: &memory{
				ply: xsync.NewMapOf[byte, *player.Player](),
			},
			uni: unique.New[common.Address, byte](),
		}
	}

	var uid byte
	{
		uid = eng.uni.Ensure(common.Address{})
	}

	var ply *player.Player
	{
		ply = &player.Player{
			Cli: client.New(client.Config{
				Con: tesCon("localhost:30003", "bench", tesHan),
			}),
		}
	}

	{
		go ply.Cli.Daemon()
	}

	{
		eng.mem.ply.Store(uid, ply)
	}

	tic := time.Now()

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				eng.buf.Store(uid, tc.buf)
				eng.send(tic)
			}
		})
	}
}

func tesHan(w http.ResponseWriter, r *http.Request) {
	con, err := websocket.Accept(w, r, nil)
	if err != nil {
		return
	}

	// We disable the read limit to work around some default settings causing
	// runtime panics.
	//
	//     panic serving 127.0.0.1:61559: failed to read: read limited at 32769 bytes
	//

	{
		con.SetReadLimit(-1)
	}

	for {
		_, _, err := con.Read(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func tesCon(add string, pat string, han http.HandlerFunc) *websocket.Conn {
	go func() {
		{
			http.HandleFunc("/"+pat, han)
		}

		err := http.ListenAndServe(add, nil)
		if err != nil {
			panic(err)
		}
	}()

	{
		time.Sleep(100 * time.Millisecond)
	}

	con, _, err := websocket.Dial(context.Background(), "ws://"+add+"/"+pat, nil)
	if err != nil {
		panic(err)
	}

	return con
}
