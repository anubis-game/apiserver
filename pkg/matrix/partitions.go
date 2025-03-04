package matrix

type Partitions []Partition

func (c Partitions) Len() int {
	return len(c)
}

func (c Partitions) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Partitions) Less(i, j int) bool {
	// Sort by X as first order, if X is not equal.

	if c[i].X != c[j].X {
		return c[i].X < c[j].X
	}

	// Sort by Y as second order, if X is equal.

	return c[i].Y < c[j].Y
}
