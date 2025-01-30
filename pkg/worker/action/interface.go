package action

import (
	"github.com/anubis-game/apiserver/pkg/worker/record"
	"github.com/google/uuid"
)

type Interface interface {
	Arg() []byte
	Rec() record.Interface
	Typ() string
	Uid() uuid.UUID
}
