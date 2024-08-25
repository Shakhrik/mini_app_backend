package service

import (
	"context"

	"github.com/Shakhrik/mini_app_backend/dto"
	srdto "github.com/Shakhrik/mini_app_backend/dto/service"
	stdto "github.com/Shakhrik/mini_app_backend/dto/store"
	"github.com/Shakhrik/mini_app_backend/store"
)

type Service struct {
	Product     ProductI
	TelegramBot TelegramBotI
}

func NewService(repo *store.Repository) *Service {
	return &Service{
		Product:     newProduct(repo.Product),
		TelegramBot: newTelegramBot(repo.TelegramBot),
	}
}

type ProductI interface {
	CreateProduct(product dto.Product) (dto.Product, error)
}

type TelegramBotI interface {
	Create(ctx context.Context, params *srdto.TelegramBotCreateParams) (string, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*stdto.TelegramBot, error)
}
