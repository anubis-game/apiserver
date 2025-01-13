package registry

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

func IsTransactionNotFoundError(err error) bool {
	return errors.Is(err, transactionNotFoundError)
}

var transactionNotFoundError = &tracer.Error{
	Kind: "transactionNotFoundError",
	Desc: "The request expects the provided transaction to exist. The provided transaction was not found to exist. Therefore the request failed.",
}

func IsTransactionNotSuccessfulError(err error) bool {
	return errors.Is(err, transactionNotSuccessfulError)
}

var transactionNotSuccessfulError = &tracer.Error{
	Kind: "transactionNotSuccessfulError",
	Desc: "The request expects the provided transaction to be successfully included in a block. The provided transaction was not found to be successfully included in a block. Therefore the request failed.",
}

func IsTransactionStillPending(err error) bool {
	return errors.Is(err, transactionStillPendingError)
}

var transactionStillPendingError = &tracer.Error{
	Kind: "transactionStillPendingError",
	Desc: "The request expects the provided transaction to not be pending anymore. The provided transaction was found to still be pending. Therefore the request failed.",
}
