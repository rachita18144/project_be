package auth

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(username, password string) (User, error) {
	return login(username, password)
}
