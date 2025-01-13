package resolve

import (
	"context"
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	TTL = 10 * time.Second
)

type Config struct {
	Log logger.Interface
	Reg *registry.Registry
}

type Resolve struct {
	log logger.Interface
	reg *registry.Registry
}

func New(c Config) *Resolve {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Reg == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Reg must not be empty", c)))
	}

	return &Resolve{
		log: c.Log,
		reg: c.Reg,
	}
}

func (r *Resolve) Ensure(pac Packet) (Packet, common.Address, time.Duration) {
	var err error

	var ttl time.Duration
	{
		pac, ttl, err = r.ensure(pac)
		if err != nil {
			r.log.Log(
				context.Background(),
				"level", "error",
				"message", err.Error(),
				"stack", tracer.Stack(err),
			)
		}
	}

	return pac, pac.Loser, ttl
}

// TODO abstract the reconciliation away for transactions
func (r *Resolve) ensure(pac Packet) (Packet, time.Duration, error) {
	// var err error

	// if transaction.Empty(pac.Transaction) {
	// 	var txn *types.Transaction
	// 	{
	// 		txn, err = r.reg.Resolve(pac.Kill, pac.Winner, pac.Loser)
	// 		if err != nil {
	// 			return pac, TTL, tracer.Mask(err)
	// 		}
	// 	}

	// 	{
	// 		pac.Transaction = txn.Hash()
	// 	}

	// 	return pac, TTL, nil
	// }

	// {
	// 	_, err = r.reg.Search(pac.Transaction)
	// 	if registry.IsTransactionNotFoundError(err) {
	// 		return pac, TTL, nil
	// 	} else if registry.IsTransactionStillPending(err) {
	// 		return pac, TTL, nil
	// 	} else if registry.IsTransactionNotSuccessfulError(err) {
	// 		{
	// 			pac.Transaction = common.Hash{}
	// 		}

	// 		return pac, TTL, tracer.Mask(err)
	// 	} else if err != nil {
	// 		return pac, 0, tracer.Mask(err)
	// 	}
	// }

	fmt.Printf("RESOLVE %#v\n", pac)
	return pac, 0, nil
}
