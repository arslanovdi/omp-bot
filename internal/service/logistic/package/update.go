package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
)

// Update изменяем имя клиента с id: cursor, нумерация с 1
func (c *DummyPackageService) Update(cursor uint64, pkg logistic.Package) error {
	//TODO implement me

	if cursor > uint64(len(c.packages)) {
		return fmt.Errorf("выход за границу массива")
	}

	c.packages[cursor-1] = pkg

	return nil
}
