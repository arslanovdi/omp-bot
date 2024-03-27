package grpc

import (
	"github.com/arslanovdi/omp-bot/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
)

import pb "github.com/arslanovdi/logistic-package-api/pkg/logistic-package-api"

type grpcClient struct {
	send pb.LogisticPackageApiServiceClient
	conn *grpc.ClientConn
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
