package repository

import (
	"context"
	"user-service/user"
)

type DB struct {
	users map[int]user.User
	genID func() int
}

func NewDB() *DB {
	return &DB{
		users: make(map[int]user.User),
		genID: NewIDGenerator(),
	}
}

func (d *DB) CreateUser(ctx context.Context, name string) (int, error) {
	id := d.genID()
	d.users[id] = user.User{id, name}

	return id, nil
}

func (d *DB) ListUsers(ctx context.Context) ([]user.User, error) {
	users := make([]user.User, 0, len(d.users))
	for _, u := range d.users {
		users = append(users, u)
	}

	return users, nil
}

func (d *DB) GetUserByID(ctx context.Context, id int) (user.User, error) {
	return d.users[id], nil
}

func NewIDGenerator() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}
