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

func (r *Release) Router(pac Packet) (Packet, bool) {
	var err error

	var req bool
	{
		pac, req, err = r.resolve(pac)
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

func (r *Release) resolve(pac Packet) (Packet, bool, error) {
	// TODO resolve player
	return pac, false, nil
}
