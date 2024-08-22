package logic

import (
	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/Shakhrik/mini_app_backend/store"
)

func CreateProduct(product dto.Product) (dto.Product, error) {
	createdProduct, err := store.SaveProduct(product)
	if err != nil {
		return dto.Product{}, err
	}
	return createdProduct, nil
}
