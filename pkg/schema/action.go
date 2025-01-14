package schema

type Action byte

const (
	Ping Action = 0x01
	Pong Action = 0x02
	Auth Action = 0x03
	Join Action = 0x04
	Cast Action = 0x05
	Move Action = 0x06
	Kill Action = 0x07
)
