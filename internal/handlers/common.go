package handlers

import "github.com/gorilla/mux"

var router *mux.Router

func InitRouter() *mux.Router{
	router = mux.NewRouter().StrictSlash(true)

	initCardsHandlers()

	return router
}
