package repository

import (
	"context"

	"github.com/Leaeraso/max_inventory/internal/entity"
)

const (
	qryInsertUser = `
		insert into users (email, name, password)
		values (?, ?, ?);
	`
	qryGetUserByEmail = `
		select 
			id, name, email, password
		from users
		where email = ?;
	`

	qryInsertUserRole = `
		insert into user_roles (user_id, role_id) values (:user_id, :role_id);
	`

	qryRemoveUserRole = `
		delete from user_roles where id = :user_id and role_id = :role_id;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}

	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {

	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)
	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}

	err := r.db.SelectContext(ctx, &roles, "select user_id, role_id from user_roles where user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return roles, nil

}
