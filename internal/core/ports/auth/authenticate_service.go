package auth

import "auth-service/internal/core/entities/auth"

type Authenticate interface {
	Login(username, password string) (*auth.TokenPair, error)
}
