package utils

import "errors"

var (
	ErrNoPath         = errors.New("path required")
	ErrUnknownCommand = errors.New("unknown command")
)
