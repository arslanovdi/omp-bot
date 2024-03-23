package service

import (
	"context"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// Get возвращаем package с id: cursor
func (c *LogisticPackageService) Get(cursor uint64) (model.Package, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	pkg, err := c.api.GetPackage(ctx, cursor)
	if err != nil {
		return model.Package{}, err
	}

	return *pkg, nil
}
