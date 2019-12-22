package auth

import (
	"log"

	"project_be/base"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getUserEntity(username, password string) (User, error) {

	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("Error:", err)
		return User{}, err
	}
	defer session.Close()

	c := session.DB(base.DB_PORTAL).C(base.COL_USERS)
	var user User
	query := bson.M{
		"username": username,
		"password": password,
	}
	err = c.Find(query).One(&user)
	if err != nil {
		log.Println("ERROR:", err)
	}
	return user, err
}
