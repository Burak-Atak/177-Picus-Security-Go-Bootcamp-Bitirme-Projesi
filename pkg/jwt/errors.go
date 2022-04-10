package jwt

import "errors"

var (
	InvalidTokenError = errors.New("invalid token")
	ExpiredTokenError = errors.New("expired token")
	UnAuthorizedError = errors.New("unauthorized")
)
