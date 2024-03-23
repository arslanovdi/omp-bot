package service

import (
	"context"
)

// Delete удаляем пакет с id: cursor
func (c *LogisticPackageService) Delete(cursor uint64) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := c.api.DeletePackage(ctx, cursor)
	if err != nil {
		return false, err
	}

	return true, nil
}
