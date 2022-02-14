package boot

import (
	"auth-service/internal/core/app/auth"
	"auth-service/internal/driven/db/cockroachdb"
)

var (
	authenticateService  auth.AuthenticateService
	authorizationService auth.AuthorizeService
)

func LoadServices() {
	authrRepo := cockroachdb.NewAuthenticationRepository(client)
	authenticateService = auth.NewAuthenticateService(authrRepo, []byte(accessSecret), []byte(refreshSecret))
	authorizationService = auth.NewAuhtorizeService([]byte(accessSecret), []byte(refreshSecret))
}
