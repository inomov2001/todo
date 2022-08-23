package service

import (
	"context"
	"todo-service/todo"
)

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

type Service struct {
	repo Repository
}

func (s Service) CreateTodo(ctx context.Context, t todo.Todo) (todo.Todo, error) {
	return s.repo.CreateTodo(ctx, t)
}

func (s Service) GetTodosByUserID(ctx context.Context, userID int32) ([]todo.Todo, error) {
	return s.repo.GetTodosByUserID(ctx, userID)
}
