package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type domainHandler func(http.ResponseWriter, *http.Request) error

func NewHandler(f domainHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func Send(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}
