package vector

func roundI(f float64) int {
	if f >= 0 {
		return int(f + 0.5)
	}

	return int(f - 0.5)
}
