package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/tracer"
)

// Releasing a losing player is done by the respective Guardian in case a player
// got killed by the environment, e.g. bots or obstacles. In such a case there
// is no winning player, but the losing player has to be resolved anyway. Upon
// release, the losing player will receive their allocated balance back, minus
// the relevant Guardian and Protocol fees required to cover operational costs.
//
//     https://sepolia.arbiscan.io/tx/0x38f59f5d50e6e2cb72bb71971fc26bc60149711a41a64d8607cba923d1ef6185
//

func (r *Registry) Release(los common.Address) (*types.Transaction, error) {
	var err error

	var txn *types.Transaction
	{
		txn, err = r.bin.Release(r.writerOption(), los)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	r.log.Log(
		"level", "debug",
		"message", "submitted Registry.Release transaction onchain",
		"signer", r.opt.From.Hex(),
		"loser", los.Hex(),
		"transaction", txn.Hash().Hex(),
	)

	return txn, nil
}
