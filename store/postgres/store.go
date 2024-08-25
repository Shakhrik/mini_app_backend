package postgres

import (
	"github.com/Shakhrik/mini_app_backend/store"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *store.Repository {
	return &store.Repository{
		Product:     newProductRepo(db),
		TelegramBot: newTelegramBotRepo(db),
	}
}
