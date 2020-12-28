package facades

import "errors"

var (
	NotAuthenticatedError     = errors.New("not authenticated")
	AlreadyAuthenticatedError = errors.New("already authenticated")
)
