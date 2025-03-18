package vector

// Motion contains a Vector's current direction of travel.
type Motion struct {
	// Qdr is the quadrant byte, which is one of [0x01 0x02 0x03 0x04], indicating
	// the logical quadrant a player is moving towards during an update cycle.
	//
	//                       0°
	//
	//             +---------+---------+
	//             |         |         |
	//             |    4    |    1    |
	//             |         |         |
	//     270°    +---------+---------+    90°
	//             |         |         |
	//             |    3    |    2    |
	//             |         |         |
	//             +---------+---------+
	//
	//                      180°
	//
	Qdr byte
	// Agl is the angle byte. This is the radian dividing 90° into 256 equal
	// parts. Given any quadrant, agl is the clock wise angle encoded as a single
	// byte in the range of [0, 255], dividing 90 degrees of any quadrant into 256
	// possible angles. The measurement of agl starts at 0° for quadrant 1, 90°
	// for quadrant 2, 180° for quadrant 3, and 270° for quadrant 4. Note that a
	// single angle byte represents 0.3515625° within any given quadrant.
	//
	//     ( 1 / 256 ) * 90° = 0.3515625°
	//
	Agl byte
}

func (v *Vector) Motion() Motion {
	return v.mot
}
