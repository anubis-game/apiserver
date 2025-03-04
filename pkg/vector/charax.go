package vector

type Charax struct {
	// Aln, angle limit normal, is the angle limit for the maximum allowed range
	// of motion per frame at normal speed.
	Aln byte
	// Alr, angle limit racing, is the angle limit for the maximum allowed range
	// of motion per frame at racing speed.
	Alr byte
	// Fos is the factor of sight in partitions.
	Fos int
	// Rad is the current radius of a player's head and body parts.
	Rad byte
	// Siz is the total amount of points that a player is worth.
	Siz int
	// Typ is the player's style indicator.
	Typ byte
}

func (v *Vector) Charax() Charax {
	return v.crx
}
