package resolve

import (
	"context"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Release struct {
	log logger.Logger
}

func New(log logger.Logger) *Release {
	return &Release{
		log: log,
	}
}

func (r *Release) Ensure(pac Packet) (Packet, bool) {
	var err error

	// TODO return (pac, true) if pac.Timeout (int64) is still in the future

	var req bool
	{
		pac, req, err = r.ensure(pac)
		if err != nil {
			r.log.Log(
				context.Background(),
				"level", "error",
				"message", err.Error(),
				"stack", tracer.Stack(err),
			)
		}
	}

	return pac, req
}

func (r *Release) ensure(pac Packet) (Packet, bool, error) {
	// TODO resolve player
	return pac, false, nil
}
