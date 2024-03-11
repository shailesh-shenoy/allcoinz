package db

import (
	"context"
	"database/sql"

	"github.com/shailesh-shenoy/allcoinz/domain"
)

type UserService struct {
	ds *DataStore
}

func NewUserService(ds *DataStore) *UserService {
	return &UserService{
		ds: ds,
	}
}

func (userService *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	tx, err := userService.ds.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

}

func (userService *UserService) FindUserById(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

// TODO: Retreive a list of users
// Returns nil, nil if no users exist
func (userService *UserService) FindUsers(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}

// TODO: Update a user object.
// TODO: Returns a custom error if the user in context is not the user being updated
func (userService *UserService) UpdateUser(ctx context.Context, id int, userUpdate domain.UserUpdate) (*domain.User, error) {
	return nil, nil
}

// TODO: Delete a user and all associated keys.
// TODO: Returns a custom error if the user in context is not the user being deleted.
func (userService *UserService) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func createUserInDb(ctx context.Context, tx *sql.Tx, user *domain.User) error {
	result, err := tx.ExecContext(ctx, `
		INSERT INTO users (
			name
		)
		VALUES (?)
	`,
		user.Name,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = int(id)
	return nil
}
