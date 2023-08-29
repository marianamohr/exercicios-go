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

type Note struct {
	Text string `json:"data"`
}

var GetNote = func(key string) (string, error) {
	db := GetDatabase()
	jsonNote, err := db.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", errors.New("not found")
	} else if err != nil {
		// Some other error
		return "", err
	}

	var note Note
	err = json.Unmarshal([]byte(jsonNote), &note)
	if err != nil {
		return "", err
	}

	return note.Text, nil
}

var SaveNote = func(data string) (string, error) {
	fmt.Println(data)
	stringUuid := (uuid.New()).String()
	db := GetDatabase()
	var note Note
	note.Text = data
	jsonNote, err := json.Marshal(note)
	if err != nil {
		return "", err
	}
	exp := 24 * time.Hour
	err = db.SetEx(ctx, stringUuid, jsonNote, exp).Err()
	if err != nil {
		return "", err
	}

	return stringUuid, nil
}

var DeleteNote = func(key string) error {
	db := GetDatabase()
	_, err := db.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
