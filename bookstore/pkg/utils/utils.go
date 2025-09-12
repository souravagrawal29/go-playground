package utils

import (
	"encoding/json"
	"net/http"
	"io"
)


func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func SetErrorResponse(w http.ResponseWriter, err error) {
	errResponse , _ := json.Marshal(err.Error())
	setJsonContentType(w)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(errResponse)
}

func SetSuccessResponse(w http.ResponseWriter, res interface{}) {
	response, _ := json.Marshal(res)
	setJsonContentType(w)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}