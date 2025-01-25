package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Registry) Resolve(kil objectid.ID, win common.Address, los common.Address) (*types.Transaction, error) {
	var err error

	var txn *types.Transaction
	{
		txn, err = r.bin.Resolve(r.writerOption(), kil.Big(), win, los)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	r.log.Log(
		"level", "debug",
		"message", "submitted Registry.Resolve transaction onchain",
		"signer", r.opt.From.Hex(),
		"kill", kil.String(),
		"winner", win.Hex(),
		"loser", los.Hex(),
		"transaction", txn.Hash().Hex(),
	)

	return txn, nil
}
