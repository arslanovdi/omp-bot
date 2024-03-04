package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
)

// Get возвращаем package с id: cursor, нумерация с 1
func (c *DummyPackageService) Get(cursor uint64) (logistic.Package, error) {
	//TODO implement me

	if cursor > uint64(len(c.packages)) {
		return logistic.Package{}, fmt.Errorf("выход за границу массива")
	}

	return c.packages[cursor-1], nil
}
