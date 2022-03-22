package auth

import "context"

type Authorize interface {
	AccessAuthorize(ctx context.Context, accessToken string) (bool, error)
	RefreshAuthorize(ctx context.Context, refreshToken string) (bool, error)
}
