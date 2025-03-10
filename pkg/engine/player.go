package engine

import "github.com/anubis-game/apiserver/pkg/router"

type player struct {
	// buf contains the fanout buffers ready to be sent out to every player during
	// the ticker based fanout procedure. Any respective byte slice may be empty,
	// or contain one, or multiple encoded messages.
	buf [][]byte
	// cli contains the fanout channels for every client. It is critically
	// important that modifications on cli are only done sequentially by a single
	// writer.
	cli []chan<- []byte
	// rac
	rac []byte
	// tur
	tur []router.Turn
}
