package engine

func (e *Engine) tick() {
	for u := range e.uni.Length() {
		var tur Turn
		{
			tur = e.tur[u]
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

		// ...
	}
}
