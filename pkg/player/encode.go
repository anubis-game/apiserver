package player

import "github.com/anubis-game/apiserver/pkg/object"

func (p Player) EncodeBuffer(prt object.Object) []byte {
	buf := p.Vec.Buffer(prt)
	byt := make([]byte, 2+len(buf))

	copy(byt[0:2], p.Uid[:])
	copy(byt[2:], buf)

	return byt
}

func (p Player) EncodeVector() []byte {
	vec := p.Vec.Encode()
	byt := make([]byte, 8+len(vec))
	crx := p.Vec.Charax().Get()
	mot := p.Vec.Motion().Get()

	copy(byt[0:2], p.Uid[:])

	byt[2] = byte(crx.Rad)
	byt[3] = byte(crx.Siz)
	byt[4] = crx.Typ

	byt[5] = mot.Qdr
	byt[6] = mot.Agl
	byt[7] = mot.Vlc

	copy(byt[8:], vec)

	return byt
}

func (p Player) EncodeWallet() []byte {
	byt := make([]byte, 22)

	copy(byt[:2], p.Uid[:])
	copy(byt[2:], p.Cli.Wallet().Bytes())

	return byt
}
