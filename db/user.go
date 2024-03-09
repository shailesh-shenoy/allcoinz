package db

import (
	"context"
	"fmt"

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
	if err := userService.ds.db.PingContext(ctx); err != nil {
		return err
	}
	fmt.Println("Create user called in db")
	fmt.Printf("Context: %+v", ctx)
	return nil
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
