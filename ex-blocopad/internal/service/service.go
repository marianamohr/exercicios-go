package service

import (
	"errors"
	"ex/blocopad/internal/db"
	"fmt"
)

func GetKey(key string) (string, error) {
	// testado
	if len(key) == 0 || len(key) > 36 {
		return "", errors.New("Key with wrong size")
	}
	onetime, data, err := db.GetNote(key)
	fmt.Println(err)
	// testado
	if err != nil {
		return "", err
	}

	if onetime {
		err := db.DeleteNote(key)
		if err != nil {
			panic("Cannot delete onetime note")
		}
	}
	return data, nil
}

func SaveKey(data string, oneTime bool) (string, error) {
	// verifica o tamanho em bites do dado
	// rune retorna um vetor de bytes
	biteSize := len([]rune(data))

	if biteSize == 0 || biteSize > (36*1024) {
		return "", errors.New(("Invalid Note Size"))
	}

	id, err := db.SaveNote(data, oneTime)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return id, nil
}
