package _package

import (
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
)

type DummyPackageService struct {
	packages []logistic.Package
}

func NewPackageService() *DummyPackageService {
	return &DummyPackageService{
		packages: []logistic.Package{ // Заглушка 10 клиентов
			{"one"},
			{"two"},
			{"three"},
			{"four"},
			{"five"},
			{"six"},
			{"seven"},
			{"eight"},
			{"nine"},
			{"ten"},
		},
	}
}
