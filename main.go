package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	infoLogger := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &App{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}

	mux := httprouter.New()
	mux.GET("/health", app.Home)
	mux.GET("/snippet", app.ShowSnippet)
	mux.POST("/snippet", app.CreateSnippet)

	app.InfoLogger.Println("Server initialised on :4000")
	err := http.ListenAndServe(":4000", mux)
	app.ErrorLogger.Fatalf("Server init failed with err %s", err.Error())
}
