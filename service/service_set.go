package service

import (
	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/Shakhrik/mini_app_backend/store"
)

type Service struct {
	Product ProductI
}

func NewService(repo *store.Repository) *Service {
	return &Service{
		Product: newProduct(repo.Product),
	}
}

type ProductI interface {
	CreateProduct(product dto.Product) (dto.Product, error)
}
