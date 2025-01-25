package connect

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) HandlerFunc(w http.ResponseWriter, r *http.Request) error {
	var err error

	// We use a semaphore pattern in order to ensure a concurrent connection
	// limit. If there is an available slot for this request handler, then this
	// free slot may be occupied by the calling client, and we process the
	// request. If no slot is available, then we return an error response.
	select {
	case h.sem <- struct{}{}:
		{
			err = h.handlerFunc(w, r)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			<-h.sem
		}
	default:
		http.Error(
			w,
			fmt.Sprintf("%s: %d", http.StatusText(http.StatusTooManyRequests), cap(h.sem)),
			http.StatusTooManyRequests,
		)
	}

	return nil
}

func (h *Handler) handlerFunc(w http.ResponseWriter, r *http.Request) error {
	var err error

	// The reuqest headers contain the desired protocol method as well as various
	// credentials used to authorize stream engine clients. See the links below
	// for more information about what we are doing here and why.
	//
	//     https://stackoverflow.com/questions/4361173/http-headers-in-websockets-client-api/77060459#77060459
	//     https://github.com/kubernetes/kubernetes/commit/714f97d7baf4975ad3aa47735a868a81a984d1f0
	//
	var hea []string
	{
		hea = reqHea(r.Header["Sec-Websocket-Protocol"])
	}

	var wal common.Address
	switch schema.Header(hea[0]) {
	case schema.DualHandshake:
		{
			wal, err = h.verify(hea)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	case schema.UserChallenge:
		wal, err = h.search(hea)
		if err != nil {
			return tracer.Mask(err)
		}
	default:
		return tracer.Maskf(protocolMethodInvalidError, "%s", hea[0])
	}

	{
		exi := h.wxp.Exists(wal)
		if exi {
			return tracer.Mask(walletAddressRegisteredError)
		}
	}

	var con *websocket.Conn
	{
		con, err = websocket.Accept(w, r, h.opt)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.client(wal, con)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func reqHea(lis []string) []string {
	var spl []string

	for _, x := range lis {
		for _, y := range strings.Split(x, ",") {
			spl = append(spl, strings.TrimSpace(y))
		}
	}

	return spl
}
