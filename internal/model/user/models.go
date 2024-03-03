package user

import "errors"

var EndOfList = errors.New("end of list")

type Client struct {
	Name string
}

func (c *Client) String() string {
	return c.Name
}
