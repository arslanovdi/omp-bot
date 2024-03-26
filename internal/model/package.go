package model

import (
	"errors"
	"fmt"
	pb "github.com/arslanovdi/logistic-package-api/pkg/logistic-package-api"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

var EndOfList = errors.New("end of list")

type Package struct {
	ID        uint64
	Title     string
	Weight    uint64
	CreatedAt time.Time
}

func (c *Package) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Weight: %d, CreatedAt: %s", c.ID, c.Title, c.Weight, c.CreatedAt)
}

func (c *Package) ToProto() *pb.Package {
	return &pb.Package{
		Id:     c.ID,
		Title:  c.Title,
		Weight: &c.Weight,
		Created: &timestamp.Timestamp{
			Seconds: c.CreatedAt.Unix(),
			Nanos:   int32(c.CreatedAt.Nanosecond()),
		},
	}
}

func (c *Package) FromProto(pkg *pb.Package) {
	c.ID = pkg.Id
	c.Title = pkg.Title
	c.Weight = *pkg.Weight
	c.CreatedAt = time.Unix(pkg.Created.Seconds, int64(pkg.Created.Nanos))
}
