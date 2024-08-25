package postgres

import (
	"context"

	sdto "github.com/Shakhrik/mini_app_backend/dto/store"
	"github.com/jmoiron/sqlx"
)

type telegramBotRepo struct {
	db *sqlx.DB
}

func newTelegramBotRepo(db *sqlx.DB) *telegramBotRepo {
	return &telegramBotRepo{db: db}
}

func (t *telegramBotRepo) Create(ctx context.Context, params sdto.TelegramBotCreateParams) error {
	query := `INSERT INTO telegram_bot(id, name, token, description, image_url)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := t.db.ExecContext(ctx, query, params.ID, params.Name, params.Token, params.Description, params.ImageUrl)
	return err
}

func (t *telegramBotRepo) Delete(ctx context.Context, id string) error {
	_, err := t.db.ExecContext(ctx, "DELETE FROM telegram_bot WHERE id=$1", id)
	return err
}

func (t *telegramBotRepo) Get(ctx context.Context, id string) (*sdto.TelegramBot, error) {
	var res sdto.TelegramBot
	query := `
	SELECT id, name, token, description, image_url, created_at
	FROM telegram_bot
	WHERE id = $1
	`
	if err := t.db.GetContext(ctx, &res, query); err != nil {
		return nil, err
	}
	return &res, nil
}
