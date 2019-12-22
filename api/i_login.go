package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"project_be/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	UserId  string `json:"user_id"`
	Success bool   `json:"success"`
}

const (
	DATA_LIMIT = 25 * 1024 * 1024 // limiting request size
)

func login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	var response LoginResponse

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, DATA_LIMIT))
	if err != nil {
		log.Println("Error: Data exceeds limit", err)
		returnResponse(response, http.StatusBadRequest, w)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("ERROR:", err)
		returnResponse(response, http.StatusBadRequest, w)
		return
	}

	if user, err := auth.Login(request.Username, request.Password); err != nil {
		log.Println("ERROR:", err)
		returnResponse(response, http.StatusUnauthorized, w)
		return
	} else {
		response.Success = true
		response.UserId = user.Username
		returnResponse(response, http.StatusOK, w)
		return
	}

}

func returnResponse(response interface{}, statusCode int, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error Authenticating Request: ", err)
	}
}
