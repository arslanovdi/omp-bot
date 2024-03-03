package client

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/user"
)

// List возвращаем клиентов id: cursor, нумерация с 1, количество - limit
func (c *DummyClientService) List(cursor uint64, limit uint64) ([]user.Client, error) {
	//TODO implement me

	if cursor > uint64(len(c.clients)) {
		return nil, fmt.Errorf("выход за границу массива")
	}

	if cursor+limit > uint64(len(c.clients)) {
		return c.clients[cursor-1:], user.EndOfList
	}

	if cursor+limit <= uint64(len(c.clients)) {
		return c.clients[cursor-1 : cursor-1+limit], nil
	} else {
		return c.clients[cursor-1 : cursor-1+limit], user.EndOfList
	}
}
