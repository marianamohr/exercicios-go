package router

import (
	"ex/blocopad/config"
	"ex/blocopad/internal/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Execute() {
	serverPort := config.ConfigVariables()

	router := mux.NewRouter()
	router.HandleFunc("/api/note/{id}", controller.ReadNote).Methods("GET")
	router.HandleFunc("/api", controller.Check).Methods("GET")
	router.HandleFunc("/api/note", controller.SaveNote).Methods("POST")
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), router)
	fmt.Println("tware", err)
}
