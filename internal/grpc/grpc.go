package grpc

import (
	"context"
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/config"
	"github.com/arslanovdi/omp-bot/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
)

import pb "github.com/arslanovdi/logistic-package-api/pkg/logistic-package-api"

type grpcClient struct {
	send pb.LogisticPackageApiServiceClient
	conn *grpc.ClientConn
}

// CreatePackage вызывает gRPC функцию CreatePackageV1
func (client *grpcClient) CreatePackage(ctx context.Context, pkg model.Package) (*uint64, error) {

	response, err := client.send.CreatePackageV1(
		ctx,
		&pb.CreatePackageRequestV1{
			Value: pkg.ToProto(),
		})

	if err != nil {
		return nil, fmt.Errorf("grpc.CreatePackage: %w", err)
	}

	return &response.PackageId, nil
}

// DeletePackage вызывает gRPC функцию DeletePackageV1
func (client *grpcClient) DeletePackage(ctx context.Context, id uint64) (bool, error) {

	response, err := client.send.DeletePackageV1(
		ctx,
		&pb.DeletePackageV1Request{
			PackageId: id,
		})

	if err != nil {
		return false, fmt.Errorf("grpc.DeletePackage: %w", err)
	}

	return response.Deleted, nil
}

// GetPackage вызывает gRPC функцию GetPackageV1
func (client *grpcClient) GetPackage(ctx context.Context, id uint64) (*model.Package, error) {

	response, err := client.send.GetPackageV1(
		ctx,
		&pb.GetPackageV1Request{
			PackageId: id,
		})

	if err != nil {
		return nil, fmt.Errorf("grpc.GetPackageV1: %w", err)
	}

	pkg := model.Package{}
	pkg.FromProto(response.Value)

	return &pkg, nil
}

// ListPackages вызывает gRPC функцию ListPackagesV1
func (client *grpcClient) ListPackages(ctx context.Context, offset uint64, limit uint64) ([]model.Package, error) {

	response, err1 := client.send.ListPackagesV1(
		ctx,
		&pb.ListPackagesV1Request{
			Offset: offset,
			Limit:  limit,
		})

	if err1 != nil {
		status, ok := status.FromError(err1)
		if !ok {
			return nil, fmt.Errorf("grpc.ListPackages: %w", err1)
		}
		if status.Code() == codes.NotFound {
			return nil, model.EndOfList
		}
		return nil, fmt.Errorf("grpc.ListPackages: %w", err1)
	}

	packages := make([]model.Package, len(response.Packages))
	for i := 0; i < len(response.Packages); i++ {
		packages[i].FromProto(response.Packages[i])
	}

	return packages, nil
}

// UpdatePackage вызывает gRPC функцию UpdatePackageV1
func (client *grpcClient) UpdatePackage(ctx context.Context, id uint64, pkg model.Package) (bool, error) {

	response, err := client.send.UpdatePackageV1(
		ctx,
		&pb.UpdatePackageV1Request{
			PackageId: id,
			Value:     pkg.ToProto(),
		})

	if err != nil {
		return false, fmt.Errorf("grpc.UpdatePackage: %w", err)
	}

	return response.Updated, nil
}

// NewGrpcClient инициализирует соединение с gRPC сервером
func NewGrpcClient() *grpcClient {

	log := slog.With("func", "grpcClient.NewGrpcClient")

	cfg := config.GetConfigInstance()

	// подключение к grpc серверу без TLS
	conn, err := grpc.Dial(
		cfg.GRPC.Host+":"+cfg.GRPC.Port,
		grpc.WithBlock(), // ожидание подключения
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Warn("did not connect", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("gRPC client connected", slog.Any("address", cfg.GRPC.Host+":"+cfg.GRPC.Port))
	return &grpcClient{
		send: pb.NewLogisticPackageApiServiceClient(conn), // инициализируем интерфейс через который будут вызываться удаленные методы
		conn: conn,
	}
}

// Close закрывает соединение с gRPC сервером
func (client *grpcClient) Close() {
	log := slog.With("func", "grpcClient.Close")

	client.conn.Close()

	log.Info("gRPC client disconnected")
}
