package registry

import (
	"github.com/ethereum/go-ethereum/common"
)

func (r *Registry) Address() common.Address {
	return r.add
}
