package entities

import "github.com/golang-jwt/jwt"

type TokenPair struct {
	Access  jwt.Token `json:"access_token"`
	Refresh jwt.Token `json:"refresh_token"`
}

type SignedTokenPair struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
}
