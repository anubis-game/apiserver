// Package schema defines the buffer encoding used between all client-server
// communication. Websocket messages from the game engine start with a sequence
// byte, followed by at least one game message. Clients can keep track of the
// sequence byte to ensure that every message was properly received in order.
// Every game message starts with an action byte, followed by action specific
// information.
//
//	[ sqnc ]    [action]    [  ...  ]    [action]    [  ...  ]
//	[1 byte]    [1 byte]    [N bytes]    [1 byte]    [N bytes]
package schema
