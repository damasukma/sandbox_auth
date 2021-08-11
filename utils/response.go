package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     int                `json:"status"`
	Message    string             `json:"message"`
	Data       *interface{}       `json:"data,omitempty"`
	Validation *map[string]string `json:"validations,omitempty"`
}

func (r *Response) ToJson(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	json.NewEncoder(w).Encode(r)
	return
}
