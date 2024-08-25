package store

import (
	"context"

	"github.com/Shakhrik/mini_app_backend/dto"
	sdto "github.com/Shakhrik/mini_app_backend/dto/store"
)

type Repository struct {
	Product     ProductRepo
	TelegramBot TelegramBotRepo
}

type ProductRepo interface {
	CreateProduct(product dto.Product) (dto.Product, error)
}

type TelegramBotRepo interface {
	Create(ctx context.Context, params sdto.TelegramBotCreateParams) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*sdto.TelegramBot, error)
}
