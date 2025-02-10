package vector

import "github.com/anubis-game/apiserver/pkg/setter"

// Motion contains all information necessary to calculate a player's change in
// direction.
type Motion struct {
	// Qdr is the index for the quadrant byte. This is one of [0x01, 0x02, 0x03,
	// 0x04], indicating the logical quadrant any given player is moving towards.
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
	// Agl is the index for the angle byte. This is the radian dividing 90° into
	// 256 equal parts. Given any quadrant, alpha is the clock wise angle encoded
	// as a single byte in the range of [0, 255], dividing 90 degrees of any
	// quadrant into 256 possible angles. The measurement of alpha starts at 0°
	// for quadrant 1, 90° for quadrant 2, 180° for quadrant 3, and 270° for
	// quadrant 4. Note that a single angle byte represents 0.3515625° within any
	// given quadrant.
	//
	//     ( 1 / 256 ) * 90° = 0.3515625°
	//
	Agl byte
	// Vlc is the index for the velocity byte. Every additional byte represents a
	// factor of 100%.
	Vlc byte
}

func (v *Vector) Motion() setter.Interface[Motion] {
	return v.mot
}
