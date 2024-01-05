package controllers

import (
	"encoding/json"
	"exercicios-go/api-rest/database"
	"exercicios-go/api-rest/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	fmt.Println("TodasPersonalidades")
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)

}

func NovaPersonalidade(w http.ResponseWriter, r *http.Request) {

	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}
func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.Delete(&p, id)

	json.NewEncoder(w).Encode(p)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	// err := database.DB.First(&p, id).Error
	err := database.DB.Where("Id = ?", id).First(&p).Error
	if err != nil {
		data := map[string]string{"message": "id não encontrado"}
		w.WriteHeader(404)

		//payload, _ := json.Marshal(data)
		//w.Write(payload)
		json.NewEncoder(w).Encode(data)
		return
	}
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Model(&p).Where("id = ?", id).Updates(models.Personalidade{Nome: p.Nome, Historia: p.Historia})
	json.NewEncoder(w).Encode(p)
}
