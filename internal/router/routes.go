package routes

import (
	"ASCII-ART-WEB/internal/handlers"

	"github.com/gorilla/mux"
)


func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/ascii-art", handlers.AsciiController)

	return r
}
