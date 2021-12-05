package auth

type User struct {
	//
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
