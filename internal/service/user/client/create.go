package client

import "github.com/arslanovdi/omp-bot/internal/model/user"

func (c *DummyClientService) Create(client user.Client) (uint64, error) {
	//TODO implement me

	c.clients = append(c.clients, client)

	return uint64(len(c.clients)), nil
}
