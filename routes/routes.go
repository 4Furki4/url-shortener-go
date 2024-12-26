package routes

import (
	"url-shortener/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", controllers.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortURL}", controllers.RedirectURL).Methods("GET")
	return router
}
