package postgres

import "github.com/Shakhrik/mini_app_backend/store"

func NewRepository(db string) *store.Repository {
	return &store.Repository{
		Product: newProductRepo(db),
	}
}
