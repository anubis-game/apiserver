package release

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/tracer"
)

const (
	Typ = "release"
)

type Config struct {
	Reg *registry.Registry
}

type Release struct {
	reg *registry.Registry
}

func New(c Config) *Release {
	if c.Reg == nil {
		tracer.Panic(fmt.Errorf("%T.Reg must not be empty", c))
	}

	return &Release{
		reg: c.Reg,
	}
}

func (r *Release) Sign(byt []byte) (common.Hash, error) {
	var err error

	var act Action
	{
		act = fromBytes(byt)
	}

	var txn *types.Transaction
	{
		txn, err = r.reg.Release(act.Los)
		if err != nil {
			return common.Hash{}, tracer.Mask(err)
		}
	}

	return txn.Hash(), nil
}

func (r *Release) Type() string {
	return Typ
}
