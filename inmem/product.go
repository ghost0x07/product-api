package inmem

import (
	"context"
	"fmt"
	"product-api"

	"github.com/google/uuid"
)

type ProductStore struct {
	data []*product.Product
}

var _ product.Store = (*ProductStore)(nil)

func NewProductStore() *ProductStore {
	return &ProductStore{
		data: make([]*product.Product, 0),
	}
}

func (ps *ProductStore) FindAll(ctx context.Context) ([]*product.Product, error) {
	return ps.data, nil
}

func (ps *ProductStore) FindById(ctx context.Context, id int) (*product.Product, error) {
	for _, v := range ps.data {
		if v.ID == id {
			return v, nil
		}
	}
	return &product.Product{}, nil
}

func (ps *ProductStore) Create(ctx context.Context, p *product.Product) (*product.Product, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("unable to generate uuid: %w", err)
	}
	p.ID = int(u.ID())

	ps.data = append(ps.data, p)

	return p, nil
}

func (ps *ProductStore) Update(ctx context.Context, id int, p *product.Product) (*product.Product, error) {
	p.ID = id
	for i, v := range ps.data {
		if v.ID == id {
			ps.data[i] = p
			return p, nil
		}
	}
	return &product.Product{}, nil
}

func (ps *ProductStore) Delete(ctx context.Context, id int) (*product.Product, error) {
	for i, v := range ps.data {
		if v.ID == id {
			ps.data = append(ps.data[:i], ps.data[i+1:]...)
			return v, nil
		}
	}

	return &product.Product{}, nil
}
