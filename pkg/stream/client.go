package stream

import "github.com/coder/websocket"

type Client struct {
	Clo func(bool)
	Wri func(websocket.MessageType, []byte)
}
