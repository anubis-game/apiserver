package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/anubis-game/apiserver/pkg/runtime"
	"github.com/anubis-game/apiserver/pkg/stream"
	"github.com/anubis-game/apiserver/pkg/wallet"
	"github.com/coder/websocket"
	"github.com/gorilla/mux"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don chan struct{}
	Lis net.Listener
	Log logger.Interface
	Rtr *mux.Router
	Str *stream.Stream
}

type Server struct {
	ctx context.Context
	don chan struct{}
	lis net.Listener
	log logger.Interface
	opt *websocket.AcceptOptions
	rtr *mux.Router
	str *stream.Stream
}

func New(c Config) *Server {
	if c.Don == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Don must not be empty", c)))
	}
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

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var opt *websocket.AcceptOptions
	{
		opt = &websocket.AcceptOptions{
			InsecureSkipVerify: true, // TODO
			Subprotocols:       []string{"personal_sign"},
		}
	}

	return &Server{
		ctx: ctx,
		don: c.Don,
		lis: c.Lis,
		log: c.Log,
		opt: opt,
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

func (s *Server) Stream(w http.ResponseWriter, r *http.Request) error {
	var err error

	//
	//     https://stackoverflow.com/questions/4361173/http-headers-in-websockets-client-api/77060459#77060459
	//     https://github.com/kubernetes/kubernetes/commit/714f97d7baf4975ad3aa47735a868a81a984d1f0
	//
	var add string
	{
		add, err = s.verify(reqHea(r.Header["Sec-Websocket-Protocol"]))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var con *websocket.Conn
	{
		con, err = websocket.Accept(w, r, s.opt)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = s.stream(add, con)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (s *Server) stream(add string, con *websocket.Conn) error {
	var onc sync.Once

	var don chan struct{}
	var rep chan struct{}
	{
		don = make(chan struct{})
		rep = make(chan struct{})
	}

	var cli stream.Client
	{
		cli = stream.Client{
			Clo: func(add bool) {
				onc.Do(func() {
					// If close is called by Add, then we do not want to call Rem again in
					// the read loop below.
					if add {
						close(rep)
					}

					{
						close(don)
						con.CloseNow() //nolint:errcheck
					}
				})
			},
			Wri: func(typ websocket.MessageType, byt []byte) {
				err := con.Write(s.ctx, typ, byt)
				if err != nil {
					go s.str.Rem(add)
				}
			},
		}
	}

	{
		s.str.Add(add, cli)
	}

	go func() {
		for {
			typ, byt, err := con.Read(s.ctx)
			if err != nil {
				// If this connection is closed from the outside, then we want to remove
				// the client from our internal state. If the same client replaces
				// itself, then we are reading from a closed connection and do not want
				// to remove the client again.
				select {
				case <-rep:
					// fall through
				default:
					go s.str.Rem(add)
				}

				{
					return
				}
			} else {
				{
					go s.str.Wri(add, typ, byt)
				}
			}
		}
	}()

	select {
	case <-don:
	case <-s.don:
		s.str.Rem(add)
	}

	return nil
}

func (s *Server) verify(aut []string) (string, error) {
	var err error

	//
	//     aut[0] signature method
	//     aut[1] message hash
	//     aut[2] public key
	//     aut[3] signature hash
	//
	if len(aut) != 4 {
		return "", tracer.Mask(fmt.Errorf("auth failed: invalid format"))
	}

	var add string
	{
		add, err = wallet.Verify(aut[0], aut[1], aut[2], aut[3], time.Now().UTC())
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return add, nil
}

func linBrk(byt []byte) []byte {
	return append(byt, []byte("\n")...)
}

func reqHea(lis []string) []string {
	var spl []string

	for _, x := range lis {
		// Split comma-separated values and trim whitespace
		for _, y := range strings.Split(x, ",") {
			spl = append(spl, strings.TrimSpace(y))
		}
	}

	return spl
}
