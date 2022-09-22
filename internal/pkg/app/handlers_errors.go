package app

import (
	"errors"
)

var (
	errorMissingID     = errors.New("ID не задан")
	errorMissingAmount = errors.New("сумма не задана")
)
