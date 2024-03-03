package client

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/user"
)

// Get возвращаем клиента с id: cursor, нумерация с 1
func (c *DummyClientService) Get(cursor uint64) (user.Client, error) {
	//TODO implement me

	if cursor > uint64(len(c.clients)) {
		return user.Client{}, fmt.Errorf("выход за границу массива")
	}

	return c.clients[cursor-1], nil
}
