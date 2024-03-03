package client

import "github.com/arslanovdi/omp-bot/internal/model/user"

func (c *DummyClientService) Describe(ClientID uint64) (*user.Client, error) {
	//TODO implement me
	return &user.Client{
		Name: "stub",
	}, nil
}
