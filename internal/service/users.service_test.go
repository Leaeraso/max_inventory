package service

import (
	"context"
	"os"
	"testing"

	"github.com/Leaeraso/max_inventory/encryption"
	"github.com/Leaeraso/max_inventory/internal/entity"
	"github.com/Leaeraso/max_inventory/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	u := &entity.User{Email: "test@exists.com", Password: encryptedPassword}

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(u, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("GetUserRoles", mock.Anything, int64(1)).Return([]entity.UserRole{{UserID: 1, RoleID: 1}}, nil)
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUSer(t *testing.T) {
	testCases := []struct {
		Name     string
		Email    string
		UserName string
		Password string
		Expected error
	}{
		{
			Name:     "RegisterUSer_Success",
			Email:    "test@test.com",
			UserName: "test",
			Password: "validPassword",
			Expected: nil,
		},
		{
			Name:     "RegisterUser_UserAlreadyExists",
			Email:    "test@exists.com",
			UserName: "test",
			Password: "validPassword",
			Expected: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)
			if err != tc.Expected {
				t.Errorf("expected error %v, got %v", ErrUserAlreadyExists, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name     string
		Email    string
		Password string
		Expected error
	}{
		{
			Name:     "LoginUser_Success",
			Email:    "test@exists.com",
			Password: "validPassword",
			Expected: nil,
		},
		{
			Name:     "LoginUser_InvalidPassword",
			Email:    "test@exists.com",
			Password: "invalidPassword",
			Expected: ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(
			tc.Name,
			func(t *testing.T) {
				t.Parallel()

				repo.Mock.Test(t)

				_, err := s.LoginUser(ctx, tc.Email, tc.Password)
				if err != tc.Expected {
					t.Errorf("expected error %v, got %v", tc.Expected, err)
				}
			},
		)
	}
}

func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		Name     string
		UserID   int64
		RoleID   int64
		Expected error
	}{
		{
			Name:     "AddUserRole_Success",
			UserID:   1,
			RoleID:   2,
			Expected: nil,
		},
		{
			Name:     "AddUserRoel_UserAlreadyHaveRole",
			UserID:   1,
			RoleID:   1,
			Expected: ErrRoleAlreadyAdded,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.AddUserRole(ctx, tc.UserID, tc.RoleID)
			if err != tc.Expected {
				t.Errorf("expected error %v, got %v", tc.Expected, err)
			}
		})
	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		Name     string
		UserID   int64
		RoleID   int64
		Expected error
	}{
		{
			Name:     "RemoveUserRole_Success",
			UserID:   1,
			RoleID:   1,
			Expected: nil,
		},
		{
			Name:     "RemoveUserRole_UserDoesNotHaveRole",
			UserID:   1,
			RoleID:   3,
			Expected: ErrRoleNotFound,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RemoveUserRole(ctx, tc.UserID, tc.RoleID)
			if err != tc.Expected {
				t.Errorf("expected error %v, got %v", tc.Expected, err)
			}
		})
	}

}
