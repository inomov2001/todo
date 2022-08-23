package service

import (
	"api-gateway/entity"
	"api-gateway/todopb"
	"api-gateway/userpb"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceGRPCClient(url string) userServiceGRPCClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	client := userpb.NewUserServiceClient(conn)

	return userServiceGRPCClient{
		client: client,
	}
}

type userServiceGRPCClient struct {
	client userpb.UserServiceClient
}

func (c userServiceGRPCClient) Register(ctx context.Context, name string) (entity.User, error) {
	resp, err := c.client.RegisterUser(ctx, &userpb.RegisterUserRequest{
		Name: name,
	})
	if err != nil {
		return entity.User{}, err
	}

	return toEntityUser(resp), nil
}

func (c userServiceGRPCClient) GetByID(ctx context.Context, id int) (entity.User, error) {
	resp, err := c.client.GetUserByID(ctx, &userpb.GetUserByIDRequest{Id: int32(id)})
	if err != nil {
		return entity.User{}, err
	}

	return toEntityUser(resp), nil
}

func toEntityUser(u *userpb.User) entity.User {
	return entity.User{
		ID:   int(u.GetId()),
		Name: u.GetName(),
	}
}

func NewTodoServiceGRPCClient(url string) todoServiceGRPCClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := todopb.NewTodoServiceClient(conn)

	return todoServiceGRPCClient{
		client: client,
	}
}

type todoServiceGRPCClient struct {
	client todopb.TodoServiceClient
}

func (c todoServiceGRPCClient) CreateTodo(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	resp, err := c.client.CreateTodo(ctx, &todopb.CreateTodoRequest{
		Body:   todo.Body,
		UserId: int32(todo.UserID),
	})
	if err != nil {
		return entity.Todo{}, err
	}

	return toEntityTodo(resp), nil
}

func (c todoServiceGRPCClient) GetTodosByUserID(ctx context.Context, userID int) ([]entity.Todo, error) {
	resp, err := c.client.GetTodosByUserID(ctx, &todopb.GetTodosByUserIDRequest{UserId: int32(userID)})
	if err != nil {
		return nil, err
	}

	return toEntityTodoSlice(resp.Todos), nil
}

func toEntityTodo(protoTodo *todopb.Todo) entity.Todo {
	return entity.Todo{
		ID:     int(protoTodo.GetId()),
		Body:   protoTodo.GetBody(),
		Done:   protoTodo.GetDone(),
		UserID: int(protoTodo.GetUserId()),
	}
}

func toEntityTodoSlice(protoTodos []*todopb.Todo) []entity.Todo {
	todos := make([]entity.Todo, 0, len(protoTodos))
	for _, pt := range protoTodos {
		todos = append(todos, toEntityTodo(pt))
	}

	return todos
}
