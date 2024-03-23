package service

import (
	"context"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// Update изменяем пакет
func (c *LogisticPackageService) Update(cursor uint64, pkg model.Package) error {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := c.api.UpdatePackage(ctx, cursor, pkg)
	if err != nil {
		return err
	}

	return nil
}
