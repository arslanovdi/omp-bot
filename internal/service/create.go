package service

import (
	"context"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// Create добавляем пакет
func (c *LogisticPackageService) Create(pkg model.Package) (uint64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	id, err := c.api.CreatePackage(ctx, pkg)
	if err != nil {
		return 0, err
	}

	return *id, nil
}
