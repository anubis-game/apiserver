package server

import (
	"context"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

func (s *Server) listen() {
	var err error

	var srv *http.Server
	{
		srv = &http.Server{
			Handler: s.rtr,
		}
	}

	{
		s.log.Log(
			context.Background(),
			"level", "info",
			"message", "server listening for calls",
			"addr", s.lis.Addr().String(),
		)
	}

	{
		err = srv.Serve(s.lis)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}
}
