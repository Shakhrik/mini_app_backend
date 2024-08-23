package main

import (
	"log"
	"net/http"

	"github.com/Shakhrik/mini_app_backend/api"
	"github.com/Shakhrik/mini_app_backend/api/handler"
	"github.com/Shakhrik/mini_app_backend/service"
	"github.com/Shakhrik/mini_app_backend/store/postgres"

	// "github.com/oapi-codegen/oapi-codegen/v2/examples/minimal-server/gin/api"
	"github.com/gin-gonic/gin"
)

func main() {
	store := postgres.NewRepository("db-string")

	service := service.NewService(store)
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := handler.NewHandler(service)

	r := gin.Default()
	api.RegisterHandlers(r, server)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
