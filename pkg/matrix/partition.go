package matrix

type Partition Coordinate // TODO:test unit test Partition.Byt and ensure the results are correct.

func (p Partition) Eql(prt Partition) bool {
	return p == prt
}

func (p Partition) Ins(stp int, srg int, sbt int, slf int) bool {
	return !(stp < p.Y || srg < p.X || sbt > p.Y || slf > p.X) // TODO:test partitions are detected to be inside screen boundaries
}

var zrp Partition

func (p Partition) Zer() bool {
	return p == zrp
}
