package app

import (
	routes "ASCII-ART-WEB/internal/router"
	"fmt"
	"net/http"
)

func App() {

	r := routes.Router()
	fmt.Println("Serveur lancé sur : http://localhost:8080")


	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Erreur:", err)
	}
	
}