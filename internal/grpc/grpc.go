package grpc

import (
	"context"
	"github.com/arslanovdi/omp-bot/internal/config"
	"github.com/arslanovdi/omp-bot/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
)

import pb "github.com/arslanovdi/logistic-package-api/pkg/logistic-package-api"

type grpsClient struct {
	Send pb.LogisticPackageApiServiceClient
	Conn *grpc.ClientConn
}

// CreatePackage вызывает gRPC функцию CreatePackageV1
func (client *grpsClient) CreatePackage(ctx context.Context, pkg model.Package) (*uint64, error) {

	log := slog.With("func", "grpsClient.CreatePackageV1")

	response, err := client.Send.CreatePackageV1(
		ctx,
		&pb.CreatePackageRequestV1{
			Value: pkg.ToProto(),
		})

	if err != nil {
		log.Error("Create package failed:", slog.Any("error", err))
		return nil, err
	}

	log.Debug("Package created", slog.Uint64("id", response.PackageId))

	return &response.PackageId, nil
}

// DeletePackage вызывает gRPC функцию DeletePackageV1
func (client *grpsClient) DeletePackage(ctx context.Context, id uint64) error {

	log := slog.With("func", "grpsClient.DeletePackageV1")

	response, err := client.Send.DeletePackageV1(
		ctx,
		&pb.DeletePackageV1Request{
			PackageId: id,
		})

	if err != nil {
		log.Error("Delete package failed:", slog.Uint64("id", id), slog.Any("error", err))
		return err
	}

	if response.Removed {
		log.Debug("Package deleted", slog.Uint64("id", id))
		return nil
	} else { // false прилетает если package not found
		log.Debug("Package not found", slog.Uint64("id", id))
		return model.NotFound
	}
}

// GetPackage вызывает gRPC функцию GetPackageV1
func (client *grpsClient) GetPackage(ctx context.Context, id uint64) (*model.Package, error) {

	log := slog.With("func", "grpsClient.GetPackageV1")

	response, err := client.Send.GetPackageV1(
		ctx,
		&pb.GetPackageV1Request{
			PackageId: id,
		})

	if err != nil {
		log.Error("Get package failed:", slog.Any("error", err))
		return nil, err
	}

	pkg := model.Package{}
	pkg.FromProto(response.Value)

	log.Debug("Get package", slog.String("package", pkg.String()))

	return &pkg, nil
}

// ListPackages вызывает gRPC функцию ListPackagesV1
func (client *grpsClient) ListPackages(ctx context.Context, offset uint64, limit uint64) ([]model.Package, error) {

	log := slog.With("func", "grpsClient.ListPackagesV1")

	response, err := client.Send.ListPackagesV1(
		ctx,
		&pb.ListPackagesV1Request{
			Offset: offset,
			Limit:  limit,
		})

	if err != nil {
		log.Error("List packages failed:", slog.Any("error", err))
		return nil, err
	}

	packages := make([]model.Package, len(response.Packages))
	for i := 0; i < len(response.Packages); i++ {
		packages[i].FromProto(response.Packages[i])
	}

	log.Debug("List packages", slog.Uint64("offset", offset), slog.Uint64("limit", limit))

	return packages, nil
}

// UpdatePackage вызывает gRPC функцию UpdatePackageV1
func (client *grpsClient) UpdatePackage(ctx context.Context, id uint64, pkg model.Package) error {

	log := slog.With("func", "grpsClient.UpdatePackageV1")

	response, err := client.Send.UpdatePackageV1(
		ctx,
		&pb.UpdatePackageV1Request{
			PackageId: id,
			Value:     pkg.ToProto(),
		})

	if err != nil {
		log.Error("Update package failed:", slog.Uint64("id", id), slog.Any("error", err))
		return err
	}

	if response.Updated {
		log.Debug("Package updated", slog.String("package", pkg.String()))
		return nil
	} else { // false прилетает если package not found
		log.Debug("Package not found", slog.Uint64("id", id))
		return model.NotFound
	}
}

// NewGrpcClient инициализирует соединение с gRPC сервером
func NewGrpcClient() *grpsClient {
	log := slog.With("func", "grpsClient.NewGrpcClient")

	cfg := config.GetConfigInstance()

	// подключение к grpc серверу без TLS
	conn, err := grpc.Dial(
		cfg.GRPC.Host+":"+cfg.GRPC.Port,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Warn("did not connect", slog.Any("error", err))
		os.Exit(1)
	}

	log.Info("gRPC client connected", slog.Any("address", cfg.GRPC.Host+":"+cfg.GRPC.Port))
	return &grpsClient{
		Send: pb.NewLogisticPackageApiServiceClient(conn), // инициализируем интерфейс через который будут вызываться удаленные методы
		Conn: conn,
	}
}

// Close закрывает соединение с gRPC сервером
func (client *grpsClient) Close() {
	log := slog.With("func", "grpsClient.Close")

	client.Conn.Close()

	log.Info("gRPC client disconnected")
}
