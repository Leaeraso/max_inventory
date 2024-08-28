package service

import (
	"context"
	"testing"

	"github.com/Leaeraso/max_inventory/internal/models"
)

func TestAddProduct(t *testing.T) {
	testCases := []struct {
		Name string
		Product models.Product
		Email string
		Expected error
	}{
		{
			Name: "AddProduct_Success",
			Product: models.Product{
				Name: "Test Product",
				Description: "Test Description",
				Price: 10.00,
			},
			Email: "admin@email.com",
			Expected: nil,
		},
		{
			Name: "AddProduct_IvalidPermissions",
			Product: models.Product{
				Name: "Test Product",
				Description: "Test Description",
				Price: 10.00,
			},
			Email: "customer@email.com",
			Expected: ErrInvalidPermission,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.AddProduct(ctx, tc.Product, tc.Email)
			if err != tc.Expected {
				t.Errorf("expected error %v, got %v", tc.Expected, err)
			}
		})
	}
}