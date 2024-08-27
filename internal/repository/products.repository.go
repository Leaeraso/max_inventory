package repository

import (
	"context"

	"github.com/Leaeraso/max_inventory/internal/entity"
)

const (
	qryInsertProduct = `
		insert into products (name, description, price, created_by) values (?, ?, ?, ?);
	`

	qryGetAllProducts = `
		select 
			id,
			name,
			description,
			price,
			created_by
		from products;
	`

	qrygetProductById = `
		select
			id,
			name,
			description,
			price,
			created_by
		from products
		where id = ?;
	`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertProduct, name, description, price, createdBy)
	if err != nil {
		return err
	}
	
	return nil
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}

	err := r.db.SelectContext(ctx, &pp,qryGetAllProducts)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.db.GetContext(ctx, p, qrygetProductById, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}