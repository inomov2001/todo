package service

import (
	"api-gateway/entity"
	"context"
)

func New(userServiceURL, todoServiceURL string) Service {
	return Service{
		User: NewUserServiceGRPCClient(userServiceURL),
		Todo: NewTodoServiceGRPCClient(todoServiceURL),
	}
}

type Service struct {
	User UserServiceClient
	Todo TodoServiceClient
}

type UserServiceClient interface {
	GetByID(ctx context.Context, id int) (entity.User, error)
	Register(ctx context.Context, name string) (entity.User, error)
}

type TodoServiceClient interface {
	CreateTodo(ctx context.Context, todo entity.Todo) (entity.Todo, error)
	GetTodosByUserID(ctx context.Context, userID int) ([]entity.Todo, error)
}
