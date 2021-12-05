package errors

import "net/http"

const (
	invalidTokenFormat = "Token format is invalid !"
)

func InvalidTokenFormat(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(invalidTokenFormat))
}
