package unit_tests

import (
	"errors"
	"ex/blocopad/internal/db"
	"ex/blocopad/internal/service"
	"strings"
	"testing"
)

var (
	deleteInvoked bool
	deletedKey    string
)

func TestGetKeyOk(t *testing.T) {
	// mock da função GetNote de db
	// o retorno deve ser exatamento o que é esperado para o case de teste
	// Given
	db.GetNote = func(key string) (bool, string, error) {
		return false, "OK", nil
	}

	data, err := service.GetKey("key1")
	if err != nil {
		t.Fatal("TestGetKeyOk Should not return error")
	}
	if data != "OK" {
		t.Fatal("TestGetKeyOk Invalid return")
	}
}

func TestGetErrorSize(t *testing.T) {
	keywrongA := ""
	keywrongB := strings.Repeat("a", 40)

	_, err := service.GetKey(keywrongA)
	if err == nil {
		t.Fatal("TestGetErrorSize Chave error deveria ter vindo nil")
	}
	// When
	_, err = service.GetKey(keywrongB)

	// Then
	if err == nil {
		t.Fatal("TestGetErrorSize Deveria ter retornado erro de key com mais de 36 char")
	}
}
func TestGetKeyDbError(t *testing.T) {
	// fazendo o mock  da função GetNote do db
	// onde retorna erro
	db.GetNote = func(key string) (bool, string, error) {
		return false, "OK", errors.New("Error")
	}

	_, err := service.GetKey("key1")

	if err == nil {
		t.Fatal("TestGetKeyDbError Should return error")
	}

}

func TestGetKeyDeleteOk(t *testing.T) {
	// Arrange
	db.GetNote = func(key string) (bool, string, error) {
		return true, "Ok", nil
	}
	// Mock que a note foi deletada com sucesso
	db.DeleteNote = func(key string) error {
		deleteInvoked = true
		deletedKey = key
		return nil
	}
	data, err := service.GetKey("key1")

	if err != nil {
		t.Fatal("TestGetKeyDeleteOk não deveria retornar nil")
	}

	if data != "Ok" {
		t.Fatal("TestGetKeyDeleteOk retorno invalido do data")
	}
	if !deleteInvoked {
		t.Fatal("TestGetKeyDeleteOk Deveria ter invocado a função db.DeleteNote")
	}
	if deletedKey != "key1" {
		t.Fatal("TestGetKeyDeleteOk função delete deveria ter sido chamada com key1")
	}
}

func TestGetKeyDeleteDbError(t *testing.T) {
	// função anonima que testa se o panic foi chamado

	defer func() {
		// a função recover() retorna
		r := recover()
		if r == nil {
			t.Errorf("TestGetKeyDeleteDbError não deu panic")
		}
	}()

	db.GetNote = func(key string) (bool, string, error) {
		return true, "Ok", nil
	}

	db.DeleteNote = func(key string) error {
		deleteInvoked = true
		deletedKey = key
		return errors.New("error")
	}
	service.GetKey("key1")

}

func TestInvalidSize(t *testing.T) {
	noteZeroLength := ""
	noteTooBig := strings.Repeat("a", 50000)

	_, err := service.SaveKey(noteZeroLength, false)

	if err == nil {
		t.Fatal("TestInvalidSize deveria retornar um erro por nota vazia")
	}

	_, err = service.SaveKey(noteTooBig, false)

	if err == nil {
		t.Fatal("TestInvalidSize deveria retornar um erro nota muito grande")
	}
}

func TestSaveKeyOK(t *testing.T) {
	noteOk := "Mari"
	oneTime := true

	db.SaveNote = func(data string, oneTime bool) (string, error) {
		return "10", nil
	}
	id, err := service.SaveKey(noteOk, oneTime)

	if err != nil {
		t.Fatal("TestSaveKeyOK deveria ter vindo sem erro")
	}

	if id != "10" {
		t.Fatal("TestSaveKeyOK deveria ter o Id 10")
	}
}
func TestSaveKeyDbError(t *testing.T) {
	noteOk := "Mari"
	oneTime := true

	db.SaveNote = func(data string, oneTime bool) (string, error) {
		return "10", errors.New(("Error"))
	}

	_, err := service.SaveKey(noteOk, oneTime)

	if err == nil {
		t.Fatal("TestSaveKeyDbError deveria ter vindo um erro do db")
	}
}
