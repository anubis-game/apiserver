package object

type Sorter []Object

func (s Sorter) Len() int {
	return len(s)
}

func (s Sorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sorter) Less(i, j int) bool {
	// Sort by X as first order, if X is not equal.

	if s[i].X != s[j].X {
		return s[i].X < s[j].X
	}

	// Sort by Y as second order, if X is equal.

	return s[i].Y < s[j].Y
}
