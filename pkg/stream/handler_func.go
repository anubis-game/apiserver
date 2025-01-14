package stream

import (
	"net/http"
	"strings"

	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

func (s *Stream) HandlerFunc(w http.ResponseWriter, r *http.Request) error {
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
			wal, err = s.verify(hea)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	case schema.UserChallenge:
		wal, err = s.search(hea)
		if err != nil {
			return tracer.Mask(err)
		}
	default:
		return tracer.Maskf(protocolMethodInvalidError, "%s", hea[0])
	}

	{
		exi := s.cli.Exists(wal)
		if exi {
			return tracer.Mask(walletAddressRegisteredError)
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
		err = s.client(wal, con)
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
