package client

import (
	"fmt"
	"slices"
)

// Remove удаляем клиента с id: cursor, нумерация с 1
func (c *DummyClientService) Remove(cursor uint64) (bool, error) {
	//TODO implement me

	if cursor > uint64(len(c.clients)) {
		return false, fmt.Errorf("выход за границу массива")
	}

	c.clients = slices.Delete(c.clients, int(cursor-1), int(cursor))

	return true, nil
}
