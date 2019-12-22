package main

import (
	"log"
	"net/http"
	"project_be/api"
	"project_be/base"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	serverDetails, err := base.SetupServer(base.SERVER_TYPE_LOCALHOST)
	if err != nil {
		log.Println("ERROR SETTING UP SERVER")
		return
	}
	log.Println(serverDetails)
	//go maiPareshanKarungiSabko()
	http.ListenAndServe(base.PORT, api.GetRouter())
}

/*
func maiPareshanKarungiSabko() {

	// Mera nam rachita hai
	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("Error:", err)
	}
	defer session.Close()

	c := session.DB(base.DB_PORTAL).C(base.COL_USERS)
	type Temp struct {
		Hello string `bson:"hello" json:"hello"`
	}
	var temp Temp
	temp.Hello = nil
	err = c.Insert(temp)
	if err != nil {
		log.Println("ERROR:", err)
	}
	log.Println("Kya hua?", err)
}
*/
