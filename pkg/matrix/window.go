package matrix

const (
	BL int = 0
	TR int = 1
)

type Window [2]Bucket

// Has returns whether the given bucket resides inside the underlying Window. So
// if bck turns out to be outside of w, then Has returns false.
func (w Window) Has(bck Bucket) bool {
	wbl, wtr := w[BL], w[TR]

	lx0, ly0, rx0, ty0 := wbl[X0], wbl[Y0], wtr[X0], wtr[Y0]
	kx0, ky0, kx1, ky1 := bck[X0], bck[Y0], bck[X1], bck[Y1]

	if lx0 > kx0 || ly0 > ky0 || rx0 < kx0 || ty0 < ky0 {
		return false
	}

	if lx0 == kx0 && wbl[X1] > kx1 {
		return false
	}

	if ly0 == ky0 && wbl[Y1] > ky1 {
		return false
	}

	if rx0 == kx0 && wtr[X1] < kx1 {
		return false
	}

	if ty0 == ky0 && wtr[Y1] < ky1 {
		return false
	}

	return true
}
