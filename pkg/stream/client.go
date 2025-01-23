package stream

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/window"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

var (
	pong = []byte{byte(schema.Pong)}
)

func (s *Stream) client(wal common.Address, con *websocket.Conn) error {
	var win *window.Window
	{
		win = window.New(window.Config{
			Bck: matrix.Bucket{
				s.crd.Random(), // x0
				s.crd.Random(), // y0
				s.crd.Random(), // x1
				s.crd.Random(), // y1
			},
			Pxl: matrix.Pixel{
				s.crd.Random(), // x2
				s.crd.Random(), // y2
			},
			Spc: matrix.Space{
				s.qdr.Random(), // quadrant
				s.ang.Random(), // angle
			},
		})
	}

	var cli *client.Client
	{
		cli = client.New(client.Config{
			Con: con,
			Ctx: s.ctx,
			Wal: wal,
			Win: win,
		})
	}

	// With setting up the client connection, we also setup an expiration callback
	// in order to limit the connection lifetime of every client.

	{
		s.wxp.Ensure(wal, s.ttl, func() {
			defer close(cli.Expiry())
		})
	}

	// Below we manage the connection specific read loop. Any error occuring here
	// causes the connection to close.

	go func() {
		{
			defer close(cli.Reader())
		}

		for {
			var err error

			var byt []byte
			{
				_, byt, err = con.Read(s.ctx)
				if err != nil {
					return
				}
			}

			// TODO prevent DDOS attacks and rate limit stream input somehow so that
			// the 25 millisecond schedule cannot be overloaded artificially.

			switch schema.Action(byt[0]) {
			case schema.Ping:
				cli.Stream(pong)
			case schema.Auth:
				err = s.auth(cli)
			case schema.Join:
				s.rtr.Create <- Packet{byt, cli}
			case schema.Move:
				// TODO move must also adapt the client window coordinates
			case schema.Race:
				// TODO
			}

			if err != nil {
				s.log.Log(
					s.ctx,
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
	case <-s.rtr.Closer:
	}

	{
		s.rtr.Delete <- Packet{nil, cli}
		s.wxp.Delete(wal)
	}

	{
		err := con.CloseNow()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
