package schema

type Action byte

const (
	// Ping is sent from the client to the server, with a buffer length of 2.
	// There is one additional parameter, the roundtrip byte used to identify the
	// ping request.
	//
	//     [action]    [ seed ]
	//     [1 byte]    [1 byte]
	//
	Ping Action = 0x1

	// Pong is sent from the server to the client, with a buffer length of 2.
	// There is one additional parameter, the roundtrip byte as received from the
	// ping request.
	//
	//     [action]    [ seed ]
	//     [1 byte]    [1 byte]
	//
	Pong Action = 0x2

	// Auth is sent from the server to the client, with a buffer length of 17.
	// There is one additional parameter, the granted session token.
	//
	//     [action]    [  uuid  ]
	//     [1 byte]    [16 bytes]
	//
	Auth Action = 0x3

	// Uuid is sent from the server to the client, with a buffer length of 22.
	// There are two additional parameters, the player's byte ID and their
	// associated wallet address.
	//
	//     [action]    [  id  ]    [ wallet ]
	//     [1 byte]    [1 byte]    [20 bytes]
	//
	Uuid Action = 0x4

	// Size is sent from the server to the client, with a buffer length of 3.
	// There are two additional parameters, the player's byte ID and their
	// associated body part radius.
	//
	//     [action]    [  id  ]    [ size ]
	//     [1 byte]    [1 byte]    [1 byte]
	//
	Size Action = 0x5

	// Type is sent from the server to the client, with a buffer length of 3.
	// There are two additional parameters, the player's byte ID and their
	// associated character type.
	//
	//     [action]    [  id  ]    [ type ]
	//     [1 byte]    [1 byte]    [1 byte]
	//
	Type Action = 0x6

	// Body is sent from the server to the client, with a variable buffer length.
	// There are at least 3 additional parameters, the player's byte ID, the
	// amount of vector coordinates, and consecutively sets of 6 bytes for every
	// transmitted body part as defined by the second parameter. The buffer length
	// of a a new Vector is 33.
	//
	//     [action]    [  id  ]    [amount]    [  x/y  ] [  x/y  ] [  x/y  ]
	//     [1 byte]    [1 byte]    [1 byte]    [6 bytes] [6 bytes] [6 bytes]
	//
	Body Action = 0x7

	// Turn is sent from the client to the server, with a buffer length of 3.
	// There are two additional parameters, the quadrant and angle bytes
	// indicating the player's desired direction.
	//
	//     [action]    [ qudr ]    [ angl ]
	//     [1 byte]    [1 byte]    [1 byte]
	//
	Turn Action = 0x8

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

	// Kill TODO:game
	//
	//     [0xb]
	//
	Kill Action = 0xb
)
