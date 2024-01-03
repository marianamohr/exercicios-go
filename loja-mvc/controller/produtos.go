package controller

import (
	"exercicios-go/alura_loja/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAll()

	temp.ExecuteTemplate(w, "index.html", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "new.html", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Fatal("Erro na conversão do preço:", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Fatal("Erro na conversão da quantidade:", err)
		}
		models.Create(nome, descricao, precoConvertido, quantidadeConvertida)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.Delete(id)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	id := r.URL.Query().Get("id")

	produto := models.Edit(id)

	temp.ExecuteTemplate(w, "edit.html", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		fmt.Println(nome)
		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Erro na conversão do ID:", err)
		}
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Fatal("Erro na conversão do preço:", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Fatal("Erro na conversão da quantidade:", err)
		}
		models.Update(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
