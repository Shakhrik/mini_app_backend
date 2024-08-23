package store

import "github.com/Shakhrik/mini_app_backend/dto"

type Repository struct {
	Product ProductRepo
}

type ProductRepo interface {
	CreateProduct(product dto.Product) (dto.Product, error)
}
