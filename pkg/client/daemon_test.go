package client

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/coder/websocket"
)

func Test_Client_Daemon_overload(t *testing.T) {
	var cli *Client
	{
		cli = New(Config{
			Con: tesCon("localhost:30004", "overload", tesHan),
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
		cli.buf <- buf               // the first message processes immediatelly, the 5 following messages accumulate
		time.Sleep(time.Millisecond) // give the test time to propagate
	}

	select {
	case <-cli.tic:
		// ticker channel closed as expected
	default:
		t.Fatalf("expected ticker channel to close")
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
