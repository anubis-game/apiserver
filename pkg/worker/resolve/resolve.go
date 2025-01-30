package resolve

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/tracer"
)

const (
	Typ = "resolve"
)

type Config struct {
	Reg *registry.Registry
}

type Resolve struct {
	reg *registry.Registry
}

func New(c Config) *Resolve {
	if c.Reg == nil {
		tracer.Panic(fmt.Errorf("%T.Reg must not be empty", c))
	}

	return &Resolve{
		reg: c.Reg,
	}
}

func (r *Resolve) Sign(byt []byte) (common.Hash, error) {
	var err error

	var act Action
	{
		act = fromBytes(byt)
	}

	var txn *types.Transaction
	{
		txn, err = r.reg.Resolve(act.Kil, act.Win, act.Los)
		if err != nil {
			return common.Hash{}, tracer.Mask(err)
		}
	}

	return txn.Hash(), nil
}

func (r *Resolve) Type() string {
	return Typ
}
