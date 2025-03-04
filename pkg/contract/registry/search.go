package registry

import (
	"context"
	"errors"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/tracer"
)

func (c *Registry) Search(hsh common.Hash) (*types.Transaction, error) {
	var err error

	var txn *types.Transaction
	var pen bool
	{
		txn, pen, err = c.cli.TransactionByHash(context.Background(), hsh)
		if errors.Is(err, ethereum.NotFound) {
			return nil, tracer.Maskf(transactionNotFoundError, "%s", hsh)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We cannot accept the provided transaction as long as it is not successfully
	// included in a block. So as long as the provided transaction is still
	// pending, we reject it.
	if pen {
		return nil, tracer.Maskf(transactionStillPendingError, "%s", hsh)
	}

	var rec *types.Receipt
	{
		rec, err = c.cli.TransactionReceipt(context.Background(), hsh)
		if errors.Is(err, ethereum.NotFound) {
			return nil, tracer.Maskf(transactionNotFoundError, "%s", hsh)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// At this point the transaction was found and there was no error, which means
	// we have a receipt for a mined transaction. What we want to see now is the
	// status field set to 1, which is the specified success status code as per
	// EIP-658.
	//
	//     https://eips.ethereum.org/EIPS/eip-658
	//
	if rec.Status != types.ReceiptStatusSuccessful {
		return nil, tracer.Maskf(transactionNotSuccessfulError, "%s", hsh)
	}

	return txn, nil
}
