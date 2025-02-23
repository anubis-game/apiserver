package connect

import (
	"time"
)

func (h *Handler) Daemon() {
	go h.txp.Expire(time.Minute)
	go h.wxp.Expire(time.Minute)
}
