package client

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/user"
)

// Update изменяем имя клиента с id: cursor, нумерация с 1
func (c *DummyClientService) Update(cursor uint64, client user.Client) error {
	//TODO implement me

	if cursor > uint64(len(c.clients)) {
		return fmt.Errorf("выход за границу массива")
	}

	c.clients[cursor-1] = client

	return nil
}
