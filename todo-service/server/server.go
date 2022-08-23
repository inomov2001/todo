package server

import (
	"context"
	"todo-service/service"
	"todo-service/todo"
	"todo-service/todopb"
)

func NewGRPCServer(svc service.Service) Server {
	return Server{
		svc: svc,
	}
}

type Server struct {
	svc service.Service
	todopb.UnimplementedTodoServiceServer
}

func (s Server) CreateTodo(ctx context.Context, req *todopb.CreateTodoRequest) (*todopb.Todo, error) {
	t, err := s.svc.CreateTodo(ctx, todo.Todo{
		Body:   req.GetBody(),
		UserID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}

	return toProtoTodo(t), nil
}

func (s Server) GetTodosByUserID(ctx context.Context, req *todopb.GetTodosByUserIDRequest) (*todopb.TodoList, error) {
	todos, err := s.svc.GetTodosByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return toProtoTodoList(todos), nil
}

func toProtoTodo(t todo.Todo) *todopb.Todo {
	return &todopb.Todo{
		Id:     t.ID,
		Body:   t.Body,
		Done:   t.Done,
		UserId: t.UserID,
	}
}

func toProtoTodoList(todos []todo.Todo) *todopb.TodoList {
	protoTodos := make([]*todopb.Todo, 0, len(todos))
	for _, t := range todos {
		protoTodos = append(protoTodos, toProtoTodo(t))
	}

	return &todopb.TodoList{Todos: protoTodos}
}
