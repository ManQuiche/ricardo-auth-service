package boot

import (
	"auth-service/internal/core/app/auth"
	"auth-service/internal/driven/db/cockroachdb"
)

var (
	authenticateService auth.AuthenticateService
)

func LoadServices() {
	authrRepo := cockroachdb.NewAuthenticationRepository(client)
	authenticateService = auth.NewAuthenticateService(authrRepo)
}
