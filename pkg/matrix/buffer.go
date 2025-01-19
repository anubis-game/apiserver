package matrix

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

	var buf [4]Bucket
	{
		buf[Q3] = ogn
	}

	{
		buf[Q4] = ogn
		buf[Q4][Y0], buf[Q4][Y1] = incByt(ogn[Y0], ogn[Y1])
	}

	{
		buf[Q1] = buf[Q4]
		buf[Q1][X0], buf[Q1][X1] = incByt(buf[Q4][X0], buf[Q4][X1])
	}

	{
		buf[Q2] = ogn
		buf[Q2][X0], buf[Q2][X1] = incByt(ogn[X0], ogn[X1])
	}

	return buf
}

func qdrTwo(ogn Bucket) [4]Bucket {
	//
	//     buf = [4]Bucket{
	//       ogn, q01,
	//       q03, q02,
	//     }
	//

	var buf [4]Bucket
	{
		buf[Q4] = ogn
	}

	{
		buf[Q1] = ogn
		buf[Q1][X0], buf[Q1][X1] = incByt(ogn[X0], ogn[X1])
	}

	{
		buf[Q2] = buf[Q1]
		buf[Q2][Y0], buf[Q2][Y1] = decByt(buf[Q1][Y0], buf[Q1][Y1])
	}

	{
		buf[Q3] = ogn
		buf[Q3][Y0], buf[Q3][Y1] = decByt(ogn[Y0], ogn[Y1])
	}

	return buf
}

func qdrThr(ogn Bucket) [4]Bucket {
	//
	//     buf = [4]Bucket{
	//       q04, ogn,
	//       q03, q02,
	//     }
	//

	var buf [4]Bucket
	{
		buf[Q1] = ogn
	}

	{
		buf[Q2] = ogn
		buf[Q2][Y0], buf[Q2][Y1] = decByt(ogn[Y0], ogn[Y1])
	}

	{
		buf[Q3] = buf[Q2]
		buf[Q3][X0], buf[Q3][X1] = decByt(buf[Q2][X0], buf[Q2][X1])
	}

	{
		buf[Q4] = ogn
		buf[Q4][X0], buf[Q4][X1] = decByt(ogn[X0], ogn[X1])
	}

	return buf
}

func qdrFou(ogn Bucket) [4]Bucket {
	//
	//     buf = [4]Bucket{
	//       q04, q01,
	//       q03, ogn,
	//     }
	//

	var buf [4]Bucket
	{
		buf[Q2] = ogn
	}

	{
		buf[Q3] = ogn
		buf[Q3][X0], buf[Q3][X1] = decByt(ogn[X0], ogn[X1])
	}

	{
		buf[Q4] = buf[Q3]
		buf[Q4][Y0], buf[Q4][Y1] = incByt(buf[Q3][Y0], buf[Q3][Y1])
	}

	{
		buf[Q1] = ogn
		buf[Q1][Y0], buf[Q1][Y1] = incByt(ogn[Y0], ogn[Y1])
	}

	return buf
}
