package service

import (
	"context"
	"errors"

	"github.com/Leaeraso/max_inventory/encryption"
	"github.com/Leaeraso/max_inventory/internal/models"
)

var (
	ErrUserAlreadyExists error = errors.New("user already exists")
	ErrInvalidCredentials error = errors.New("invalid credentials")
)
func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)
	
	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID: u.ID,
		Email: u.Email,
		Name: u.Name,
	}, nil
}