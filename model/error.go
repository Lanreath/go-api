package model

import "errors"

var (
	ErrNotFound = errors.New("no rows in result set")
)
