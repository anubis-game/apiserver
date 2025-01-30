package release

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/worker/record"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Action struct {
	Act uuid.UUID
	Kil uuid.UUID
	Los common.Address
}

func (a Action) Arg() []byte {
	var buf [36]byte

	copy(buf[0:16], a.Kil[:])
	copy(buf[16:36], a.Los[:])

	return buf[:]
}

func (a Action) Rec() record.Interface {
	return record.NewSlicer(record.SlicerConfig{})
}

func (a Action) Typ() string {
	return Typ
}

func (a Action) Uid() uuid.UUID {
	return a.Act
}

func fromBytes(byt []byte) Action {
	if len(byt) != 36 {
		panic(fmt.Sprintf("expected 36 packet bytes, got %d", len(byt)))
	}

	var act Action
	{
		copy(act.Kil[:], byt[0:16])
		copy(act.Los[:], byt[16:36])
	}

	return act
}
