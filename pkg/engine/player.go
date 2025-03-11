package engine

type player struct {
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
