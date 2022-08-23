package service

import (
	"context"
	"user-service/user"
)

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) RegisterUser(ctx context.Context, name string) (int, error) {
	return s.repo.CreateUser(ctx, name)
}

func (s Service) ListUsers(ctx context.Context) ([]user.User, error) {
	return s.repo.ListUsers(ctx)
}

func (s Service) GetUserByID(ctx context.Context, id int) (user.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

type Repository interface {
	CreateUser(ctx context.Context, name string) (int, error)
	ListUsers(ctx context.Context) ([]user.User, error)
	GetUserByID(ctx context.Context, id int) (user.User, error)
}
