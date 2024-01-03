package models

import (
	"exercicios-go/alura_loja/db"
	"fmt"
	"log"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAll() []Produto {
	db := db.ConectaDb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		log.Fatal(err)
	}

	p := Produto{}
	produtos := []Produto{}
	for rows.Next() {

		err := rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			log.Fatal(err)
		}

		produtos = append(produtos, p)
	}

	return produtos

}

func Create(nome string, descricao string, preco float64, quantidade int) {

	db := db.ConectaDb()
	defer db.Close()

	insereDados, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	insereDados.Exec(nome, descricao, preco, quantidade)

}

func Delete(id string) {

	db := db.ConectaDb()
	defer db.Close()

	deletaDados, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}

	deletaDados.Exec(id)

}

func Edit(id string) Produto {

	db := db.ConectaDb()
	defer db.Close()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	produto := Produto{}
	for produtoDoBanco.Next() {

		err := produtoDoBanco.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			log.Fatal(err)
		}
	}

	return produto

}

func Update(id int, nome string, descricao string, precoConvertido float64, quantidadeConvertida int) {
	db := db.ConectaDb()
	defer db.Close()
	fmt.Println(id, nome, descricao, precoConvertido, quantidadeConvertida)
	atualizaDados, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		log.Fatal(err)
	}

	atualizaDados.Exec(nome, descricao, precoConvertido, quantidadeConvertida, id)
}
