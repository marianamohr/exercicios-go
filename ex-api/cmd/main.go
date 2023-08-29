package main

import (
	"ex-api/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serverPort := "8080"

	router := mux.NewRouter()
	router.HandleFunc("/api/save", controller.SaveNote).Methods("POST")
	router.HandleFunc("/api/read/{id}", controller.ReadNote).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), router)
	fmt.Println("Iniciando o servidor Rest com Go")
	fmt.Println(err)

}
