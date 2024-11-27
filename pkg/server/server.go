package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/anubis-game/apiserver/pkg/runtime"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Lis net.Listener
	Log logger.Interface
	Rtr *mux.Router
	Sig chan os.Signal
	Upg websocket.Upgrader
}

type Server struct {
	lis net.Listener
	log logger.Interface
	rtr *mux.Router
	sig chan os.Signal
	upg websocket.Upgrader
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
	if c.Sig == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sig must not be empty", c)))
	}

	return &Server{
		lis: c.Lis,
		log: c.Log,
		rtr: c.Rtr,
		sig: c.Sig,
		upg: c.Upg,
	}
}

func (s *Server) Daemon() {
	var err error

	// Add a simple health check response to the root.
	{
		s.rtr.NewRoute().Methods("GET").Path("/").HandlerFunc(func(wri http.ResponseWriter, req *http.Request) {
			wri.WriteHeader(http.StatusOK)
			_, _ = wri.Write(linBrk([]byte("OK")))
		})
	}

	// Add the anubis streaming handler. All GET requests will be upgraded to
	// manage websocket connections.
	{
		s.rtr.NewRoute().Methods("GET").Path("/anubis").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := s.Stream(w, r)
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
		s.rtr.NewRoute().Methods("GET").Path("/version").HandlerFunc(func(wri http.ResponseWriter, req *http.Request) {
			wri.Header().Set("Content-Type", "application/json")
			wri.WriteHeader(http.StatusOK)
			_, _ = wri.Write(linBrk(runtime.JSON()))
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

func (s *Server) Stream(w http.ResponseWriter, r *http.Request) error {
	var err error

	var con *websocket.Conn
	{
		con, err = s.upg.Upgrade(w, r, nil)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		defer con.Close()
	}

	for {
		var typ int
		var mes []byte
		{
			typ, mes, err = con.ReadMessage()
			if _, ok := err.(*websocket.CloseError); ok {
				return nil
			} else if err != nil {
				return tracer.Mask(err)
			}
		}

		// TODO process streaming data

		{
			err = con.WriteMessage(typ, mes)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}
}

func linBrk(byt []byte) []byte {
	return append(byt, []byte("\n")...)
}
