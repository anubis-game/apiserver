package engine

import (
	"github.com/anubis-game/apiserver/pkg/address"
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/worker/action"
	"github.com/anubis-game/apiserver/pkg/worker/release"
	"github.com/anubis-game/apiserver/pkg/worker/resolve"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

// TODO:game implement Engine.Kill
//
//     Should this even be a client input driven reconciliation?
//

func (e *Engine) Kill(_ byte, _ *client.Client, inp []byte) error {
	var err error

	{
		// e.ply.act[u] = false
	}

	//
	//     inp[0]        action
	//     inp[1:21]     winner
	//     inp[21:41]    loser
	//

	var act uuid.UUID
	{
		act, err = uuid.NewRandom()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var kil uuid.UUID
	{
		kil, err = uuid.NewRandom()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var win common.Address
	var los common.Address
	{
		win = common.BytesToAddress(inp[1:21])
		los = common.BytesToAddress(inp[21:41])
	}

	// If the winner is the zero address, then the loser got killed by the
	// environment, and must therefore be released from the Registry without
	// benefiting any other player.
	//
	// If the winner is not the zero address, then the winner killed the loser,
	// so the given kill state must be resolved to benefit the winner.

	if address.Empty(win) {
		e.wrk.Ensure(action.New(release.Action{
			Act: act,
			Kil: kil,
			Los: los,
		}))
	} else {
		e.wrk.Ensure(action.New(resolve.Action{
			Act: act,
			Kil: kil,
			Los: los,
			Win: win,
		}))
	}

	// Fan out the kill response to all participating users.

	// var out []byte
	// {
	// 	out = schema.Encode(schema.Kill, win.Bytes(), los.Bytes())
	// }

	// e.cli.Ranger(func(_ common.Address, val *client.Client) {
	// 	val.Stream(out)
	// })

	{
		// Since the loser got killed, they must be removed from the current game.
		// Regardless, we can only remove the loser from the broadcasting worker
		// pool, once they received their own kill signal over the broadcast in the
		// step above. So only once the kill state got communicated to everyone,
		// including the losing player themselves, only then can we remove the loser
		// from the broadcasting worker pool. E.g. e.cli.Delete(los).
	}

	{
		// Player UIDs are not deleted when the client connection is lost. Only
		// losing the game causes UIDs to be freed. E.g. e.uni.Delete(wal).
	}

	{
		// We map the player's wallet address to their respective session tokens,
		// because we want to allow active players to quickly reconnect if they
		// experience intermittent network issues. This means the engine must clear
		// a player's session token if a player leaves the game. E.g.
		// e.tkx.Delete(wal)
	}

	return nil
}
