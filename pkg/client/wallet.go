package client

import "github.com/ethereum/go-ethereum/common"

func (c *Client) Wallet() common.Address {
	return c.wal
}
