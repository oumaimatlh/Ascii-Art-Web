package app

import (
	"ASCII-ART-WEB/handlers"
	"fmt"
	"net/http"
)

func App() {

	fmt.Println("Serveur lancé sur : http://localhost:8080")

	
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ascii-art", handlers.AsciiController)


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur:", err)
	}
	
}