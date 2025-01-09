package stream

import (
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/xh3b4sd/tracer"
)

type Client struct {
	Close func()
	Write func([]byte)
}

func (s *Stream) client(add string, con *websocket.Conn) error {
	var clo chan struct{}
	var exp chan struct{}
	var rea chan struct{}
	var wri chan struct{}
	{
		clo = make(chan struct{})
		exp = make(chan struct{})
		rea = make(chan struct{})
		wri = make(chan struct{})
	}

	var cli Client
	{
		cli = Client{
			Close: func() {
				close(clo)
			},
			Write: func(byt []byte) {
				err := con.Write(s.ctx, websocket.MessageText, byt)
				if err != nil {
					close(wri)
				}
			},
		}
	}

	// The first functional step in our connection management is to associate the
	// user's client connection with their verified Wallet address. Creating this
	// association adds the user to the system and provides them with realtime
	// data primitives over the given client connection. With setting up the
	// client connection, we also setup an expiration callback in order to limit
	// the connection lifetime of every client.

	{
		s.cli.Update(add, cli)
	}

	{
		s.exp.Ensure(add, func() {
			defer close(exp)
		})
	}

	// Below we manage the connection specific read loop. Any error occuring here
	// causes the connection to close.

	go func() {
		{
			defer close(rea)
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
				{
					err = s.ping(con)
				}
			case schema.Auth:
				{
					err = s.auth(con, add)
				}
			case schema.Cast:
				{
					err = s.cast(byt, add)
				}
			case schema.Move:
				// TODO
			case schema.Kill:
				// TODO
			}

			if err != nil {
				return
			}
		}
	}()

	// We block this client specific goroutine until either the client or the
	// server shuts down, each of which may happen for various reasons. Once
	// a signal got emitted to close this client connection, we cleanup any
	// internal references.

	select {
	case <-clo:
	case <-exp:
	case <-rea:
	case <-wri:
	case <-s.don:
	}

	{
		s.cli.Delete(add)
		s.exp.Delete(add)
	}

	{
		err := con.CloseNow()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
