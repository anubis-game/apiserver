package engine

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
	"go.uber.org/ratelimit"
)

func Test_Engine_worker_read(t *testing.T) {
	var eng *Engine
	{
		eng = tesEng(250)
	}

	var uid byte
	{
		uid = 0x5
	}

	var fcn chan []byte
	var frc chan []byte
	{
		fcn = make(chan []byte, 1024)
		frc = make(chan []byte, 1024)
	}

	var cli *client.Client
	{
		cli = client.New(client.Config{
			Con: tesCon(frc),
			Don: make(<-chan struct{}),
			Fcn: fcn,
			Lim: ratelimit.New(1),
			Log: logger.Fake(),
			Tkx: tokenx.New[common.Address](),
		})
	}

	{
		go cli.Daemon()
	}

	{
		eng.ply.cli[uid] = fcn
	}

	//

	for i := range 10 {
		// Verify that the sequence byte is being incremented properly with every
		// update cycle.

		if eng.ply.buf[uid][0] != byte(i) {
			t.Fatalf("expected %#v got %#v", byte(i), eng.ply.buf[uid][0])
		}

		var wal common.Address
		{
			wal = tesWal(uid)
		}

		{
			eng.ply.buf[uid] = append(eng.ply.buf[uid], wal.Bytes()...)
		}

		var dur time.Duration
		{
			sta := time.Now()
			eng.send(sta)
			dur = time.Since(sta)
		}

		// We are effectively only making sure that some message is actually being
		// send to a non-blocking channel. This should happen in under 100
		// microseconds, even on limited CI providers.

		if dur > 100*time.Microsecond {
			t.Fatalf("expected %s got %s", "under 100 microseconds", dur)
		}

		// Ensure that the fanout receiver channel recorded the message that we sent
		// within the receiving websocket client of our test connection.

		msg := <-frc
		fmt.Printf("%#v\n", msg)
		if msg[0] != byte(i) {
			t.Fatalf("expected %#v got %#v", byte(i), msg[0])
		}
		if !bytes.Contains(msg, wal.Bytes()) {
			t.Fatalf("expected %#v got %#v", wal, msg)
		}
	}
}

func Benchmark_Engine_send(b *testing.B) {
	testCases := []struct {
		buf []byte
	}{
		// Case 000, ~4,000 ns/op, 2 allocs/op
		{
			buf: make([]byte, 16),
		},
		// Case 001, ~4,000 ns/op, 3 allocs/op
		{
			buf: make([]byte, 32),
		},
		// Case 002, ~4,000 ns/op, 3 allocs/op
		{
			buf: make([]byte, 64),
		},
		// Case 003, ~4,000 ns/op, 3 allocs/op
		{
			buf: make([]byte, 128),
		},
		// Case 004, ~4,000 ns/op, 3 allocs/op
		{
			buf: make([]byte, 256),
		},
		// Case 005, ~4,000 ns/op, 3 allocs/op
		{
			buf: make([]byte, 512),
		},
		// Case 006, ~4,300 ns/op, 4 allocs/op
		{
			buf: make([]byte, 1024),
		},
		// Case 007, ~7,300 ns/op, 7 allocs/op
		{
			buf: make([]byte, 2048),
		},
		// Case 008, ~11,100 ns/op, 9 allocs/op
		{
			buf: make([]byte, 4096),
		},
		// Case 009, ~16,300 ns/op, 11 allocs/op
		{
			buf: make([]byte, 8192),
		},
	}

	var eng *Engine
	{
		eng = tesEng(250)
	}

	var uid byte
	{
		uid = 0x5
	}

	var fcn chan []byte
	var frc chan []byte
	{
		fcn = make(chan []byte, 1024)
		frc = make(chan []byte, 1024)
	}

	var cli *client.Client
	{
		cli = client.New(client.Config{
			Con: tesCon(frc),
			Don: make(<-chan struct{}),
			Fcn: fcn,
			Lim: ratelimit.New(1),
			Log: logger.Fake(),
			Tkx: tokenx.New[common.Address](),
		})
	}

	{
		go cli.Daemon()
	}

	go func() {
		for {
			<-frc
		}
	}()

	{
		eng.ply.cli[uid] = fcn
	}

	tic := time.Now()

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				eng.ply.buf[uid] = append(eng.ply.buf[uid], tc.buf...)
				eng.send(tic)
			}
		})
	}

	//

	err := cli.Stream([]byte("ping"))
	if err != nil {
		b.Fatalf("expected %#v got %#v", nil, err)
	}
}

func tesCon(c chan<- []byte) *websocket.Conn {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			_, byt, err := con.Read(context.Background())
			if err != nil {
				panic(err)
			}

			{
				c <- byt
			}
		}
	}))

	con, _, err := websocket.Dial(context.Background(), srv.URL, nil)
	if err != nil {
		panic(err)
	}

	return con
}
