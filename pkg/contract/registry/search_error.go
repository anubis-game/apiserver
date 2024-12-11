package registry

import "github.com/xh3b4sd/tracer"

var transactionNotFoundError = &tracer.Error{
	Kind: "TransactionNotFoundError",
	Desc: "The request expects the provided transaction to exist. The provided transaction was not found to exist. Therefore the request failed.",
}

var transactionNotSuccessfulError = &tracer.Error{
	Kind: "TransactionNotSuccessfulError",
	Desc: "The request expects the provided transaction to be successfully included in a block. The provided transaction was not found to be successfully included in a block. Therefore the request failed.",
}
