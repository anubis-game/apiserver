package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Circle generates new coordinates around the given root coordinate in a
// circular fashion based on the provided budget and deducatable unit size. The
// given callback is executed for the given root coordinate plus all derived
// circular coordinates surrounding the root. Every time the given callback
// returns true the provided unit size is deducated from the provided budget.
// That means every time the provided callback returns false, an additional
// iteration is being done until the entire budget got consumed.
func Circle(crd matrix.Coordinate, bud int, unt int, fnc func(matrix.Coordinate) bool) {
	var cpy matrix.Coordinate
	var dis float64
	var qdr byte
	var agl byte
	{
		cpy = crd
		dis = nrm
		qdr = 1
	}

	for bud >= 0 {
		// Executing the callback first allows us to use the injected coordinate
		// itself without any modification. If the callback returns true, then we
		// deduct the given unit size from the given budget.

		if fnc(cpy) {
			bud -= unt
		}

		// Apply the coordinate change to the copy based on the original version.

		{
			cpy.X, cpy.Y = target(crd.X, crd.Y, qdr, agl, dis)
		}

		// Move the angle along.

		{
			agl += 128
		}

		// Manage the quadrant byte if the angle resets at zero again. Once we made
		// a full circle we also increment the target distance.

		if agl == 0 {
			if qdr < 4 {
				qdr++
			} else {
				qdr = 1
				dis += nrm
			}
		}
	}
}
