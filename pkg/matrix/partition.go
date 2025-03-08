package matrix

type Partition Coordinate // TODO:test unit test Partition.Byt and ensure the results are correct.

func (p Partition) Byt() [CoordinateBytes]byte {
	return [CoordinateBytes]byte{
		byte(p.X / 4096),
		byte(p.Y / 4096),
		byte((p.X % 4096) / 64),
		byte((p.Y % 4096) / 64),
		byte(p.X % 64),
		byte(p.Y % 64),
	}
}

func (p Partition) Eql(prt Partition) bool {
	return p.X == prt.X && p.Y == prt.Y
}

func (p Partition) Zer() bool {
	return p.X == 0 && p.Y == 0
}
