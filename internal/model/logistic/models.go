package logistic

import (
	"errors"
	"time"
)

var EndOfList = errors.New("end of list")

type Package struct {
	ID        uint64
	Title     string
	CreatedAt time.Time
}

func (c *Package) String() string {
	return c.Title
}
