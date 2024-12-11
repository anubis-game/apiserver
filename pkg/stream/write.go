package stream

import "github.com/coder/websocket"

func (s *Stream) write(_ string, typ websocket.MessageType, byt []byte) {
	s.mut.RLock()
	for _, x := range s.cli {
		x.Write(typ, byt)
	}
	s.mut.RUnlock()
}
