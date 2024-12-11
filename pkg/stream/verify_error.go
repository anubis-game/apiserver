package stream

import "github.com/xh3b4sd/tracer"

var handshakeHeaderInvalidError = &tracer.Error{
	Kind: "handshakeHeaderInvalidError",
	Desc: "The request expects the Sec-Websocket-Protocol header to contain exactly 3 elements. The Sec-Websocket-Protocol header was not found to comply with this requirement. Therefore the request failed.",
}

var handshakeMethodInvalidError = &tracer.Error{
	Kind: "handshakeMethodInvalidError",
	Desc: "The request expects the signature method to be one of [dual-handshake]. The signature method was not found to be one of those options. Therefore the request failed.",
}

var signatureHashLengthError = &tracer.Error{
	Kind: "signatureHashLengthError",
	Desc: "The request expects the signature hash to have 132 characters. The signature hash was not found to have 132 characters. Therefore the request failed.",
}

var transactionHashLengthError = &tracer.Error{
	Kind: "transactionHashLengthError",
	Desc: "The request expects the transaction hash to have 66 characters. The transaction hash was not found to have 66 characters. Therefore the request failed.",
}
