package home

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request){


	temp,err := template.ParseFiles("views/home/index.html")


	if err != nil {
		
	}

	temp.Execute(w, nil)
}