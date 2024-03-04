package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
)

// List возвращаем клиентов id: cursor, нумерация с 1, количество - limit
func (c *DummyPackageService) List(cursor uint64, limit uint64) ([]logistic.Package, error) {
	//TODO implement me

	if cursor > uint64(len(c.packages)) {
		return nil, fmt.Errorf("выход за границу массива")
	}

	if cursor+limit > uint64(len(c.packages)) {
		return c.packages[cursor-1:], logistic.EndOfList
	}

	if cursor+limit <= uint64(len(c.packages)) {
		return c.packages[cursor-1 : cursor-1+limit], nil
	} else {
		return c.packages[cursor-1 : cursor-1+limit], logistic.EndOfList
	}
}
