package schema

type Action byte

const (
	Ping Action = 1
	Pong Action = 2
	Auth Action = 3
	Join Action = 4
	Body Action = 5
	Move Action = 6
	Food Action = 7
	Race Action = 8
	Kill Action = 9
)
