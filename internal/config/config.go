package config

import (
	"fmt" 
	"os"
)

type Config struct {
	DatabaseURL string
}

func Load() (Config , error) {

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{} , fmt.Errorf("DATABASE_URL is requierd")
	}

	return Config {
		DatabaseURL: databaseURL,
	} , nil
}