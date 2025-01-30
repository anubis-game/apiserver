package connect

import (
	"errors"
	"net"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) client(wal common.Address, con *websocket.Conn) error {
	var err error

	var uid uuid.UUID
	{
		uid, err = uuid.NewRandom()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cli *client.Client
	{
		cli = client.New(client.Config{
			Con: con,
			Ctx: h.ctx,
			Wal: wal,
		})
	}

	// With setting up the client connection, we also setup an expiration callback
	// in order to limit the connection lifetime of every client.

	{
		h.wxp.Ensure(wal, h.ttl, func() {
			defer close(cli.Expiry())
		})
	}

	// Below we manage the connection specific reader loop. Any error occuring
	// here causes the connection to close, regardless where the underlying error
	// originated from. In some cases, reading from the websocket connection may
	// fail. In other cases some internal logic may cause the reader loop to
	// produce an error, either due to invalid reconciliation results, or
	// websocket writes.

	go func() {
		err := h.reader(con, uid, cli)
		if errors.Is(err, net.ErrClosed) {
			// fall through
		} else if err != nil {
			h.log.Log(
				"level", "error",
				"message", err.Error(),
				"stack", tracer.Stack(err),
			)
		}

		{
			close(cli.Reader())
		}
	}()

	// We block this client specific goroutine until either the client or the
	// server shuts down, each of which may happen for various reasons. Once a
	// signal got emitted to close this client connection, we remove the client
	// from all internal references, but keep the player in the game. The reason
	// for not removing players from games during disconnects is that connections
	// might be dropped intermittently. That means the client may very well come
	// back quickly using its auth token and continue playing their game.

	select {
	case <-cli.Expiry():
	case <-cli.Reader():
	case <-cli.Writer():
	case <-h.don:
	}

	{
		h.wxp.Delete(wal)
	}

	{
		err := con.CloseNow()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *Handler) reader(con *websocket.Conn, uid uuid.UUID, cli *client.Client) error {
	for {
		_, byt, err := con.Read(h.ctx)
		if err != nil {
			return tracer.Mask(err)
		}

		switch schema.Action(byt[0]) {
		case schema.Ping:
			err = h.ping(uid, cli, byt)
		case schema.Auth:
			err = h.auth(uid, cli, byt)
		case schema.Join:
			err = h.join(uid, cli, byt)
		case schema.Move:
			err = h.move(uid, cli, byt)
		case schema.Race:
			err = h.race(uid, cli, byt)
		}

		if err != nil {
			return tracer.Mask(err)
		}
	}
}
