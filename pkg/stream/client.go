package stream

import "github.com/coder/websocket"

type Client struct {
	Close func(bool)
	Write func(websocket.MessageType, []byte)
}
