package auth

type Token struct {
	User  string
	Roles []string
	Data  string
}

type TokenPair struct {
	Access  Token
	Refresh Token
}

type UserClaims
