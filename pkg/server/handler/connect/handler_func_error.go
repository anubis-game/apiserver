package connect

import "github.com/xh3b4sd/tracer"

var protocolMethodInvalidError = &tracer.Error{
	Kind: "protocolMethodInvalidError",
	Desc: "The request expects the signature method to be one of [dual-handshake user-challenge]. The signature method was not found to be one of those options. Therefore the request failed.",
}

var walletAddressRegisteredError = &tracer.Error{
	Kind: "walletAddressRegisteredError",
	Desc: "The request expects the associated Wallet addresses to not be registered already. The associated Wallet addresses was found to be registered already. Therefore the request failed.",
}
