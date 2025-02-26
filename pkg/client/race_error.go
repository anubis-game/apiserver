package client

import (
	"github.com/xh3b4sd/tracer"
)

var raceBytesInvalidError = &tracer.Error{
	Kind: "raceBytesInvalidError",
	Desc: "The request expects exactly 1 input byte, one byte for the action. The input bytes were not found to comply with this requirement. Therefore the request failed.",
}
