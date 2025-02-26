package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
	"go.uber.org/ratelimit"
)

func Test_Client_Daemon_cleanup(t *testing.T) {
	var err error
	var wai sync.WaitGroup

	// Remember the amount of goroutines that we start out with.

	var num int
	{
		num = runtime.NumGoroutine()
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		{
			wai.Add(1)
		}

		con, err := websocket.Accept(w, r, nil)
		if err != nil {
			panic(err)
		}

		var rtr *router.Router
		{
			rtr = router.New(router.Config{
				Cap: 50,
			})
		}

		var cli *Client
		{
			cli = New(Config{
				Con: con,
				Don: make(<-chan struct{}),
				Fcn: make(chan []byte, 1024),
				Lim: ratelimit.New(1),
				Log: logger.Fake(),
				Rtr: rtr.Client(),
				Tkx: tokenx.New[common.Address](),
			})
		}

		{
			go cli.Daemon()
		}

		{
			time.Sleep(time.Millisecond)
		}

		{
			cli.Delete()
		}

		{
			wai.Done()
		}
	}))

	var con *websocket.Conn
	{
		con, _, err = websocket.Dial(context.Background(), srv.URL, nil)
		if err != nil {
			panic(err)
		}
	}

	// Ensure some more goroutines got actually started.

	{
		ngo := runtime.NumGoroutine()
		if num >= ngo {
			t.Fatalf("expected %#v got %#v", "more goroutines", ngo)
		}
	}

	// Ensure the handler code completed successfully.

	{
		wai.Wait()
		con.Close(websocket.StatusNormalClosure, "")
		srv.Close()
	}

	// Ensure the amount of goroutines equals the amount that we started with.

	{
		ngo := runtime.NumGoroutine()
		if num != ngo {
			t.Fatalf("expected %#v got %#v", num, ngo)
		}
	}
}

func Test_Client_Daemon_overload(t *testing.T) {
	var cli *Client
	{
		cli = New(Config{
			Con: tesCon(),
			Don: make(<-chan struct{}),
			Fcn: make(chan []byte, 1024),
			Lim: ratelimit.New(1),
			Log: logger.Fake(),
			Tkx: tokenx.New[common.Address](),
		})
	}

	{
		cli.cap = 5                           // allow for 5 concurrent messages to accumulate
		cli.tiC = time.Tick(time.Microsecond) // check for congestion very often
	}

	{
		go cli.Daemon()
	}

	var buf []byte
	{
		buf = make([]byte, 50*1024*1024)
	}

	for range 6 {
		cli.fcn <- buf               // the first message processes immediatelly, the 5 following messages accumulate
		time.Sleep(time.Millisecond) // give the test time to propagate
	}

	select {
	case <-cli.tic:
		// ticker channel closed as expected
	default:
		t.Fatalf("expected ticker channel to close")
	}
}

func tesCon() *websocket.Conn {
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
	}))

	con, _, err := websocket.Dial(context.Background(), srv.URL, nil)
	if err != nil {
		panic(err)
	}

	return con
}
