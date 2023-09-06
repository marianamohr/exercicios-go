package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	DatabaseUrl      string
	DatabasePassword string
	rDB              *redis.Client
	ctx              = context.Background()
)

type Note struct {
	Text    string `json:"data"`
	OneTime bool   `json:"onetime"`
}

// função que cria a conexão com o redis
var GetDatabase = func() *redis.Client {
	if rDB == nil {
		rDB = redis.NewClient(&redis.Options{
			Addr:     DatabaseUrl,
			Password: DatabasePassword,
			DB:       0,
		})
	}
	return rDB
}

var GetNote = func(key string) (bool, string, error) {
	db := GetDatabase()
	jsonNote, err := db.Get(ctx, key).Result()
	pessoa, err := db.HGetAll(ctx, "pessoa").Result()
	fmt.Println(fmt.Sprintf("%#v", pessoa))
	if err == redis.Nil {
		return false, "", errors.New("not Found")
	} else if err != nil {
		return false, "", err
	}
	var note Note
	err = json.Unmarshal([]byte(jsonNote), &note)

	if err != nil {
		return false, "", err
	}

	return note.OneTime, note.Text, nil
}

var SaveNote = func(data string, oneTime bool) (string, error) {
	id := (uuid.New()).String()
	db := GetDatabase()
	var note Note
	note.Text = data
	note.OneTime = oneTime
	jsonNote, err := json.Marshal(note)
	if err != nil {
		return "", err
	}

	exp := 24 * time.Hour

	err = db.SetEx(ctx, id, jsonNote, exp).Err()
	if err != nil {
		return "", err
	}
	return id, nil
}

var DeleteNote = func(key string) error {
	db := GetDatabase()
	_, err := db.Del(ctx, key).Result()

	if err != nil {
		return err
	}
	return nil
}
