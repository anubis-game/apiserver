package vector

import "github.com/anubis-game/apiserver/pkg/setter"

const (
	// Rad is the initial radius of a player's head and body parts.
	Rad byte = 10
	// Siz is the initial amount of points that a player is worth.
	Siz byte = 50
)

type Charax struct {
	// Rad is the current radius of a player's head and body parts.
	Rad byte
	// Siz is the total amount of points that a player is worth.
	Siz byte
	// Typ is the player's style indicator.
	Typ byte
}

func (v *Vector) Charax() setter.Interface[Charax] {
	return v.crx
}
