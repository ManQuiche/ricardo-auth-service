package boot

import (
	"fmt"
	"log"
	"ricardo/auth-service/internal/core/app/auth"
	"ricardo/auth-service/internal/driven/db/cockroachdb"
	"ricardo/auth-service/internal/driven/firebase"

	"github.com/nats-io/nats.go"
	ricardoNats "ricardo/auth-service/internal/driven/broker/nats"
)

var (
	authenticateService  auth.AuthenticateService
	authorizationService auth.AuthorizeService

	natsEncConn *nats.EncodedConn
)

func LoadServices() {

	natsConn, err := nats.Connect(fmt.Sprintf("nats://%s:%s@%s", natsUsr, natsPwd, natsURL))
	if err != nil {
		log.Fatal(err)
	}
	natsEncConn, err = nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	authrRepo := cockroachdb.NewAuthenticationRepository(client)
	registerNotifier := ricardoNats.NewRegisterNotifier(natsEncConn, natsRegisterTopic)
	authenticateService = auth.NewAuthenticateService(authrRepo, registerNotifier, []byte(accessSecret), []byte(refreshSecret))
	authorizationService = auth.NewAuthorizeService([]byte(accessSecret), []byte(refreshSecret))

	if !noFirebase {
		firebase.InitFirebaseSDK()
	}
}
