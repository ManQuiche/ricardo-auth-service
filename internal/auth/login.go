package auth

import (
	"auth-service/internal/auth/firebase"
	"auth-service/pkg/errors"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func LoginUser(writer http.ResponseWriter, req *http.Request, p httprouter.Params) {
	reqToken := req.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 {
		errors.InvalidTokenFormat(writer)
		return
	}

	switch splitToken[0] {
	case "Basic":
		LoginUserFromCreds(writer, req)
		break
	case "Bearer":
		LoginUserFromToken(writer, req, splitToken[1])
		break
	}
	reqToken = strings.TrimSpace(splitToken[1])
}

func LoginUserFromToken(writer http.ResponseWriter, req *http.Request, token string) {
	respToken, err := firebase.FireAuth.VerifyIDToken(req.Context(), token)
	if err != nil {
		_ = json.NewEncoder(writer).Encode(err.Error())
	}

	// Generate JWT

	_ = json.NewEncoder(writer).Encode(respToken)
}

func LoginUserFromCreds(writer http.ResponseWriter, req *http.Request) {
	username, password, ok := req.BasicAuth()
	if ok {
		// Check in DB if exists

		// Generate JWT
	} else {
		// Error: Cannot extract credentials from token
	}
}
