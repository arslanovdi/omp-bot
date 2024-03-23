package service

import (
	"context"
	"github.com/arslanovdi/omp-bot/internal/config"
	"github.com/arslanovdi/omp-bot/internal/model"
	"time"
)

// Client интерфейс grpc клиента
type Client interface {
	CreatePackage(ctx context.Context, pkg model.Package) (*uint64, error)
	DeletePackage(ctx context.Context, id uint64) (bool, error)
	GetPackage(ctx context.Context, id uint64) (*model.Package, error)
	ListPackages(ctx context.Context, offset uint64, limit uint64) ([]model.Package, error)
	UpdatePackage(ctx context.Context, cursor uint64, pkg model.Package) (bool, error)
}

// LogisticPackageService слой бизнес-логики
type LogisticPackageService struct {
	api        Client        // gRPC клиент
	ctxTimeout time.Duration // Таймаут контекста gRPC запросов
}

func NewPackageService(grpc Client) *LogisticPackageService {
	cfg := config.GetConfigInstance()
	srv := &LogisticPackageService{
		api:        grpc,
		ctxTimeout: cfg.GRPC.CtxTimeout,
	}
	return srv
}
