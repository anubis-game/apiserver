package connect

import "github.com/xh3b4sd/tracer"

var challengeHeaderInvalidError = &tracer.Error{
	Kind: "challengeHeaderInvalidError",
	Desc: "The request expects the Sec-Websocket-Protocol header to contain exactly 2 elements. The Sec-Websocket-Protocol header was not found to comply with this requirement. Therefore the request failed.",
}

var challengeValidationFailedError = &tracer.Error{
	Kind: "challengeValidationFailedError",
	Desc: "The request expects the user-challenge to succeed for any player to participate in the game. The user-challenge did not provide a valid session token. Therefore the request failed.",
}

var sessionTokenLengthError = &tracer.Error{
	Kind: "sessionTokenLengthError",
	Desc: "The request expects the session token to have 36 characters. The session token was not found to have 36 characters. Therefore the request failed.",
}
