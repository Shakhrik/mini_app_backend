package main

import (
	"log"

	"github.com/Shakhrik/mini_app_backend/api"
	"github.com/Shakhrik/mini_app_backend/service"
	"github.com/Shakhrik/mini_app_backend/store/postgres"
	// "github.com/oapi-codegen/oapi-codegen/v2/examples/minimal-server/gin/api"
)

func main() {
	store := postgres.NewRepository("db-string")

	service := service.NewService(store)

	server := api.NewServer(service)

	// And we serve HTTP until the world ends.
	log.Fatal(server.ListenAndServe())
}
