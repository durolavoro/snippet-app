package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	mux := httprouter.New()
	mux.GET("/health", Home)
	mux.GET("/snippet", ShowSnippet)
	mux.POST("/snippet", CreateSnippet)

	log.Println("Server initialised on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatalf("Server init failed with err %s", err.Error())
}
