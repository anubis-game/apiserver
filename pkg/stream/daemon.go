package stream

import "time"

func (s *Stream) Daemon() {
	go s.txp.Expire(time.Minute)
	go s.wxp.Expire(time.Minute)
}
