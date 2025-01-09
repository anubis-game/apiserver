package stream

import "github.com/xh3b4sd/tracer"

var challengeValidationFailedError = &tracer.Error{
	Kind: "challengeValidationFailedError",
	Desc: "The request expects the user-challenge to succeed for any player to participate in the game. The user-challenge did not provide a valid session token. Therefore the request failed.",
}
