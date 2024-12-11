package stream

import (
	"net/http"
	"strings"

	"github.com/coder/websocket"
	"github.com/xh3b4sd/tracer"
)

func (s *Stream) HandlerFunc(w http.ResponseWriter, r *http.Request) error {
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
		err = s.process(add, con)
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
