package matrix

type Partition Coordinate // TODO:test unit test Partition.Byt and ensure the results are correct.

// TODO:test make sure PfromS does not panic

func PfromS(top int, rig int, bot int, lef int) []Partition {
	ind := 0
	prt := make([]Partition, (((top-bot)/Prt)+1)*(((rig-lef)/Prt)+1))

	for y := top; y >= bot; y -= Prt {
		for x := lef; x <= rig; x += Prt {
			prt[ind].X = x
			prt[ind].Y = y
			ind++
		}
	}

	return prt
}

func (p Partition) Eql(prt Partition) bool {
	return p.X == prt.X && p.Y == prt.Y
}
