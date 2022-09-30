package auth

import (
	"context"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
)

type Authenticate interface {
	Login(ctx context.Context, loginRequest entities.LoginRequest) (*tokens.SignedTokens, error)
	Save(ctx context.Context, user entities.User) error
	Refresh(ctx context.Context, token string) (*tokens.SignedTokens, error)
}
