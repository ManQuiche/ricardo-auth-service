package errors

import "log"

const (
	cannotInitFirebaseApp = "cannot init firebase app, exiting..."
)

func CannotInitFirebaseApp() {
	log.Fatal(cannotInitFirebaseApp)
}
