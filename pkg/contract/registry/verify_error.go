package registry

import (
	"github.com/xh3b4sd/tracer"
)

var guardianAddressMatchError = &tracer.Error{
	Kind: "guardianAddressMatchError",
	Desc: "The request expects the associated Guardian addresses to match one another. The associated Guardian addresses were not found to match one another. Therefore the request failed.",
}

var signatureHashInvalidError = &tracer.Error{
	Kind: "signatureHashInvalidError",
	Desc: "The request expects the signature hash to be valid. The signature hash was not found to be valid. Therefore the request failed.",
}

var signatureTimeExpiredError = &tracer.Error{
	Kind: "signatureTimeExpiredError",
	Desc: "The request expects the signature time to not be older than 60 seconds. The signature time was found to be older than that threshold. Therefore the request failed.",
}

var signatureTimeFutureError = &tracer.Error{
	Kind: "signatureTimeFutureError",
	Desc: "The request expects the signature time to not be in the future. The signature time was found to be in the future. Therefore the request failed.",
}

var signerAddressMatchError = &tracer.Error{
	Kind: "signerAddressMatchError",
	Desc: "The request expects the associated Signer addresses to match one another. The associated Signer addresses were not found to match one another. Therefore the request failed.",
}
