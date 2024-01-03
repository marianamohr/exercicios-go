package main

import (
	"exercicios-go/alura_loja/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
