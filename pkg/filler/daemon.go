package filler

func (f *Filler) Daemon() {
	go f.ang.Daemon()
	go f.crd.Daemon()
	go f.qdr.Daemon()
}
