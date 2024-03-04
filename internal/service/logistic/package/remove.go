package _package

import (
	"fmt"
	"slices"
)

// Remove удаляем клиента с id: cursor, нумерация с 1
func (c *DummyPackageService) Remove(cursor uint64) (bool, error) {
	//TODO implement me

	if cursor > uint64(len(c.packages)) {
		return false, fmt.Errorf("выход за границу массива")
	}

	c.packages = slices.Delete(c.packages, int(cursor-1), int(cursor))

	return true, nil
}
