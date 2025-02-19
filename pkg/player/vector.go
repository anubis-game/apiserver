package player

func (p *Player) Vector() []byte {
	return p.Vec.Encode()
}
