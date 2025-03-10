package matrix

type Partition Coordinate // TODO:test unit test Partition.Byt and ensure the results are correct.

// TODO:test make sure Pt1Scr and Pt8Scr does not panic

func Pt1Scr(top int, rig int, bot int, lef int) []Partition {
	ind := 0
	scr := make([]Partition, (((top-bot)/Pt1)+1)*(((rig-lef)/Pt1)+1))

	for y := top; y >= bot; y -= Pt1 {
		for x := lef; x <= rig; x += Pt1 {
			scr[ind].X = x
			scr[ind].Y = y
			ind++
		}
	}

	return scr
}

func Pt8Scr(prt Partition) []Partition {
	top := prt.Y + Pt8
	rig := prt.X + Pt8
	bot := prt.Y - Pt8
	lef := prt.X - Pt8

	ind := 0
	scr := make([]Partition, 9)

	for y := top; y >= bot; y -= Pt8 {
		for x := lef; x <= rig; x += Pt8 {
			scr[ind].X = x
			scr[ind].Y = y
			ind++
		}
	}

	return scr
}

func (p Partition) Eql(prt Partition) bool {
	return p.X == prt.X && p.Y == prt.Y
}
