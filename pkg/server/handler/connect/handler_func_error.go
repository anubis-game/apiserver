package connect

import "github.com/xh3b4sd/tracer"

var protocolMethodInvalidError = &tracer.Error{
	Kind: "protocolMethodInvalidError",
	Desc: "The request expects the signature method to be one of [dual-handshake user-challenge]. The signature method was not found to be one of those options. Therefore the request failed.",
}
