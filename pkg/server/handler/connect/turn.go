package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) turn(uid byte, cli *client.Client, byt []byte) error {
	// If we do not receive exactly two bytes, then we simply ignore the user
	// input. The two required bytes here are the quadrant byte and the angle
	// byte.

	if len(byt) != 2 {
		return tracer.Maskf(turnBytesInvalidError, "%d", len(byt))
	}

	// If the quadrant byte is not one of [1 2 3 4], then we simply ignore the
	// user input.

	if byt[0]-1 > 3 {
		return tracer.Maskf(turnQuadrantRangeError, "%#v", byt[0])
	}

	return h.rtr.Turn(uid, cli, byt)
}
