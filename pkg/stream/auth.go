package stream

import (
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

func (s *Stream) auth(con *websocket.Conn, wal common.Address) error {
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
		exi := s.tok.Exists(tok)
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

	{
		old, exi := s.ind.Search(wal)
		if exi {
			s.txp.Delete(old)
			s.tok.Delete(old)
		}
	}

	{
		s.tok.Update(tok, wal)
		s.ind.Update(wal, tok)
	}

	{
		s.txp.Ensure(tok, func() {
			s.tok.Delete(tok)
			s.ind.Delete(wal)
		})
	}

	// Encode the auth response and send the new session token back to the client
	// connection that requested this new credential.

	var byt []byte
	{
		byt = schema.Encode(schema.Auth, []byte(tok.String()))
	}

	{
		err = con.Write(s.ctx, websocket.MessageBinary, byt)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
