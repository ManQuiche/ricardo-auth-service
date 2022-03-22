package boot

import (
	"ricardo/auth-service/internal/core/app/auth"
	"ricardo/auth-service/internal/driven/db/cockroachdb"
	"ricardo/auth-service/internal/driven/firebase"
)

var (
	authenticateService  auth.AuthenticateService
	authorizationService auth.AuthorizeService
)

func LoadServices() {
	authrRepo := cockroachdb.NewAuthenticationRepository(client)
	authenticateService = auth.NewAuthenticateService(authrRepo, []byte(accessSecret), []byte(refreshSecret))
	authorizationService = auth.NewAuthorizeService([]byte(accessSecret), []byte(refreshSecret))

	if !noFirebase {
		firebase.InitFirebaseSDK()
	}
}
