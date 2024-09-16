package env

import (
	"log"
	"os"
)

var (
	POSTGRES_DB       string
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_HOST     string
	POSTGRES_PORT     string

	API_URL string
)

func Load() {
	POSTGRES_DB = LoadVariable("POSTGRES_DB", RequiredTag)
	POSTGRES_PASSWORD = LoadVariable("POSTGRES_PASSWORD", RequiredTag)
	POSTGRES_USER = LoadVariable("POSTGRES_USER", RequiredTag)
	POSTGRES_HOST = LoadVariable("POSTGRES_HOST", RequiredTag)
	POSTGRES_PORT = LoadVariable("POSTGRES_PORT", RequiredTag)

	API_URL = LoadVariable("API_URL", RequiredTag)
}

const (
	Production  string = "prod"
	Development string = "dev"

	RequiredTag string = "required"
)

func Required(env string) bool {
	return env != ""
}

func LoadVariable(variable string, tags ...string) string {
	value := os.Getenv(variable)

	for i := range tags {
		switch tags[i] {
		case RequiredTag:
			if !Required(value) {
				log.Fatal("Missing ", variable)
			}
		}
	}

	return value
}
