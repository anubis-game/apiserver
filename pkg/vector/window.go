package vector

import "github.com/anubis-game/apiserver/pkg/window"

func (v *Vector) Window() *window.Window {
	return v.win
}
