package filler

func (f *Filler) Daemon() {
	{
		go f.agl.Daemon()
		go f.crd.Daemon()
		go f.qdr.Daemon()
	}

	for {
		select {
		case <-f.don:
			return
		case f.vec <- f.vector():
			// repeat
		}
	}
}
