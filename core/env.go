package main

import (
	"os"
)

var (
	PORT    string
	PG_URL  string
	PG_PORT string
	PG_USER string
	PG_NAME string
	PG_PASS string
)

func init() {
	// Init env vars
	PORT = getEnvOrElse("PORT", "8080")
	PG_URL = getEnvOrElse("PG_URL", "postgres-service")
	PG_PORT = getEnvOrElse("PG_PORT", "5432")
	PG_USER = getEnvOrElse("PG_USER", "postgres")
	PG_NAME = getEnvOrElse("PG_NAME", "postgres")
	// Ideally for secrets we wouldn't use an "OrElse" version of this function, so that
	// creds aren't stored in the repo
	PG_PASS = getEnvOrElse("PG_PASS", "password")
}

func getEnvOrElse(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
