package schema

// Engine defines the game engine message type and instructs how to read those
// types of messages. Websocket messages coming from the game engine are
// prefixed with a client specific sequence byte and may contain multiple game
// messages.
type Engine byte

const (
	// Uuid is sent from the server to the client, with a buffer length of 22.
	// There are 2 additional parameters, the player's byte ID and their
	// associated wallet address.
	//
	//     [action]    [  id  ]    [ wallet ]
	//     [1 byte]    [1 byte]    [20 bytes]
	//
	Uuid Engine = 0x21

	// Size is sent from the server to the client, with a buffer length of 3.
	// There are 2 additional parameters, the player's byte ID and their
	// associated body part radius.
	//
	//     [action]    [  id  ]    [ size ]
	//     [1 byte]    [1 byte]    [1 byte]
	//
	Size Engine = 0x22

	// Type is sent from the server to the client, with a buffer length of 3.
	// There are 2 additional parameters, the player's byte ID and their
	// associated character type.
	//
	//     [action]    [  id  ]    [ type ]
	//     [1 byte]    [1 byte]    [1 byte]
	//
	Type Engine = 0x23

	// Body is sent from the server to the client, with a variable buffer length.
	// There are at least 3 additional parameters, the player's byte ID, the
	// amount of vector coordinates, and consecutively sets of 6 bytes for every
	// transmitted body part as defined by the second parameter. The amount
	// parameter is set to 3 in the example below.
	//
	//     [action]    [  id  ]    [amount]    [  x/y  ] [  x/y  ] [  x/y  ]
	//     [1 byte]    [1 byte]    [1 byte]    [6 bytes] [6 bytes] [6 bytes]
	//
	Body Engine = 0x24

	// Head is sent from the server to the client, with a buffer length of 3.
	// There are 2 additional parameters, the player's byte ID, and the
	// coordinates of the new head to add.
	//
	//     [action]    [  id  ]    [  x/y  ]
	//     [1 byte]    [1 byte]    [6 bytes]
	//
	Head Engine = 0x25

	// Tail is sent from the server to the client, with a buffer length of 3.
	// There are 2 additional parameters, the player's byte ID, and the
	// coordinates of the old tail to remove.
	//
	//     [action]    [  id  ]    [  x/y  ]
	//     [1 byte]    [1 byte]    [6 bytes]
	//
	Tail Engine = 0x26

	// Turn is sent from the client to the server, with a buffer length of 3.
	// There are 2 additional parameters, the quadrant and angle bytes indicating
	// the player's desired direction.
	//
	//     [action]    [ qudr ]    [ angl ]
	//     [1 byte]    [1 byte]    [1 byte]
	//
	Turn Engine = 0x27

	// Race is sent from the client to the server, with a buffer length of 1.
	// There are no additional parameters.
	//
	//     [action]
	//     [1 byte]
	//
	Race Engine = 0x28

	// Food is sent from the server to the client, with a variable buffer length.
	// There are at least 4 additional parameters, the amount of energy
	// coordinates, the energy coordinates themselves, the energy size and the
	// energy type. The amount parameter is set to 2 in the example below.
	//
	//     [action]    [amount]    [  x/y  ] [ size ] [ type ]    [  x/y  ] [ size ] [ type ]
	//     [1 byte]    [1 byte]    [6 bytes] [1 byte] [1 byte]    [6 bytes] [1 byte] [1 byte]
	//
	Food Engine = 0x29

	// Kill is sent from the server to the client, with a buffer length of 3.
	// There are 2 additional parameters, the winning player's byte ID, and the
	// losing player's byte ID.
	//
	//     [action]    [ wnid ]    [ lsid ]
	//     [1 byte]    [1 byte]    [1 byte]
	//
	Kill Engine = 0x2a
)
