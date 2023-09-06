package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

type Retorno struct {
	Uuid string ` json:"id"`
}

type Data struct {
	Data string ` json:"data"`
}

func TestCreateSuccess(t *testing.T) {
	postURL := "http://localhost:8080/api/note"
	jsonData := []byte(`{"data": "Mari", "onetime": false}`)
	req, error := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		panic(error)
	}

	var retorno Retorno

	json.NewDecoder(response.Body).Decode(&retorno)

	defer response.Body.Close()

	if response.StatusCode != 200 {
		t.Fatal("TestCreateSuccess falhou, retornou outro status", response.Status)
	}

	if reflect.TypeOf(retorno.Uuid).Kind() != reflect.String {
		t.Fatal("TestCreateSuccess falhou, retornou outro tipo de dados no body", reflect.TypeOf(retorno.Uuid))
	}
}

func TestReadWithSuccess(t *testing.T) {
	postURL := "http://localhost:8080/api/note"
	jsonData := []byte(`{"data": "Estudar Go", "onetime": false}`)
	req, error := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		panic(error)
	}

	var retorno Retorno

	json.NewDecoder(response.Body).Decode(&retorno)

	defer response.Body.Close()

	var data Data
	getURL := "http://localhost:8080/api/note/" + retorno.Uuid

	request, err := http.NewRequest("GET", getURL, nil)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client = &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(error)
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if data.Data != "Estudar Go" {
		t.Fatal("TestReadWithSuccess => Body não veio com o dado correto. Esperado: 'Estudar Go'. Recebido ", data.Data)
	}
}
