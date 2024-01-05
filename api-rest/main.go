package main

import (
	"exercicios-go/api-rest/database"
	"exercicios-go/api-rest/models"
	"exercicios-go/api-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico alemão"},
		{Id: 2, Nome: "Isaac Newton", Historia: "Físico inglês"},
		{Id: 3, Nome: "Stephen Hawking", Historia: "Físico inglês"},
	}

	database.Connect()

	fmt.Println("Go Web App Started on Port 8080")
	routes.HandleRequests()
}
