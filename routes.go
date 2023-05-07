package main

import (
	"github.com/julienschmidt/httprouter"
)

func (a *App) routes() *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/health", a.Home)
	mux.GET("/snippet", a.ShowSnippet)
	mux.POST("/snippet", a.CreateSnippet)
	mux.POST("/snippets", a.CreateSnippets)
	return mux
}
