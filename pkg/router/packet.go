package router

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/google/uuid"
)

type Packet struct {
	Byt []byte
	Cli *client.Client
	Uid uuid.UUID
}
