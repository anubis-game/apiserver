package connect

import (
	"net/http"
	"strconv"

	"github.com/xh3b4sd/tracer"
)

var tokenAlreadyExistsError = &tracer.Error{
	Kind: "tokenAlreadyExistsError",
	Code: strconv.Itoa(http.StatusInternalServerError),
	Desc: "The request expects the generated session token to be unique. The generated session token was found to exist already. Therefore the request failed.",
}
