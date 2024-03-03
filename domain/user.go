package domain

import "math/rand"

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
