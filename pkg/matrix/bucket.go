package matrix

import (
	"fmt"

	"github.com/xh3b4sd/tracer"
)

const (
	// Max is the upper byte boundary of any bucket and pixel. Buckets and pixels
	// above this level overflow.
	Max byte = 163
	// Min is the lower byte boundary of any bucket and pixel. Buckets and pixels
	// below this level underflow.
	Min byte = 100
	// Siz is the bucket size of the layered coordinate system. This is the total
	// amount of outer buckets within the entire coordinate system. This is also
	// the number of inner buckets within any given outer bucket. And this is the
	// quadratic length in pixels per inner bucket. Player positions may overflow
	// and underflow into other buckets, if players move beyond the boundaries of
	// their current position.
	Siz byte = Max - Min + 1
	// Dia is the truncated diagonal pixel length of every inner bucket, minus 1.
	// This is the maximum amount of movement allowed to happen within a single
	// frame. Any player movement beyond this pixel length is invalid, causing an
	// illegal overflow.
	//
	//     ( Siz * math.Sqrt2 ) - 1
	//
	Dia float64 = 88
)

const (
	X0 int = 0
	Y0 int = 1
	X1 int = 2
	Y1 int = 3
)

type Bucket [4]byte

func (b Bucket) Dec(siz byte) Bucket {
	del := Min + siz

	sx0, sy0, sx1, sy1 := b[X0], b[Y0], b[X1], b[Y1]

	if del > sx1 {
		sx0, sx1 = sx0-1, Max-(del-sx1)+1
	} else {
		sx1 -= siz
	}

	if del > sy1 {
		sy0, sy1 = sy0-1, Max-(del-sy1)+1
	} else {
		sy1 -= siz
	}

	return Bucket{sx0, sy0, sx1, sy1}
}

func (b Bucket) Inc(siz byte) Bucket {
	del := Max - siz

	sx0, sy0, sx1, sy1 := b[X0], b[Y0], b[X1], b[Y1]

	if del < sx1 {
		sx0, sx1 = sx0+1, Min+(sx1-del)-1
	} else {
		sx1 += siz
	}

	if del < sy1 {
		sy0, sy1 = sy0+1, Min+(sy1-del)-1
	} else {
		sy1 += siz
	}

	return Bucket{sx0, sy0, sx1, sy1}
}

func (b Bucket) Key() Bucket {
	// TODO return index bucket based on mod 4
	return Bucket{}
}

// TODO we can use the player's profile size to check what other inner buckets
// have to be added to the generated buffer region.
func (b Bucket) Ngh(qdr byte) [6]Bucket {
	switch qdr {
	case 0x01:
		return b.one()
	case 0x02:
		return b.two()
	case 0x03:
		return b.thr()
	case 0x04:
		return b.fou()
	default:
		tracer.Panic(fmt.Errorf("invalid quadrant %#v", qdr))
	}

	return [6]Bucket{}
}

func (b Bucket) one() [6]Bucket {
	//
	//     buf = [6]Bucket{
	//        ? , b02, b03,
	//        ? ,  b , b06,
	//             ? ,  ?
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), b[X0], b[Y0], b[X1], b[Y1], byte(0), byte(0), byte(0), byte(0)

	// from b05 to b02 by incrementing along y
	if cy1 == Max {
		ax0, ay0, ax1, ay1 = cx0, cy0+1, cx1, Min
	} else {
		ax0, ay0, ax1, ay1 = cx0, cy0, cx1, cy1+1
	}

	// from b02 to b03 by incrementing along x
	if ax1 == Max {
		bx0, by0, bx1, by1 = ax0+1, ay0, Min, ay1
	} else {
		bx0, by0, bx1, by1 = ax0, ay0, ax1+1, ay1
	}

	// from b05 to b06 by incrementing along x
	if cx1 == Max {
		dx0, dy0, dx1, dy1 = cx0+1, cy0, Min, cy1
	} else {
		dx0, dy0, dx1, dy1 = cx0, cy0, cx1+1, cy1
	}

	return [6]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}

func (b Bucket) two() [6]Bucket {
	//
	//     buf = [6]Bucket{
	//             ? ,  ?
	//        ? ,  b , b06,
	//        ? , b08, b09,
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := b[X0], b[Y0], b[X1], b[Y1], byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0)

	// from b05 to b06 by incrementing along x
	if ax1 == Max {
		bx0, by0, bx1, by1 = ax0+1, ay0, Min, ay1
	} else {
		bx0, by0, bx1, by1 = ax0, ay0, ax1+1, ay1
	}

	// from b05 to b08 by decrementing along y
	if ay1 == Min {
		cx0, cy0, cx1, cy1 = ax0, ay0-1, ax1, Max
	} else {
		cx0, cy0, cx1, cy1 = ax0, ay0, ax1, ay1-1
	}

	// from b06 to b09 by decrementing along y
	if by1 == Min {
		dx0, dy0, dx1, dy1 = bx0, by0-1, bx1, Max
	} else {
		dx0, dy0, dx1, dy1 = bx0, by0, bx1, by1-1
	}

	return [6]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}

func (b Bucket) thr() [6]Bucket {
	//
	//     buf = [6]Bucket{
	//        ? ,  ?
	//       b04,  b ,  ?
	//       b07, b08,  ?
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := byte(0), byte(0), byte(0), byte(0), b[X0], b[Y0], b[X1], b[Y1], byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0)

	// from b05 to b08 by decrementing along y
	if by1 == Min {
		dx0, dy0, dx1, dy1 = bx0, by0-1, bx1, Max
	} else {
		dx0, dy0, dx1, dy1 = bx0, by0, bx1, by1-1
	}

	// from b05 to b04 by decrementing along x
	if bx1 == Min {
		ax0, ay0, ax1, ay1 = bx0-1, by0, Max, by1
	} else {
		ax0, ay0, ax1, ay1 = bx0, by0, bx1-1, by1
	}

	// from b04 to b07 by decrementing along y
	if ay1 == Min {
		cx0, cy0, cx1, cy1 = ax0, ay0-1, ax1, Max
	} else {
		cx0, cy0, cx1, cy1 = ax0, ay0, ax1, ay1-1
	}

	return [6]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}

func (b Bucket) fou() [6]Bucket {
	//
	//     buf = [6]Bucket{
	//       b01, b02,  ?
	//       b04,  b ,  ?
	//        ? ,  ?
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), b[X0], b[Y0], b[X1], b[Y1]

	// from b05 to b04 by decrementing along x
	if dx1 == Min {
		cx0, cy0, cx1, cy1 = dx0-1, dy0, Max, dy1
	} else {
		cx0, cy0, cx1, cy1 = dx0, dy0, dx1-1, dy1
	}

	// from b05 to b02 by incrementing along y
	if dy1 == Max {
		bx0, by0, bx1, by1 = dx0, dy0+1, dx1, Min
	} else {
		bx0, by0, bx1, by1 = dx0, dy0, dx1, dy1+1
	}

	// from b04 to b01 by incrementing along y
	if cy1 == Max {
		ax0, ay0, ax1, ay1 = cx0, cy0+1, cx1, Min
	} else {
		ax0, ay0, ax1, ay1 = cx0, cy0, cx1, cy1+1
	}

	return [6]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}
