package service

import (
	"context"
	"database/sql"
	"fmt"

	srdto "github.com/Shakhrik/mini_app_backend/dto/service"
	stdto "github.com/Shakhrik/mini_app_backend/dto/store"
	lerrors "github.com/Shakhrik/mini_app_backend/pkg/errors"

	"github.com/Shakhrik/mini_app_backend/store"
	"github.com/google/uuid"
)

type telegramBot struct {
	telegramBotStore store.TelegramBotRepo
}

func newTelegramBot(tgBotrepo store.TelegramBotRepo) *telegramBot {
	return &telegramBot{
		telegramBotStore: tgBotrepo,
	}
}

func (t *telegramBot) Create(ctx context.Context, params *srdto.TelegramBotCreateParams) (string, error) {
	id := uuid.New().String()
	p := stdto.TelegramBotCreateParams{
		ID:          id,
		Name:        params.Name,
		Token:       params.Token,
		Description: params.Description,
		ImageUrl:    params.ImageUrl,
	}
	if err := t.telegramBotStore.Create(ctx, p); err != nil {
		return "", fmt.Errorf("failed to create telegram bot: %w", err)
	}
	return id, nil
}

func (t *telegramBot) Delete(ctx context.Context, id string) error {
	err := t.telegramBotStore.Delete(ctx, id)
	if err == sql.ErrNoRows {
		return lerrors.ErrNotFound
	}
	return err
}

func (t *telegramBot) Get(ctx context.Context, id string) (*stdto.TelegramBot, error) {
	res, err := t.telegramBotStore.Get(ctx, id)
	if err == sql.ErrNoRows {
		return nil, lerrors.ErrNotFound
	}
	return res, err
}
