package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"log"
)

var (
	FireApp  *firebase.App
	FireAuth *auth.Client
)

func InitFirebaseSDK() {
	var err error = nil
	FireApp, err = firebase.NewApp(context.Background(), nil)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FireAuth, err = FireApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
}
