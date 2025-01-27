package matrix

const (
	Q1 int = 1
	Q2 int = 3
	Q3 int = 2
	Q4 int = 0
)

// TODO we can use the player's profile size to check what other inner buckets
// have to be added to the generated buffer region.
func Buffer(ogn Bucket, qdr byte) [4]Bucket {
	switch qdr {
	case 0x01:
		return qdrOne(ogn)
	case 0x02:
		return qdrTwo(ogn)
	case 0x03:
		return qdrThr(ogn)
	case 0x04:
		return qdrFou(ogn)
	}

	return [4]Bucket{}
}

func qdrOne(ogn Bucket) [4]Bucket {
	//
	//     buf = [4]Bucket{
	//       q04, q01,
	//       ogn, q02,
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), ogn[X0], ogn[Y0], ogn[X1], ogn[Y1], byte(0), byte(0), byte(0), byte(0)

	// from q03 to q04 by incrementing along y
	if cy1 == Max {
		ax0, ay0, ax1, ay1 = cx0, cy0+1, cx1, Min
	} else {
		ax0, ay0, ax1, ay1 = cx0, cy0, cx1, cy1+1
	}

	// from q04 to q01 by incrementing along x
	if ax1 == Max {
		bx0, by0, bx1, by1 = ax0+1, ay0, Min, ay1
	} else {
		bx0, by0, bx1, by1 = ax0, ay0, ax1+1, ay1
	}

	// from q03 to q02 by incrementing along x
	if cx1 == Max {
		dx0, dy0, dx1, dy1 = cx0+1, cy0, Min, cy1
	} else {
		dx0, dy0, dx1, dy1 = cx0, cy0, cx1+1, cy1
	}

	return [4]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}

func qdrTwo(ogn Bucket) [4]Bucket {
	//
	//     buf = [4]Bucket{
	//       ogn, q01,
	//       q03, q02,
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := ogn[X0], ogn[Y0], ogn[X1], ogn[Y1], byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0)

	// from q04 to q01 by incrementing along x
	if ax1 == Max {
		bx0, by0, bx1, by1 = ax0+1, ay0, Min, ay1
	} else {
		bx0, by0, bx1, by1 = ax0, ay0, ax1+1, ay1
	}

	// from q04 to q03 by decrementing along y
	if ay1 == Min {
		cx0, cy0, cx1, cy1 = ax0, ay0-1, ax1, Max
	} else {
		cx0, cy0, cx1, cy1 = ax0, ay0, ax1, ay1-1
	}

	// from q01 to q02 by decrementing along y
	if by1 == Min {
		dx0, dy0, dx1, dy1 = bx0, by0-1, bx1, Max
	} else {
		dx0, dy0, dx1, dy1 = bx0, by0, bx1, by1-1
	}

	return [4]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}

func qdrThr(ogn Bucket) [4]Bucket {
	//
	//     buf = [4]Bucket{
	//       q04, ogn,
	//       q03, q02,
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := byte(0), byte(0), byte(0), byte(0), ogn[X0], ogn[Y0], ogn[X1], ogn[Y1], byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0)

	// from q01 to q02 by decrementing along y
	if by1 == Min {
		dx0, dy0, dx1, dy1 = bx0, by0-1, bx1, Max
	} else {
		dx0, dy0, dx1, dy1 = bx0, by0, bx1, by1-1
	}

	// from q01 to q04 by decrementing along x
	if bx1 == Min {
		ax0, ay0, ax1, ay1 = bx0-1, by0, Max, by1
	} else {
		ax0, ay0, ax1, ay1 = bx0, by0, bx1-1, by1
	}

	// from q04 to q03 by decrementing along y
	if ay1 == Min {
		cx0, cy0, cx1, cy1 = ax0, ay0-1, ax1, Max
	} else {
		cx0, cy0, cx1, cy1 = ax0, ay0, ax1, ay1-1
	}

	return [4]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}

func qdrFou(ogn Bucket) [4]Bucket {
	//
	//     buf = [4]Bucket{
	//       q04, q01,
	//       q03, ogn,
	//     }
	//

	ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx0, cy0, cx1, cy1, dx0, dy0, dx1, dy1 := byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), ogn[X0], ogn[Y0], ogn[X1], ogn[Y1]

	// from q02 to q03 by decrementing along x
	if dx1 == Min {
		cx0, cy0, cx1, cy1 = dx0-1, dy0, Max, dy1
	} else {
		cx0, cy0, cx1, cy1 = dx0, dy0, dx1-1, dy1
	}

	// from q02 to q01 by incrementing along y
	if dy1 == Max {
		bx0, by0, bx1, by1 = dx0, dy0+1, dx1, Min
	} else {
		bx0, by0, bx1, by1 = dx0, dy0, dx1, dy1+1
	}

	// from q03 to q04 by incrementing along y
	if cy1 == Max {
		ax0, ay0, ax1, ay1 = cx0, cy0+1, cx1, Min
	} else {
		ax0, ay0, ax1, ay1 = cx0, cy0, cx1, cy1+1
	}

	return [4]Bucket{
		{ax0, ay0, ax1, ay1},
		{bx0, by0, bx1, by1},
		{cx0, cy0, cx1, cy1},
		{dx0, dy0, dx1, dy1},
	}
}
