package client

func (c *Client) Writer() chan struct{} {
	return c.wri
}
