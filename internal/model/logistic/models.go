package logistic

import "errors"

var EndOfList = errors.New("end of list")

type Package struct {
	Name string
}

func (c *Package) String() string {
	return c.Name
}
