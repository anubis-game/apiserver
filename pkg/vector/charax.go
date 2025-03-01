package vector

import "github.com/anubis-game/apiserver/pkg/setter"

const (
	// Rad is the initial radius of a player's head and body parts.
	Rad byte = 10
	// Siz is the initial amount of points that a player is worth.
	Siz int = 50
)

type Charax struct {
	// Als is the angle limit for the maximum allowed range of motion per frame at
	// normal speed.
	Als byte
	// Alr is the angle limit for the maximum allowed range of motion per frame at
	// racing speed.
	Alr byte
	// Prt is the range of sight in partitions.
	Prt int
	// Rad is the current radius of a player's head and body parts.
	Rad byte
	// Siz is the total amount of points that a player is worth.
	Siz int
	// Typ is the player's style indicator.
	Typ byte
}

func (v *Vector) Charax() setter.Interface[Charax] {
	return v.crx
}
