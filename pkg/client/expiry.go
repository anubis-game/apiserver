package client

func (c *Client) Expiry() chan struct{} {
	return c.exp
}
