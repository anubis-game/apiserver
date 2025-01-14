package stream

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

func (s *Stream) client(wal common.Address, con *websocket.Conn) error {
	var cli *client.Client
	{
		cli = client.New(client.Config{
			Con: con,
			Ctx: s.ctx,
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

			switch schema.Action(byt[0]) {
			case schema.Ping:
				err = s.ping(wal, cli, byt)
			case schema.Auth:
				err = s.auth(wal, cli, byt)
			case schema.Join:
				err = s.join(wal, cli, byt)
			case schema.Cast:
				err = s.cast(wal, cli, byt) // TODO we should not allow anyone to just cast anything to everyone
			case schema.Move:
				// TODO
			case schema.Kill:
				err = s.kill(wal, cli, byt)
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
	case <-s.don:
	}

	{
		s.cli.Delete(wal)
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
