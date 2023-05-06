package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	infoLogger := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	a := &App{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}

	mux := a.routes()

	a.InfoLogger.Println("Server initialised on :4000")
	err := http.ListenAndServe(":4000", mux)
	a.ErrorLogger.Fatalf("Server init failed with err %s", err.Error())
}
