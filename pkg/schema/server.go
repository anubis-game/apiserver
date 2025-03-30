package schema

// Server defines the server handler message type and instructs how to read
// those types of messages. Websocket messages coming from the server handler
// are not prefixed with a client sequence byte and only contain one server
// message.
type Server byte

const (
	// Ping is sent from the client to the server, with a buffer length of 2.
	// There is 1 additional parameter, the roundtrip byte used to identify the
	// ping request.
	//
	//     [action]    [  rt  ]
	//     [1 byte]    [1 byte]
	//
	Ping Server = 0x11

	// Pong is sent from the server to the client, with a buffer length of 2.
	// There is 1 additional parameter, the roundtrip byte as received from the
	// ping request.
	//
	//     [action]    [  rt  ]
	//     [1 byte]    [1 byte]
	//
	Pong Server = 0x12

	// Auth is sent from the server to the client, with a buffer length of 17.
	// There is 1 additional parameter, the granted session token.
	//
	//     [action]    [  auth  ]
	//     [1 byte]    [16 bytes]
	//
	Auth Server = 0x13
)
