package grpc

import (
	"context"
	"fmt"
	pb "github.com/arslanovdi/logistic-package-api/pkg/logistic-package-api"
	"github.com/arslanovdi/omp-bot/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

// Get вызывает gRPC функцию GetV1
func (client *grpcClient) Get(ctx context.Context, id uint64) (*model.Package, error) {

	log := slog.With("func", "grpcClient.Get")

	response, err := client.send.GetV1(
		ctx,
		&pb.GetV1Request{
			PackageId: id,
		})

	if err != nil {
		log.Error("fail to get package", slog.String("error", err.Error()))
		if status.Code(err) == codes.NotFound {
			return nil, model.ErrNotFound
		}

		return nil, fmt.Errorf("fail to get package error: %s", status.Code(err).String())
	}

	pkg := model.Package{}
	pkg.FromProto(response.Value)

	return &pkg, nil
}
