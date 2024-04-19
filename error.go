package requests

import "errors"

var (
	ErrIllegalUrl       = errors.New("illegal url")
	ErrMethodNotAllowed = errors.New("method not allowed")
)
