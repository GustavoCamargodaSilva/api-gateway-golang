package repository

import (
	"errors"
)

var (
	ErrorNotFound = errors.New("transaction not found")

	ErrorDuplicateKey = errors.New("duplicate idempotency key")
)
