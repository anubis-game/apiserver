package connect

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/window"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
	"go.uber.org/ratelimit"
)

func (h *Handler) client(wal common.Address, con *websocket.Conn) error {
	var lim ratelimit.Limiter
	{
		lim = ratelimit.New(
			2,
			ratelimit.Per(25*time.Millisecond),
			ratelimit.WithSlack(0),
		)
	}

	var win *window.Window
	{
		win = window.New(window.Config{
			Bck: matrix.Bucket{
				h.crd.Random(), // x0
				h.crd.Random(), // y0
				h.crd.Random(), // x1
				h.crd.Random(), // y1
			},
			Pxl: matrix.Pixel{
				h.crd.Random(), // x2
				h.crd.Random(), // y2
			},
			Spc: matrix.Space{
				h.qdr.Random(), // quadrant
				h.ang.Random(), // angle
			},
		})
	}

	var cli *client.Client
	{
		cli = client.New(client.Config{
			Con: con,
			Ctx: h.ctx,
			Lim: lim,
			Wal: wal,
			Win: win,
		})
	}

	// With setting up the client connection, we also setup an expiration callback
	// in order to limit the connection lifetime of every client.

	{
		h.wxp.Ensure(wal, h.ttl, func() {
			defer close(cli.Expiry())
		})
	}

	// Below we manage the connection specific read loop. Any error occuring here
	// causes the connection to close.

	go func() {
		var err error

		{
			defer close(cli.Reader())
		}

		for {
			var byt []byte
			{
				_, byt, err = con.Read(h.ctx)
				if err != nil {
					return
				}
			}

			switch schema.Action(byt[0]) {
			case schema.Ping:
				err = h.ping(cli, byt)
			case schema.Auth:
				err = h.auth(cli, byt)
			case schema.Join:
				err = h.join(cli, byt)
			case schema.Move:
				err = h.move(cli, byt)
			case schema.Race:
				err = h.race(cli, byt)
			}

			if err != nil {
				h.log.Log(
					"level", "error",
					"message", err.Error(),
					"stack", tracer.Stack(err),
				)

				return
			}
		}
	}()

	// We block this client specific goroutine until either the client or the
	// server shuts down, each of which may happen for various reasons. Once
	// a signal got emitted to close this client connection, we cleanup any
	// internal references.

	select {
	case <-cli.Expiry():
	case <-cli.Reader():
	case <-cli.Writer():
	case <-h.don:
	}

	{
		h.rtr.Delete(cli)
		h.wxp.Delete(wal)
	}

	{
		err := con.CloseNow()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
