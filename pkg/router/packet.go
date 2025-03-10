package router

import "github.com/ethereum/go-ethereum/common"

const (
	Join = 0x1
	Drop = 0x0
)

// Turn contains a player's desired direction of travel.
type Turn struct {
	Uid byte
	Qdr byte
	Agl byte
}

type Uuid struct {
	// Uid
	Uid byte
	// Jod is the join-or-drop flag.
	//
	//     join: 0x1
	//     drop: 0x0
	//
	Jod byte
	// Wal
	Wal common.Address
	// Cli is the fanout channel provided by a client to receive fanout buffers.
	// Using a non-blocking channel allows us to decouple the ticker based fanout
	// procedure from the blocking operations of a client connection.
	Cli chan<- []byte
}
