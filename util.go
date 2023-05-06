package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var errMarshalling = errors.New("error marshalling body json")

func SendOKResponse(w http.ResponseWriter, body interface{}) error {
	p, err := json.Marshal(body)
	if err != nil {
		return errMarshalling
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(p)
	return nil
}
