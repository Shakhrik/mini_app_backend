package store

import (
	"sync"

	"github.com/Shakhrik/mini_app_backend/dto"
)

var products = make(map[int64]dto.Product)
var productIDCounter int64 = 1
var mu sync.Mutex

func SaveProduct(product dto.Product) (dto.Product, error) {
	mu.Lock()
	defer mu.Unlock()

	product.Id = productIDCounter
	products[productIDCounter] = product
	productIDCounter++

	return product, nil
}
