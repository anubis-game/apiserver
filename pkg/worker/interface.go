package worker

import (
	"github.com/anubis-game/apiserver/pkg/worker/action"
	"github.com/ethereum/go-ethereum/common"
)

type Daemon interface {
	Daemon()
}

type Ensure interface {
	Ensure(action.Interface)
}

type Signer interface {
	Sign([]byte) (common.Hash, error)
	Type() string
}
