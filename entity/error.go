package entity

import (
	"errors"

	"github.com/NooBeeID/bee-guard/infra/contracts"
)

var (
	ErrInvalidPassword   = contracts.NewErrBadRequest(errors.New("invalid password"))
	ErrAuthAlreadyExists = contracts.NewErrBadRequest(errors.New("auth already exists"))
)
