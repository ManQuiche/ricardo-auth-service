package auth

import (
	"context"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
)

type TokenAuthenticate interface {
	Verify(ctx context.Context, token string) (*tokens.SignedTokens, error)
}
