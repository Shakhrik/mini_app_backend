package dto

type Product struct {
	Id    int64   `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
