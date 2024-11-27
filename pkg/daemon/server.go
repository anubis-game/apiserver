package daemon

import (
	"net"
	"net/http"

	"github.com/anubis-game/apiserver/pkg/server"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/xh3b4sd/tracer"
)

func (d *Daemon) Server() *server.Server {
	var err error

	var lis net.Listener
	{
		lis, err = net.Listen("tcp", net.JoinHostPort(d.env.HttpHost, d.env.HttpPort))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var rtr *mux.Router
	{
		rtr = mux.NewRouter()
	}

	var upg websocket.Upgrader
	{
		upg = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }, // Adjust for production
		}
	}

	return server.New(server.Config{
		Lis: lis,
		Log: d.log,
		Rtr: rtr,
		Sig: d.sig,
		Upg: upg,
	})
}
