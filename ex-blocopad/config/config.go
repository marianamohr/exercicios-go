package config

import (
	"ex/blocopad/internal/db"
	"os"
)

func ConfigVariables() string {
	serverPort := "8080"
	if port, hasValue := os.LookupEnv("API_PORT"); hasValue {
		serverPort = port
	}
	databaseUrl := "localhost:6379"
	if dbUrl, hasValue := os.LookupEnv("API_DB_URL"); hasValue {
		databaseUrl = dbUrl
	}
	databasePassword := ""
	if dbPassword, hasValue := os.LookupEnv("API_DB_PASSWORD"); hasValue {
		databasePassword = dbPassword
	}

	db.DatabaseUrl = databaseUrl
	db.DatabasePassword = databasePassword
	return serverPort
}
