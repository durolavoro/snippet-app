package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type App struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}

func (a *App) Home(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := SendOKResponse(w, map[string]string{
		"version": "v1",
		"status":  "OK",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error sending response for endpoint /", err.Error())
	}
}

func (a *App) ShowSnippet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	idStr := r.URL.Query().Get("id")
	var err error
	if strings.EqualFold(idStr, "") {
		err = SendOKResponse(w, map[string]string{
			"message": "Display a snippet...",
		})
	} else {
		id, convErr := strconv.Atoi(idStr)
		if convErr != nil {
			http.Error(w,
				fmt.Sprintf("%s - invalid value for param id", http.StatusText(http.StatusBadRequest)),
				http.StatusBadRequest,
			)
			return
		}
		err = SendOKResponse(w, map[string]string{
			"message": fmt.Sprintf("Display a snippet %d", id),
		})
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error sending response for endpoint /showSnippet", err.Error())
	}
}

func (a *App) CreateSnippet(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := SendOKResponse(w, map[string]string{
		"message": "Creating a new snippet...",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error sending response for endpoint /createSnippet", err.Error())
	}
}
