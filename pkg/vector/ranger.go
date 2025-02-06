package vector

import "github.com/anubis-game/apiserver/pkg/object"

func (v *Vector) Ranger(fnc func(object.Object)) {
	cur := v.tai
	for cur != nil {
		fnc(cur.val)
		cur = cur.nxt
	}
}
