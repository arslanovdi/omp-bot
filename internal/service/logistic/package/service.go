package _package

import (
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
	"time"
)

type DummyPackageService struct {
	packages []logistic.Package
}

func NewPackageService() *DummyPackageService {
	return &DummyPackageService{
		packages: []logistic.Package{ // Заглушка 10 клиентов
			{
				ID:        1,
				Title:     "one",
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Title:     "two",
				CreatedAt: time.Now(),
			},
			{
				ID:        3,
				Title:     "three",
				CreatedAt: time.Now(),
			},
			{
				ID:        4,
				Title:     "four",
				CreatedAt: time.Now(),
			},
			{
				ID:        5,
				Title:     "five",
				CreatedAt: time.Now(),
			},
			{
				ID:        6,
				Title:     "six",
				CreatedAt: time.Now(),
			},
			{
				ID:        7,
				Title:     "seven",
				CreatedAt: time.Now(),
			},
			{
				ID:        8,
				Title:     "eight",
				CreatedAt: time.Now(),
			},
			{
				ID:        9,
				Title:     "nine",
				CreatedAt: time.Now(),
			},
			{
				ID:        10,
				Title:     "ten",
				CreatedAt: time.Now(),
			},
		},
	}
}
