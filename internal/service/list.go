package service

import (
	"context"
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// List возвращаем packages с позиции offset, количество - limit
func (c *LogisticPackageService) List(offset uint64, limit uint64) ([]model.Package, error) {

	ctx, cancel := context.WithTimeout(context.Background(), c.ctxTimeout)
	defer cancel()

	pkg, err := c.api.ListPackages(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("service.List: %w", err)
	}

	return pkg, nil
}
