package postgres

import (
	"sync"

	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/jmoiron/sqlx"
)

var products = make(map[int64]dto.Product)
var productIDCounter int64 = 1
var mu sync.Mutex

type productRepo struct {
	db *sqlx.DB
}

func newProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo{db: db}
}

func (p *productRepo) CreateProduct(product dto.Product) (dto.Product, error) {
	mu.Lock()
	defer mu.Unlock()

	product.Id = productIDCounter
	products[productIDCounter] = product
	productIDCounter++

	return product, nil
}
