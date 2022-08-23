package service

import (
	"api-gateway/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type userServiceHTTPClient struct {
	url    string
	client *http.Client
}

func (a userServiceHTTPClient) GetUserByID(ctx context.Context, id int) (u entity.User, err error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/user/%d", a.url, id), nil)
	if err != nil {
		return entity.User{}, err
	}
	response, err := a.client.Do(request)
	if err != nil {
		return entity.User{}, err
	}

	if err = json.NewDecoder(response.Body).Decode(&u); err != nil {
		return entity.User{}, err
	}

	return u, nil
}

func (a userServiceHTTPClient) RegisterUser(ctx context.Context, name string) (entity.User, error) {
	panic("not implemented")
}

type todoServiceHTTPClient struct {
	url    string
	client *http.Client
}

func (a todoServiceHTTPClient) GetTodosByUserID(ctx context.Context, userID int) ([]entity.Todo, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/todos/%d", a.url, userID), nil)
	if err != nil {
		return nil, err
	}

	response, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}

	todos := []entity.Todo{}
	if err = json.NewDecoder(response.Body).Decode(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}
