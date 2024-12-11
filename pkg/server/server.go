package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/anubis-game/apiserver/pkg/runtime"
	"github.com/anubis-game/apiserver/pkg/stream"
	"github.com/gorilla/mux"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Lis net.Listener
	Log logger.Interface
	Rtr *mux.Router
	Str *stream.Stream
}

type Server struct {
	lis net.Listener
	log logger.Interface
	rtr *mux.Router
	str *stream.Stream
}

func New(c Config) *Server {
	if c.Lis == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lis must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rtr == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rtr must not be empty", c)))
	}
	if c.Str == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Str must not be empty", c)))
	}

	return &Server{
		lis: c.Lis,
		log: c.Log,
		rtr: c.Rtr,
		str: c.Str,
	}
}

func (s *Server) Daemon() {
	var err error

	// Add a simple health check response to the root.
	{
		s.rtr.NewRoute().Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(linBrk([]byte("OK")))
		})
	}

	// Add the anubis streaming handler. All GET requests will be upgraded to
	// manage websocket connections.
	{
		s.rtr.NewRoute().Methods("GET").Path("/connect").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := s.str.HandlerFunc(w, r)
			if err != nil {
				s.log.Log(
					context.Background(),
					"level", "error",
					"message", err.Error(),
					"stack", tracer.Stack(err),
				)
			}
		})
	}

	// Add a simple version response for the runtime.
	{
		s.rtr.NewRoute().Methods("GET").Path("/version").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(linBrk(runtime.JSON()))
		})
	}

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

func linBrk(byt []byte) []byte {
	return append(byt, []byte("\n")...)
}
