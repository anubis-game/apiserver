package connect

import "github.com/xh3b4sd/tracer"

var handshakeHeaderInvalidError = &tracer.Error{
	Kind: "handshakeHeaderInvalidError",
	Desc: "The request expects the Sec-Websocket-Protocol header to contain exactly 3 elements. The Sec-Websocket-Protocol header was not found to comply with this requirement. Therefore the request failed.",
}

var handshakeValidationFailedError = &tracer.Error{
	Kind: "handshakeValidationFailedError",
	Desc: "The request expects the dual-handshake to succeed for any player to participate in the game. The dual-handshake was not found to comply with the specified requirements. Therefore the request failed.",
}

var signatureHashLengthError = &tracer.Error{
	Kind: "signatureHashLengthError",
	Desc: "The request expects the signature hash to have 132 characters. The signature hash was not found to have 132 characters. Therefore the request failed.",
}

var signatureTimeInvalidError = &tracer.Error{
	Kind: "signatureTimeInvalidError",
	Desc: "The request expects the signature time to be within the range of +-60 seconds relative to the current time. The signature time was not found to be within that threshold. Therefore the request failed.",
}

var signerAddressMatchError = &tracer.Error{
	Kind: "signerAddressMatchError",
	Desc: "The request expects the associated Signer addresses to match one another. The associated Signer addresses were not found to match one another. Therefore the request failed.",
}

var transactionHashLengthError = &tracer.Error{
	Kind: "transactionHashLengthError",
	Desc: "The request expects the transaction hash to have 66 characters. The transaction hash was not found to have 66 characters. Therefore the request failed.",
}
