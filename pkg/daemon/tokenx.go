package daemon

import (
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/ethereum/go-ethereum/common"
)

func (d *Daemon) TokenX() *tokenx.TokenX[common.Address] {
	return d.tkx
}
