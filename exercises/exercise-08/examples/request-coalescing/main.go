package main

import (
	"echo-server/api"
	"echo-server/database"
	"echo-server/database/message"
	"log"
	"net/http"
)

func main() {
	config := database.PsqlConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Database: "postgres",
	}

	messageRepo, err := message.NewPsqlRepository(config)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(messageRepo)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
