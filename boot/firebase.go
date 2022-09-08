package boot

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"gitlab.com/ricardo134/auth-service/pkg/errors"
	option2 "google.golang.org/api/option"
)

var (
	firebaseApp  *firebase.App
	firebaseAuth *auth.Client
)

func InitFirebaseApp() {
	var err error
	option := option2.WithCredentialsFile("boot/ricardo-9b5d5-firebase-adminsdk-udnxf-2e4b3b051f.json")
	firebaseApp, err = firebase.NewApp(context.Background(), nil, option)
	if err != nil {
		errors.CannotInitFirebaseApp()
	}

	firebaseAuth, err = firebaseApp.Auth(context.Background())
	if err != nil {
		errors.CannotInitFirebaseApp()
	}
}
