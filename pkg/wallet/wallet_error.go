package wallet

import (
	"github.com/xh3b4sd/tracer"
)

var signatureMethodInvalidError = &tracer.Error{
	Kind: "signatureMethodInvalidError",
	Desc: "The request expects the signature method to be one of [personal_sign]. The signature method was not found to be one of those options. Therefore the request failed.",
}

var messageHashEmptyError = &tracer.Error{
	Kind: "messageHashEmptyError",
	Desc: "The request expects the message hash not to be empty. The message hash was found to be empty. Therefore the request failed.",
}

var messageHashFormatError = &tracer.Error{
	Kind: "messageHashFormatError",
	Desc: `The request expects the message hash to be in the format "signer-[wallet address]-[unix seconds]". The message hash was not found to be in that format. Therefore the request failed.`,
}

var publicKeyEmptyError = &tracer.Error{
	Kind: "publicKeyEmptyError",
	Desc: "The request expects the public key not to be empty. The public key was found to be empty. Therefore the request failed.",
}

var publicKeyLengthError = &tracer.Error{
	Kind: "publicKeyLengthError",
	Desc: "The request expects the public key to have 132 characters. The public key was not found to have 132 characters. Therefore the request failed.",
}

var signatureHashEmptyError = &tracer.Error{
	Kind: "signatureHashEmptyError",
	Desc: "The request expects the signature hash not to be empty. The signature hash was found to be empty. Therefore the request failed.",
}

var signatureHashExpiredError = &tracer.Error{
	Kind: "signatureHashExpiredError",
	Desc: "The request expects the signature time to not be older than 1 minute. The signature time was found to be older than 1 minute. Therefore the request failed.",
}

var signatureHashInvalidError = &tracer.Error{
	Kind: "signatureHashInvalidError",
	Desc: "The request expects the signature hash to be valid. The signature hash was not found to be valid. Therefore the request failed.",
}

var signatureHashLengthError = &tracer.Error{
	Kind: "signatureHashLengthError",
	Desc: "The request expects the signature hash to have 132 characters. The signature hash was not found to have 132 characters. Therefore the request failed.",
}
