package schema

type Action byte

const (
	Ping Action = 0x01
	Pong Action = 0x02
	Auth Action = 0x03
	Cast Action = 0x04
	Move Action = 0x05
	Kill Action = 0x06
)
