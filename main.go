package main

import (
	"go-native/config"
	"go-native/controllers/category"
	"go-native/controllers/home"
	"log"
	"net/http"
)

func main(){

	config.ConnectDB()


	//home
	http.HandleFunc("/", home.Welcome)

	//Categoris
	http.HandleFunc("/categories", category.Index)
	http.HandleFunc("/categories/add", category.Add)
	http.HandleFunc("/categories/edit", category.Edit)
	http.HandleFunc("/categories/delete", category.Delete)

	log.Println("Server running on port 8080")

	http.ListenAndServe(":8080",nil)

}