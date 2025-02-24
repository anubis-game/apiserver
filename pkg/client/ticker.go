package client

func (c *Client) Ticker() chan struct{} {
	return c.tic
}
