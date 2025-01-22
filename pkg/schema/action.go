package schema

type Action byte

const (
	Ping Action = 1
	Pong Action = 2
	Auth Action = 3
	Join Action = 4
	Move Action = 5
	Race Action = 6
	Kill Action = 7
	Food Action = 8
)
