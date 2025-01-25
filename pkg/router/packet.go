package router

import "github.com/anubis-game/apiserver/pkg/client"

type Packet struct {
	Byt []byte
	Cli *client.Client
}
