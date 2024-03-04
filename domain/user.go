package domain

import (
	"context"
	"math/rand"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser(name string) *User {
	return &User{
		Id:   rand.Intn(10000),
		Name: name,
	}
}

// UserUpdate represents the subset of user fields than can be updated using UpdateUser()
type UserUpdate struct {
	Name *string `json:"name"`
}

type UserService interface {
	// TODO: Retreives a user by ID
	// Returns custom error if user does not exist
	FindUserById(ctx context.Context, id int) (*User, error)

	// TODO: Retreive a list of users
	// Returns nil, nil if no users exist
	FindUsers(ctx context.Context) ([]*User, error)

	// TODO: Create a new user
	CreateUser(ctx context.Context, user *User) error

	// TODO: Update a user object.
	// TODO: Returns a custom error if the user in context is not the user being updated
	UpdateUser(ctx context.Context, id int, userUpdate UserUpdate) (*User, error)

	// TODO: Delete a user and all associated keys.
	// TODO: Returns a custom error if the user in context is not the user being deleted.
	DeleteUser(ctx context.Context, id int) error
}
