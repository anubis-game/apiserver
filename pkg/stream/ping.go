package stream

import (
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/xh3b4sd/tracer"
)

func (s *Stream) ping(con *websocket.Conn) error {
	var err error

	{
		err = con.Write(s.ctx, websocket.MessageText, []byte{byte(schema.Pong)})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
