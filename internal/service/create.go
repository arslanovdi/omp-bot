package service

import (
	"context"
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// Create добавляем пакет
func (c *LogisticPackageService) Create(pkg model.Package) (uint64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), c.ctxTimeout)
	defer cancel()

	id, err := c.api.CreatePackage(ctx, pkg)
	if err != nil {
		return 0, fmt.Errorf("service.Create: %w", err)
	}

	return *id, nil
}
