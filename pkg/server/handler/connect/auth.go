package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) auth(_ byte, cli *client.Client, _ []byte) error {
	var err error

	// Create a new session token using V4 UUIDs for the requesting Wallet
	// address. For that we also want to check for the newly generated session
	// token to exist and handle errors properly, if they do, instead of blindly
	// overwriting an existing value eventually. If such a case were ever to
	// happen, then we are dealing with a rather unlucky, but severe internal
	// error.

	var tok uuid.UUID
	{
		tok, err = uuid.NewRandom()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		exi := h.tok.Exists(tok)
		if exi {
			return tracer.Mask(tokenAlreadyExistsError)
		}
	}

	// We cache the credentials for the user-challenge in a way that we can search
	// the Wallet address using the session token, because that session token is
	// provided to us in order to search for the matching Wallet address, which is
	// requesting to continue playing the game. Because of that token based
	// indexing, and because we want to cleanup old references of still expiring
	// tokens, if the same user requests another session token, we need to track
	// an additional index pointing from Wallet address to session token, and use
	// that to search for dangling resources. If such dangling resources exist, we
	// want to remove those. That includes the token references themselves, and
	// the expiration callbacks.

	var wal common.Address
	{
		wal = cli.Wallet()
	}

	{
		old, exi := h.ind.Search(wal)
		if exi {
			h.txp.Delete(old)
			h.tok.Delete(old)
		}
	}

	{
		h.tok.Update(tok, wal)
		h.ind.Update(wal, tok)
	}

	{
		h.txp.Ensure(tok, h.ttl, func() {
			h.tok.Delete(tok)
			h.ind.Delete(wal)
		})
	}

	// Encode the auth response and send the new session token back to the client
	// connection that requested this new credential.

	return cli.Stream(schema.Encode(schema.Auth, tok[:]))
}
