package service

import (
	"context"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// List возвращаем package с позиции offset, количество - limit
func (c *LogisticPackageService) List(offset uint64, limit uint64) ([]model.Package, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	pkg, err := c.api.ListPackages(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return pkg, nil
}
