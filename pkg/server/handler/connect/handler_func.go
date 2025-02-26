package connect

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
	"go.uber.org/ratelimit"
)

const (
	// Per is the duration bucket for the client specific rate limiters that guard
	// the game engine fanout procedure from external overloading.
	Per = vector.Frm * time.Millisecond
)

func (h *Handler) HandlerFunc(w http.ResponseWriter, r *http.Request) error {
	var err error

	// We use a semaphore pattern in order to ensure a concurrent connection
	// limit. If there is an available slot for this request handler, then this
	// free slot may be occupied by the calling client, and we process the
	// request. If no slot is available, then we return an error response.
	select {
	case h.sem <- struct{}{}:
		{
			err = h.handlerFunc(w, r)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			<-h.sem
		}
	default:
		http.Error(
			w,
			fmt.Sprintf("%s: %d", http.StatusText(http.StatusTooManyRequests), cap(h.sem)),
			http.StatusTooManyRequests,
		)
	}

	return nil
}

func (h *Handler) handlerFunc(w http.ResponseWriter, r *http.Request) error {
	var err error

	// The reuqest headers contain the desired protocol method as well as various
	// credentials used to authorize stream engine clients. See the links below
	// for more information about what we are doing here and why.
	//
	//     https://stackoverflow.com/questions/4361173/http-headers-in-websockets-client-api/77060459#77060459
	//     https://github.com/kubernetes/kubernetes/commit/714f97d7baf4975ad3aa47735a868a81a984d1f0
	//
	var hea []string
	{
		hea = reqHea(r.Header["Sec-Websocket-Protocol"])
	}

	var wal common.Address
	switch schema.Header(hea[0]) {
	case schema.DualHandshake:
		{
			wal, err = h.verify(hea)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	case schema.UserChallenge:
		{
			wal, err = h.search(hea)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	default:
		return tracer.Maskf(protocolMethodInvalidError, "%s", hea[0])
	}

	// TODO:infra how can we prevent requests for wallet addresses that we are
	// already serving?

	var con *websocket.Conn
	{
		con, err = websocket.Accept(w, r, h.opt)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Create a compact byte ID and associate it with the given wallet address.
	// If clients ever reconnect using their session tokens, we can allow them to
	// continue playing their game after any intermittend interruption.

	var cli *client.Client
	{
		cli = client.New(client.Config{
			Con: con,
			Don: h.don,
			Fcn: make(chan []byte, 1024),
			Lim: newLim(),
			Log: h.log,
			Rtr: h.rtr,
			Uid: h.uni.Ensure(wal),
			Wal: wal,
		})
	}

	// Below we manage the connection specific reader loop. Any error occuring
	// here causes the connection to close, regardless where the underlying error
	// originated from. In some cases, reading from the websocket connection may
	// fail. In other cases some internal logic may cause the reader loop to
	// produce an error, either due to invalid reconciliation results, or
	// websocket writes.

	{
		go cli.Daemon()
	}

	// We block this client specific goroutine until either the client or the
	// server shuts down, each of which may happen for various reasons. Once a
	// signal got emitted to close this client connection, we remove the client
	// from all internal references, but keep the player in the game. The reason
	// for not removing players from games during disconnects is that connections
	// might be dropped intermittently. That means the client may very well come
	// back quickly using its auth token and continue playing their game.

	select {
	case <-h.don:
	case <-cli.Expiry():
	case <-cli.Reader():
	case <-cli.Ticker():
	case <-cli.Writer():
	}

	{
		cli.Delete()
	}

	return nil
}

func newLim() ratelimit.Limiter {
	return ratelimit.New(
		5,                      // 2 race, 2 turn, 1 buffer
		ratelimit.Per(Per),     // per standard frame
		ratelimit.WithSlack(0), // without re-using unused capacity
	)
}

func reqHea(lis []string) []string {
	var spl []string

	for _, x := range lis {
		for _, y := range strings.Split(x, ",") {
			spl = append(spl, strings.TrimSpace(y))
		}
	}

	return spl
}
