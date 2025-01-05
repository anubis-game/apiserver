package stream

import (
	"bytes"

	"github.com/coder/websocket"
)

var (
	Ping = []byte("ping,")
)

func (s *Stream) process(add string, con *websocket.Conn) error {
	var rea chan struct{}
	var wri chan struct{}
	{
		rea = make(chan struct{})
		wri = make(chan struct{})
	}

	var cli Client
	{
		cli = Client{
			Close: func() {
				con.CloseNow() //nolint:errcheck
			},
			Write: func(typ websocket.MessageType, byt []byte) {
				err := con.Write(s.ctx, typ, byt)
				if err != nil {
					close(wri)
				}
			},
		}
	}

	// The first functional step in our websocket management process is to
	// associate the user's client connection with their verified Wallet address.
	// Creating this association adds the user to the system and provides them
	// with realtime data primitives.
	{
		s.create(add, cli)
	}

	// Here we manage the connection specific read loop.
	go func() {
		for {
			typ, byt, err := con.Read(s.ctx)
			if err != nil {
				close(rea)
				return
			}

			{
				if bytes.HasPrefix(byt, Ping) {
					cli.Write(typ, byt)
				} else {
					go s.write(add, typ, byt)
				}
			}
		}
	}()

	// We block this websocket connection specific goroutine until either the
	// client or the server shuts down.
	select {
	case <-rea:
	case <-wri:
	case <-s.don:
	}

	{
		s.delete(add)
	}

	return nil
}
