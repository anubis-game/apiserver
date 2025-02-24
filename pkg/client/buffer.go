package client

func (c *Client) Buffer() chan []byte {
	return c.buf
}
