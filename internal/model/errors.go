package model

import "errors"

var (
	EndOfList   = errors.New("end of list")
	ErrNotFound = errors.New("not found")
)
