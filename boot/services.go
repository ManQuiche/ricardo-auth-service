package boot

import (
	"fmt"
	"gitlab.com/ricardo134/auth-service/internal/core/app/auth"
	"gitlab.com/ricardo134/auth-service/internal/driven/db/cockroachdb"
	"gitlab.com/ricardo134/auth-service/internal/driven/firebase"
	"log"

	"github.com/nats-io/nats.go"
	ricardoNats "gitlab.com/ricardo134/auth-service/internal/driven/broker/nats"
)

var (
	authenticateService  auth.AuthenticateService
	authorizationService auth.AuthorizeService
	externalTokenService auth.ExternalTokenService

	natsEncConn *nats.EncodedConn
)

func LoadServices() {

	natsConn, err := nats.Connect(fmt.Sprintf("nats://%s:%s@%s", natsUsr, natsPwd, natsURL))
	if err != nil {
		log.Fatal(err)
	}
	natsEncConn, err = nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	authrRepo := cockroachdb.NewAuthenticationRepository(client)
	tokenRepo := firebase.NewTokenRepository(firebaseAuth)
	registerNotifier := ricardoNats.NewRegisterNotifier(natsEncConn, natsRegisterTopic)

	authenticateService = auth.NewAuthenticateService(authrRepo, registerNotifier, []byte(accessSecret), []byte(refreshSecret))
	authorizationService = auth.NewAuthorizeService([]byte(accessSecret), []byte(refreshSecret))
	externalTokenService = auth.NewExternalTokenService(tokenRepo, authrRepo, registerNotifier, []byte(accessSecret), []byte(refreshSecret))
}
