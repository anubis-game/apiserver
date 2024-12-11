package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (r *Registry) Sender(txn *types.Transaction) (common.Address, error) {
	var err error

	var add common.Address
	{
		add, err = r.sig.Sender(txn)
		if err != nil {
			return common.Address{}, nil
		}
	}

	return add, nil
}
