package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {

	router := mux.NewRouter()
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()
	apiV1Router.HandleFunc("/login", login).Methods(http.MethodPost)
	apiV1Router.HandleFunc("/save_oform_details", saveOformDetails).Methods(http.MethodPost)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	return router
}
