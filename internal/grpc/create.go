package grpc

import (
	"context"
	"fmt"
	pb "github.com/arslanovdi/logistic-package-api/pkg/logistic-package-api"
	"github.com/arslanovdi/omp-bot/internal/model"
)

// Create вызывает gRPC функцию CreateV1
func (client *Client) Create(ctx context.Context, pkg model.Package) (*uint64, error) {

	response, err := client.send.CreateV1(
		ctx,
		&pb.CreateRequestV1{
			Value: pkg.ToProto(),
		})

	if err != nil {
		return nil, fmt.Errorf("grpc.Create: %w", err)
	}

	return &response.PackageId, nil
}
