package main

import (
	"github.com/godofprodev/simple-pass/internal/router"
	"github.com/godofprodev/simple-pass/internal/storage"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("there was an issue loading .env")
	}

	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal("there was an issue connecting to the db")
	}

	r := router.New(store)

	r.RegisterMiddlewares()
	r.RegisterHandlers()

	err = r.Listen()
	if err != nil {
		log.Fatal("there was an issue listening to port 8080: ", err)
	}
}
