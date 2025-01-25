package server

import (
	"net/http"

	"github.com/anubis-game/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

func (s *Server) router() {
	// Add a simple health check response to the root.
	{
		s.rtr.NewRoute().Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			{
				w.Header().Set("Access-Control-Allow-Origin", "*") // TODO change this for production
				w.Header().Set("Access-Control-Allow-Methods", "GET")
				w.Header().Set("Content-Type", "plain/text")
			}

			{
				w.WriteHeader(http.StatusOK)
			}

			{
				_, _ = w.Write(linBrk([]byte("OK")))
			}
		})
	}

	// Add the anubis streaming handler. All GET requests will be upgraded to
	// manage websocket connections.
	{
		s.rtr.NewRoute().Methods("GET").Path("/connect").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := s.con.HandlerFunc(w, r)
			if err != nil {
				s.log.Log(
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
			{
				w.Header().Set("Access-Control-Allow-Origin", "*") // TODO change this for production
				w.Header().Set("Access-Control-Allow-Methods", "GET")
				w.Header().Set("Content-Type", "application/json")
			}

			{
				w.WriteHeader(http.StatusOK)
			}

			{
				_, _ = w.Write(linBrk(runtime.Json()))
			}
		})
	}
}

func linBrk(byt []byte) []byte {
	return append(byt, []byte("\n")...)
}
