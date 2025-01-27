package engine

import (
	"github.com/anubis-game/apiserver/pkg/router"
)

func (e *Engine) delete(pac router.Packet) {
	{
		delete(e.mem.ply, pac.Uid)
	}

	// TODO find the bucket index and remove the deleted player from that lookup
	// table
	// {
	// 	e.lkp.ply.Delete(matrix.Bucket)
	// }
}
