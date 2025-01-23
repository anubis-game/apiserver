package stream

import (
	"time"
)

func (s *Stream) Daemon() {
	{
		go s.ang.Daemon()
		go s.crd.Daemon()
		go s.qdr.Daemon()
		go s.txp.Expire(time.Minute)
		go s.wxp.Expire(time.Minute)
	}

	for {
		select {
		case <-s.rtr.Closer:
			return
		case x := <-s.rtr.Create:
			s.create(x)
		case x := <-s.rtr.Delete:
			s.delete(x)
		case <-s.rtr.Update:
			s.update()
		}
	}
}
