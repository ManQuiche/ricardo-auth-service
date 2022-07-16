package boot

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"ricardo/auth-service/pkg/errors"
)

var (
	firebaseApp  *firebase.App
	firebaseAuth *auth.Client
)

func InitFirebaseApp() {
	var err error
	firebaseApp, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		errors.CannotInitFirebaseApp()
	}

	firebaseAuth, err = firebaseApp.Auth(context.Background())
	if err != nil {
		errors.CannotInitFirebaseApp()
	}
}
