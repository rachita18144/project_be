package auth

import "log"

func login(username, password string) (User, error) {
	user, err := getUserEntity(username, password)
	if err != nil {
		log.Println("no user found for username:", username, "password:", password)
		return User{}, err
	}

	return user, nil
}
