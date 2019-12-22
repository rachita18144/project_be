package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"project_be/oForm"
)

type OFormResponse struct {
	Success bool `json:"success"`
}

func saveOformDetails(w http.ResponseWriter, r *http.Request) {
	var request oForm.OForm
	var response OFormResponse

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

	if err := oForm.SaveOFormDetailsBoundary(request); err != nil {
		log.Println("ERROR:", err)
		returnResponse(response, http.StatusUnauthorized, w)
		return
	} else {
		log.Println("Incoming Request", request)
		response.Success = true
		returnResponse(response, http.StatusOK, w)
		return
	}

}
