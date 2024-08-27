package service

import (
	"context"
	"errors"

	"github.com/Leaeraso/max_inventory/internal/models"
)

var validRolesToAddProducts []int64 = []int64{1, 2}
var ErrInvalidPermission error = errors.New("user do not have permission to add product")


func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {
	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	// GetProducts me devuelve un array con structs del tipo entity, por lo que vamos a mapearlas para que seam del tipo models
	products := []models.Product{}

	for _, p := range pp {
		products = append(products, models.Product{
			ID: p.ID,
			Name: p.Name,
			Description: p.Description,
			Price: p.Price,
		})
	}
	
	return products, nil
}

func (s *serv) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	product := &models.Product{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
	}

	return product, nil
}

func (s *serv) AddProduct(ctx context.Context, product models.Product, email string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	
	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}

	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProducts {
			if vr == r.RoleID {
				userCanAdd = true
				break
			}
		}
	}

	if !userCanAdd {
		return ErrInvalidPermission
	}

	return s.repo.SaveProduct(ctx, product.Name, product.Description, product.Price, u.ID)
}