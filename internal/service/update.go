package service

import (
	"context"
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// Update изменяем существующий пакет
func (c *LogisticPackageService) Update(cursor uint64, pkg model.Package) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), c.ctxTimeout)
	defer cancel()

	ok, err := c.api.UpdatePackage(ctx, cursor, pkg)
	if err != nil {
		return false, fmt.Errorf("service.Update: %w", err)
	}

	return ok, nil
}
