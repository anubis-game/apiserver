package registry

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
		context.Background(),
		"level", "debug",
		"message", "submitted Registry.resolve transaction onchain",
		"signer", r.opt.From.Hex(),
		"kill", kil.String(),
		"winner", win.Hex(),
		"loser", los.Hex(),
		"transaction", txn.Hash().Hex(),
	)

	return txn, nil
}

func (r *Registry) writerOption() *bind.TransactOpts {
	return &bind.TransactOpts{
		From: r.opt.From,

		// Here we are trying to set some reasonable gas limits, specifically for
		// the EIP-1559 enabled minting transaction.
		//
		//     GasFeeCap is the max gas fee we are willing to pay
		//     GasTipCap is the max priority fee we are willing to pay
		//
		// Below is a testnet transaction providing some real world insight into
		// effective gas usage.
		//
		//     TODO
		//
		GasFeeCap: big.NewInt(5_000_000_000), // 5.00 gwei
		GasTipCap: big.NewInt(500_000_000),   // 0.50 gwei

		Signer: r.opt.Signer,
	}
}
