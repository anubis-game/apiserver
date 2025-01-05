package stream

import (
	"fmt"
)

func (s *Stream) create(add string, cli Client) {
	var exi bool

	{
		s.mut.Lock()
		_, exi = s.cli[add]
		s.cli[add] = cli
		s.mut.Unlock()
	}

	if exi {
		s.log.Log(
			s.ctx,
			"level", "error",
			"message", fmt.Sprintf("Wallet %q added twice", add),
		)
	}
}
