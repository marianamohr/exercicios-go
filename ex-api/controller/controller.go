package controller

import (
	"encoding/json"
	"ex-api/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Note struct {
	Note string `json:"note"`
}

func SaveNote(w http.ResponseWriter, r *http.Request) {
	// cria uma variavel de Note
	var note Note
	fmt.Println("Controller")
	// faz a leitura  do body, criando um decoder, e atribuindo o valor
	// a note, atraves do ponteiro &

	err := json.NewDecoder(r.Body).Decode(&note)

	defer r.Body.Close()

	// testa se a conversão deu erro
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	text, err := db.SaveNote(note.Note)
	fmt.Println(text, err)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(text)
	w.Write(payload)
}

func ReadNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(vars)
	text, err := db.GetNote(id)
	fmt.Println("Errorrr", err)
	if err != nil {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"message": "Id não encontrado"}
		payload, _ := json.Marshal(data)
		w.Write(payload)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{"data": text}
	payload, _ := json.Marshal(data)
	w.Write(payload)
}
