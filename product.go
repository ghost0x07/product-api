package product

import "context"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func New() *Product {
	return &Product{}
}

type Store interface {
	FindAll(ctx context.Context) ([]*Product, error)
	FindById(ctx context.Context, id int) (*Product, error)
	Create(ctx context.Context, p *Product) (*Product, error)
	Update(ctx context.Context, id int, p *Product) (*Product, error)
	Delete(ctx context.Context, id int) (*Product, error)
}
