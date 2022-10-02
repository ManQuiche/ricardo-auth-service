package boot

import (
	"fmt"
	"gitlab.com/ricardo134/auth-service/internal/core/app/auth"
	"gitlab.com/ricardo134/auth-service/internal/core/app/user"
	"gitlab.com/ricardo134/auth-service/internal/driven/db/postgresql"
	"gitlab.com/ricardo134/auth-service/internal/driven/firebase"
	"log"

	"github.com/nats-io/nats.go"
	natsext "gitlab.com/ricardo134/auth-service/internal/driven/broker/nats"
)

var (
	authenticateService  auth.AuthenticateService
	authorizationService auth.AuthorizeService
	externalTokenService auth.ExternalTokenService
	userService          user.Service

	natsEncConn *nats.EncodedConn
)

func LoadServices() {
	natsConn, err := nats.Connect(fmt.Sprintf("nats://%s:%s@%s", natsUsr, natsPwd, natsURL))
	if err != nil {
		log.Fatal(err)
	}
	natsEncConn, err = nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	authrRepo := postgresql.NewAuthenticationRepository(client)
	userRepo := postgresql.NewUserRepository(client)
	tokenRepo := firebase.NewTokenRepository(firebaseAuth)
	registerNotifier := natsext.NewUserEventsNotifier(natsEncConn, natsRegisterTopic)

	authenticateService = auth.NewAuthenticateService(authrRepo, registerNotifier, []byte(accessSecret), []byte(refreshSecret))
	authorizationService = auth.NewAuthorizeService([]byte(accessSecret), []byte(refreshSecret))
	externalTokenService = auth.NewExternalTokenService(tokenRepo, authrRepo, registerNotifier, []byte(accessSecret), []byte(refreshSecret))
	userService = user.NewService(userRepo)
}
