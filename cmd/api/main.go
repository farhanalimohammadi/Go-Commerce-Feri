package main

import (
	"context"
	"log"

	"github.com/farhanalimohammadi/Go-Commerce-Feri/internal/config"
	"github.com/farhanalimohammadi/Go-Commerce-Feri/internal/database"
)

func main() {

	cfg , err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db , err := database.Open(context.Background() , cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string

	err = db.QueryRowContext(
		context.Background(),
		"SELECT version()",
	).Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(version)

	log.Println("database connection established")
}
