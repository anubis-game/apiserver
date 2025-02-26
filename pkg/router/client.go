package router

type Client struct {
	rac chan<- byte
	tur chan<- Turn
	uid chan<- Uuid
}

func (c *Client) Race() chan<- byte {
	return c.rac
}

func (c *Client) Turn() chan<- Turn {
	return nil
}

func (c *Client) Uuid() chan<- Uuid {
	return c.uid
}
