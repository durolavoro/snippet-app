package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
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

func (a *App) ReqError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = a.ErrorLogger.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (a *App) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = a.ErrorLogger.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
