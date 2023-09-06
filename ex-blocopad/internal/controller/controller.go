package controller

import (
	"encoding/json"
	"ex/blocopad/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Note struct {
	Text    string `json:"data"`
	OneTime bool   `json:"onetime"`
}

func writeResponse(status int, body interface{}, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(body)
	w.Write(payload)
}

func ReadNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	text, err := service.GetKey(id)
	if err != nil {
		data := map[string]string{"message": "id não encontrado"}
		writeResponse(404, data, w)
		return
	}
	data := map[string]string{"data": text}
	writeResponse(202, data, w)
	return
}

func SaveNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		writeResponse(400, map[string]string{"error": err.Error()}, w)
		return
	}
	text, err := service.SaveKey(note.Text, note.OneTime)

	data := map[string]string{"id": text}
	writeResponse(200, data, w)
}

func Check(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"message": "Okkk"}
	writeResponse(200, data, w)
}
