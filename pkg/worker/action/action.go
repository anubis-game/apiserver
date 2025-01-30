package action

import (
	"github.com/anubis-game/apiserver/pkg/worker/record"
	"github.com/google/uuid"
)

type Action struct {
	arg []byte
	rec record.Interface
	typ string
	uid uuid.UUID
}

func New(c Interface) *Action {
	return &Action{
		arg: c.Arg(),
		rec: c.Rec(),
		typ: c.Typ(),
		uid: c.Uid(),
	}
}

func (p *Action) Arg() []byte {
	return p.arg
}

func (p *Action) Rec() record.Interface {
	return p.rec
}

func (p *Action) Typ() string {
	return p.typ
}

func (p *Action) Uid() uuid.UUID {
	return p.uid
}
