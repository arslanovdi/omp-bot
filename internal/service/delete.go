package service

import (
	"context"
	"fmt"
)

// Delete удаляем пакет с id: cursor
func (c *LogisticPackageService) Delete(cursor uint64) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), c.ctxTimeout)
	defer cancel()

	removed, err := c.api.DeletePackage(ctx, cursor)
	if err != nil {
		return false, fmt.Errorf("service.Delete: %w", err)
	}

	return removed, nil
}
