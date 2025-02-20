package engine

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/coder/websocket"
)

func Test_Engine_worker_Read(t *testing.T) {
	var eng *Engine
	{
		eng = &Engine{
			sem: make(chan struct{}, 1),
		}
	}

	var ply *player.Player
	{
		ply = &player.Player{
			Cli: client.New(client.Config{
				Con: tesCon("localhost:30001", "read", tesHan),
			}),
		}
	}

	var buf []byte
	{
		buf = make([]byte, 32)
	}

	var tim *time.Timer
	{
		tim = time.NewTimer(4 * time.Millisecond)
	}

	//

	var dur time.Duration
	{
		sta := time.Now()
		eng.worker(ply, buf, tim)
		dur = time.Since(sta)
	}

	// A couple of bytes should be written in under 1 millisecond, so the *time.Timer
	// should not be invoked.

	if dur > 1*time.Millisecond {
		t.Fatalf("expected %s got %s", "under 5ms", dur)
	}

	//

	err := ply.Cli.Stream([]byte("ping"))
	if err != nil {
		t.Fatalf("expected %#v got %#v", nil, err)
	}
}

func Test_Engine_worker_Timeout(t *testing.T) {
	var eng *Engine
	{
		eng = &Engine{
			sem: make(chan struct{}, 1),
		}
	}

	var ply *player.Player
	{
		ply = &player.Player{
			Cli: client.New(client.Config{
				Con: tesCon("localhost:30002", "timeout", tesHan),
			}),
		}
	}

	var buf []byte
	{
		buf = make([]byte, 50*1024*1024)
	}

	var tim *time.Timer
	{
		tim = time.NewTimer(10 * time.Millisecond)
	}

	//

	var dur time.Duration
	{
		sta := time.Now()
		eng.worker(ply, buf, tim)
		dur = time.Since(sta)
	}

	// 50 MB are not written within 10 milliseconds, so the test should timeout
	// using the configured *time.Timer.

	if dur < 9*time.Millisecond {
		t.Fatalf("expected %s got %s", "over 9ms", dur)
	}
	if dur > 11*time.Millisecond {
		t.Fatalf("expected %s got %s", "under 11ms", dur)
	}

	//

	err := ply.Cli.Stream([]byte("ping"))
	if err != nil {
		t.Fatalf("expected %#v got %#v", nil, err)
	}
}

func Benchmark_Engine_worker(b *testing.B) {
	testCases := []struct {
		buf []byte
	}{
		// Case 000, ~6,800 ns/op for 2 bytes
		{
			buf: make([]byte, 2),
		},
		// Case 001, ~6,800 ns/op for 32 bytes
		{
			buf: make([]byte, 32),
		},
		// Case 002, ~6,800 ns/op for 64 bytes
		{
			buf: make([]byte, 64),
		},
		// Case 003, ~6,800 ns/op for 128 bytes
		{
			buf: make([]byte, 128),
		},
		// Case 004, ~6,800 ns/op for 256 bytes
		{
			buf: make([]byte, 256),
		},
		// Case 005, ~6,800 ns/op for 512 bytes
		{
			buf: make([]byte, 512),
		},
		// Case 006, ~7,000 ns/op for 1024 bytes
		{
			buf: make([]byte, 1024),
		},
		// Case 007, ~7,250 ns/op for 2048 bytes
		{
			buf: make([]byte, 2048),
		},
		// Case 008, ~9,250 ns/op for 4096 bytes
		{
			buf: make([]byte, 4096),
		},
		// Case 009, ~17,250 ns/op for 8192 bytes
		{
			buf: make([]byte, 8192),
		},
	}

	var eng *Engine
	{
		eng = &Engine{
			sem: make(chan struct{}, 1),
		}
	}

	var ply *player.Player
	{
		ply = &player.Player{
			Cli: client.New(client.Config{
				Con: tesCon("localhost:30002", "bench", tesHan),
			}),
		}
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tim := time.NewTimer(4 * time.Millisecond)
				eng.worker(ply, tc.buf, tim)
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
