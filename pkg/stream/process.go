package stream

import (
	"bytes"
	"sync"

	"github.com/coder/websocket"
)

var (
	Ping = []byte("ping,")
)

func (s *Stream) process(add string, con *websocket.Conn) error {
	var onc sync.Once

	var don chan struct{}
	var rep chan struct{}
	{
		don = make(chan struct{})
		rep = make(chan struct{})
	}

	var cli Client
	{
		cli = Client{
			Close: func(add bool) {
				onc.Do(func() {
					// If close is called by Add, then we do not want to call Rem again in
					// the read loop below.
					if add {
						close(rep)
					}

					{
						close(don)
						con.CloseNow() //nolint:errcheck
					}
				})
			},
			Write: func(typ websocket.MessageType, byt []byte) {
				err := con.Write(s.ctx, typ, byt)
				if err != nil {
					go s.remove(add)
				}
			},
		}
	}

	{
		s.add(add, cli)
	}

	go func() {
		for {
			typ, byt, err := con.Read(s.ctx)
			if err != nil {
				// If this connection is closed from the outside, then we want to remove
				// the client from our internal state. If the same client replaces
				// itself, then we are reading from a closed connection and do not want
				// to remove the client again.
				select {
				case <-rep:
					// fall through
				default:
					go s.remove(add)
				}

				{
					return
				}
			} else {
				if bytes.HasPrefix(byt, Ping) {
					cli.Write(typ, byt)
				} else {
					go s.write(add, typ, byt)
				}
			}
		}
	}()

	select {
	case <-don:
	case <-s.don:
		s.remove(add)
	}

	return nil
}
