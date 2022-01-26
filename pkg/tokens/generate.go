package tokens

import "github.com/golang-jwt/jwt"

func GenerateHS256SignedToken(claims jwt.Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
