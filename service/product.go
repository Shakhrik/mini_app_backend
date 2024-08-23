package service

import (
	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/Shakhrik/mini_app_backend/store"
)

type product struct {
	productRepo store.ProductRepo
}

func newProduct(productRepo store.ProductRepo) *product {
	return &product{
		productRepo: productRepo,
	}
}
func (p *product) CreateProduct(product dto.Product) (dto.Product, error) {
	createdProduct, err := p.productRepo.CreateProduct(product)
	if err != nil {
		return dto.Product{}, err
	}
	return createdProduct, nil
}
