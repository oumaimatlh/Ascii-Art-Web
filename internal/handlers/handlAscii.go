package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)


func AsciiController(w http.ResponseWriter, r *http.Request) {
	var errorMsg string

	if r.Method != http.MethodPost {
		fmt.Println("Testing")
		http.Error(w, "400", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	inputText := r.PostForm.Get("content")

	if strings.TrimSpace(inputText) == "" {
		errorMsg = "Vous devez taper quelque chose..."
	} else {
		for _, r := range inputText {
			if !(r >= 32 && r <= 126 || r == 13 || r == 10) {
				errorMsg = "Input non valide : les caractères doivent être en ASCII 32-126"
				break
			}
		}
	}
	if len(inputText) >= 3000 {
		errorMsg = "Votre Text est dépassée 3000 caractéres"
	}

	font := r.PostForm.Get("types")
	if errorMsg == "" && font == "" {
		errorMsg = "Vous devez choisir un Art"
	}

	var result string
	if errorMsg == "" {
		var err error
		result, err = ApplyingFont(inputText, font)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	data := map[string]string{
		"result": result,
		"error":  errorMsg,
	}

	template, e := template.ParseFiles("web/templates/home.html")
	if e != nil {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	err := template.Execute(w, data)
	if err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

}