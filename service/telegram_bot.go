package service

import (
	"context"
	"fmt"

	srdto "github.com/Shakhrik/mini_app_backend/dto/service"
	stdto "github.com/Shakhrik/mini_app_backend/dto/store"

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
