package connect

import "github.com/xh3b4sd/tracer"

var pingBufferInvalidError = &tracer.Error{
	Kind: "pingBufferInvalidError",
	Desc: "The request expects the ping message to contain exactly 1 roundtrip byte. The ping request was not found to comply with this requirement. Therefore the request failed.",
}
