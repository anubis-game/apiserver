package schema

type Action byte

const (
	// Ping is sent from the client to the server. TODO
	//
	//     [0x1]
	//
	Ping Action = 0x1

	// Pong TODO
	//
	//     [0x2]
	//
	Pong Action = 0x2

	// Auth is sent from the server to the client, with a buffer length of 17.
	// There is one additional parameter, the granted session token.
	//
	//     [action]    [  uuid  ]
	//     [1 byte]    [16 bytes]
	//
	Auth Action = 0x3

	// Uuid is sent from the server to the client, with a buffer length of 23.
	// There are two additional parameters, the player's 2 byte UID and their
	// associated wallet address.
	//
	//     [action]    [  uid  ]    [ wallet ]
	//     [1 byte]    [2 bytes]    [20 bytes]
	//
	Uuid Action = 0x4

	// Size is sent from the server to the client, with a buffer length of 4.
	// There are two additional parameters, the player's 2 byte UID and their
	// associated body part radius.
	//
	//     [action]    [  uid  ]    [ size ]
	//     [1 byte]    [2 bytes]    [1 byte]
	//
	Size Action = 0x5

	// Type is sent from the server to the client, with a buffer length of 4.
	// There are two additional parameters, the player's 2 byte UID and their
	// associated character type.
	//
	//     [action]    [  uid  ]    [ type ]
	//     [1 byte]    [2 bytes]    [1 byte]
	//
	Type Action = 0x6

	// Body is sent from the server to the client, with a variable buffer length.
	// There are at least 3 additional parameters, the player's 2 byte UID, the
	// amount of vector coordinates, and consecutively sets of 6 bytes for every
	// transmitted body part as defined by the second parameter.
	//
	//     [action]    [  uid  ]    [amount]    [  x/y  ] [  x/y  ] [  x/y  ]
	//     [1 byte]    [2 bytes]    [1 byte]    [6 bytes] [6 bytes] [6 bytes]
	//
	Body Action = 0x7

	// Move
	//
	//     [0x8]
	//
	Move Action = 0x8

	// Food is sent from the server to the client, with a buffer length of 9.
	// There are three additional parameters, the energy coordinates, the energy
	// size and the energy type.
	//
	//     [action]    [  x/y  ]    [ size ]    [ type ]
	//     [1 byte]    [6 bytes]    [1 byte]    [1 byte]
	//
	Food Action = 0x9

	// Race is sent from the client to the server, with a buffer length of 1.
	// There are no additional parameters.
	//
	//     [0xa]
	//
	Race Action = 0xa

	// Kill
	//
	//     [0xb]
	//
	Kill Action = 0xb
)
