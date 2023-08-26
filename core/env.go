package main

import (
	"os"
)

var (
	PORT    string
	PG_URL  string
	PG_PORT string
)

func init() {
	// Init env vars
	PORT = getEnvOrElse("PORT", "8080")
	PG_URL = getEnvOrElse("PG_URL", "postgres-service")
	PG_PORT = getEnvOrElse("PG_PORT", "5432")
}

func getEnvOrElse(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
