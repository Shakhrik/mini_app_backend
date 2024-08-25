package main

import (
	"fmt"

	"github.com/Shakhrik/mini_app_backend/api"
	"github.com/Shakhrik/mini_app_backend/config"
	"github.com/Shakhrik/mini_app_backend/pkg/logger"
	"github.com/Shakhrik/mini_app_backend/service"
	"github.com/Shakhrik/mini_app_backend/store/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "user_service")
	defer logger.Cleanup(log)

	log.Info("main: pgxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		log.Error("Error while connecting database: %v", logger.Error(err))
		return
	}
	defer connDb.Close()

	store := postgres.NewRepository(connDb)

	service := service.NewService(store)

	server := api.NewServer(service)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error while listening: %v", logger.Error(err))
	}
}
