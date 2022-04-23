package errors

import "net/http"

const (
	ErrInvalidToken            = "INVALID_TOKEN"
	ErrInvalidTokenCode        = http.StatusUnauthorized
	ErrInvalidTokenDescription = "token is invalid"

	ErrCannotFindUserDescription = "cannot find user"
)
