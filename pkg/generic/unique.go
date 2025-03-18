package generic

// Unique returns the distinct items of the given list without preserving any
// particular input order. Unique does not allocate and is considered faster for
// lists with fewer than 100 items. T must not contain pointers. For larger
// lists the standard map based approach is recommended instead.
func Unique[T comparable](lis []T) []T {
	var r int

	for i := range lis {
		var n int

		for j := range lis {
			if lis[i-r] == lis[j] {
				{
					n++
				}

				if n > 1 {
					{
						lis[i-r] = lis[len(lis)-1]
						lis = lis[:len(lis)-1]
					}

					{
						r++
					}

					{
						break
					}
				}
			}
		}
	}

	return lis
}
