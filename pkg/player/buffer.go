package player

import "github.com/anubis-game/apiserver/pkg/setter"

func (p *Player) Buffer() setter.Interface[[]byte] {
	return p.buf
}
