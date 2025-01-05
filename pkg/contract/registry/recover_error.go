package registry

import (
	"github.com/xh3b4sd/tracer"
)

var signatureInvalidError = &tracer.Error{
	Kind: "signatureInvalidError",
	Desc: "The request expects the signature hash to be valid. The signature hash was not found to be valid. Therefore the request failed.",
}
