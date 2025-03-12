package engine

type player struct {
	// act contains all active flags for us to know whether players are still in
	// the game or not. Players may be part of the game without having a client
	// connected. In that case we want to keep reconciling all relevant game
	// state, even if we do not send any data to the disconnected client.
	act []bool
	// agl
	agl []byte
	// buf contains the fanout buffers ready to be sent out to every player during
	// the ticker based fanout procedure. Any respective byte slice may be empty,
	// or contain one, or multiple encoded messages.
	buf [][]byte
	// cli contains the fanout channels for every client. It is critically
	// important that modifications on cli are only done sequentially by a single
	// writer.
	cli []chan<- []byte
	// qdr
	qdr []byte
	// rac
	rac []byte
}

func newPlayer(c int) *player {

	var buf [][]byte
	{
		buf = make([][]byte, c)
	}

	// Pre-allocate a fanout buffer for every potential player using the same
	// sequence byte 0x0. This sequence byte is incremented every update cycle.
	//
	//     [
	//         [0x0, ...], // player 1
	//         [0x0, ...], // player 2
	//         ...
	//         [0x0, ...], // player N
	//     ]
	//

	for i := range buf {
		buf[i] = make([]byte, 1, 64)
	}

	return &player{
		act: make([]bool, c),
		agl: make([]byte, c),
		buf: buf,
		cli: make([]chan<- []byte, c),
		qdr: make([]byte, c),
		rac: make([]byte, c),
	}
}
