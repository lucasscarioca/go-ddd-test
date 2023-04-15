package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	status int
	Data   any `json:"data"`
}

func NewResponse(status int, data any) *Response {
	return &Response{status, data}
}

func (r *Response) WriteJSON(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.status)

	err := json.NewEncoder(w).Encode(r.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Failed to encode JSON")
	}
}
