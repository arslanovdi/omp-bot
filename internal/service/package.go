package service

import (
	"context"
	"github.com/arslanovdi/omp-bot/internal/model"
	"time"
)

// интерфейс grpc клиента
type Client interface {
	CreatePackage(ctx context.Context, pkg model.Package) (*uint64, error)
	DeletePackage(ctx context.Context, id uint64) error
	GetPackage(ctx context.Context, id uint64) (*model.Package, error)
	ListPackages(ctx context.Context, offset uint64, limit uint64) ([]model.Package, error)
	UpdatePackage(ctx context.Context, cursor uint64, pkg model.Package) error
}

const timeout = time.Second * 5

// LogisticPackageService реализует интерфейс слоя бизнес-логики
type LogisticPackageService struct {
	api Client
}

func NewPackageService(grpc Client) *LogisticPackageService {
	return &LogisticPackageService{
		api: grpc,
	}
}
