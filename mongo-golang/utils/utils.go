package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	json.NewDecoder(r.Body).Decode(x)
}


func SetNotFoundErrorResponse(w http.ResponseWriter, err error) {
	SetErrorResponse(w, err, http.StatusNotFound)
}

func SetErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	var errResponse []byte
	if err != nil {
		errResponse, _ = json.Marshal(err.Error())
	}
	setJsonContentType(w)
	w.WriteHeader(statusCode)
	w.Write(errResponse)
}

func SetSuccessResponse(w http.ResponseWriter, statusCode int, res interface{}) {
	setJsonContentType(w)
	w.WriteHeader(statusCode)
	if res != nil {
		response, _ := json.Marshal(res)
		w.Write(response)
	}
}


func setJsonContentType(w http.ResponseWriter) {
	setContentType(w, "application/json")
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}