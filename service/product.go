package service

import (
	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/Shakhrik/mini_app_backend/store"
)

type product struct {
	productStore store.ProductRepo
}

func newProduct(productRepo store.ProductRepo) *product {
	return &product{
		productStore: productRepo,
	}
}
func (p *product) CreateProduct(product dto.Product) (dto.Product, error) {
	createdProduct, err := p.productStore.CreateProduct(product)
	if err != nil {
		return dto.Product{}, err
	}
	return createdProduct, nil
}
