package stream

import "github.com/coder/websocket"

type Client struct {
	Close func()
	Write func(websocket.MessageType, []byte)
}
