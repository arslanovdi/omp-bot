package _package

import "github.com/arslanovdi/omp-bot/internal/model/logistic"

func (c *DummyPackageService) Describe(PackageID uint64) (*logistic.Package, error) {
	//TODO implement me
	return &logistic.Package{
		Title: "stub",
	}, nil
}
