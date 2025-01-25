package connect

import (
	"time"
)

func (h *Handler) Daemon() {
	{
		go h.ang.Daemon()
		go h.crd.Daemon()
		go h.qdr.Daemon()
		go h.txp.Expire(time.Minute)
		go h.wxp.Expire(time.Minute)
	}
}
