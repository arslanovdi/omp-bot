package _package

import "github.com/arslanovdi/omp-bot/internal/model/logistic"

func (c *DummyPackageService) Create(pkg logistic.Package) (uint64, error) {
	//TODO implement me

	c.packages = append(c.packages, pkg)

	return uint64(len(c.packages)), nil
}
