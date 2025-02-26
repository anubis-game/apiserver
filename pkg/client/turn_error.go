package client

import (
	"github.com/xh3b4sd/tracer"
)

var turnBytesInvalidError = &tracer.Error{
	Kind: "turnBytesInvalidError",
	Desc: "The request expects exactly 3 input bytes, one byte for the action, one byte for the quadrant, and one byte for the angle. The input bytes were not found to comply with this requirement. Therefore the request failed.",
}

var turnQuadrantRangeError = &tracer.Error{
	Kind: "turnQuadrantRangeError",
	Desc: "The request expects the quadrant byte to be one of [0x1 0x2 0x3 0x4]. The quadrant byte was found to be out of range. Therefore the request failed.",
}
