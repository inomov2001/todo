package handler

import "api-gateway/entity"

type GetUserTodosResponse struct {
	User  entity.User   `json:"user"`
	Todos []entity.Todo `json:"todos"`
}
