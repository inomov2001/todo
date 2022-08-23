package service

import (
	"context"
	"todo-service/todo"
)

type Repository interface {
	CreateTodo(ctx context.Context, todo todo.Todo) (todo.Todo, error)
	GetTodosByUserID(ctx context.Context, userID int32) ([]todo.Todo, error)
}
