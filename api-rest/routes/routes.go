package routes

import (
	"exercicios-go/api-rest/controllers"
	"exercicios-go/api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.NovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
