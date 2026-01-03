package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	
	template, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	er := template.Execute(w, nil)
    if er != nil {
        http.Error(w, "Erreur lors de l'affichage du template", http.StatusInternalServerError)
    }

}
