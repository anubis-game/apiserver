package engine

import "github.com/anubis-game/apiserver/pkg/router"

func (e *Engine) tick() {
	for u := range e.uni.Length() {
		// Skip all inactive players.

		if e.ply.cli[u] == nil {
			continue
		}

		var tur router.Turn
		{
			tur = e.ply.tur[u]
		}

		// It may happen that a new player is being processed here, while said
		// player has not yet provided their own motion specific update. In such a
		// case the new player's quadrant byte is still empty, forcing us to move
		// the player along the game map using the current direction of travel. We
		// will consider the player choise of movement as soon as they provide their
		// own motion update.

		if tur.Qdr == 0 {
			_ = 0 // TODO:game use current motion
		}

		// TODO:infra manage all changes in movement
	}
}
