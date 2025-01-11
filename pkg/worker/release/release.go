package release

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

	// TODO return (pac, true) if pac.Timeout (int64) is still in the future

	var req bool
	{
		pac, req, err = r.release(pac)
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

func (r *Release) release(pac Packet) (Packet, bool, error) {
	// TODO release player
	return pac, false, nil
}
