package handler

import (
	"api-gateway/entity"
	"api-gateway/service"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func New(svc service.Service) Handler {
	return Handler{
		service: svc,
	}
}

type Handler struct {
	service service.Service
}

func (h Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var u entity.User
	if err := render.DecodeJSON(r.Body, &u); err != nil {
		panic(err)
	}

	user, err := h.service.User.Register(context.Background(), u.Name)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, user)
}

func (h Handler) GetUserTodos(w http.ResponseWriter, r *http.Request) {
	userID, err := extractUserIDParam(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, render.M{
			"error": "path parameter 'id' is not a number",
		})
		return
	}

	// STEP 1 - get user data
	user, err := h.service.User.GetByID(context.Background(), userID)
	if err != nil {
		panic(err)
	}

	// STEP 2 - get user's todos
	todos, err := h.service.Todo.GetTodosByUserID(context.Background(), userID)
	if err != nil {
		panic(err)
	}

	// STEP 3 - compose (add) data from services to single response
	render.JSON(w, r, GetUserTodosResponse{
		User:  user,
		Todos: todos,
	})
}

func (h Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var request entity.Todo
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		panic(err)
	}

	todo, err := h.service.Todo.CreateTodo(context.Background(), request)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, todo)
}

func extractUserIDParam(r *http.Request) (int, error) {
	userIDParam := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
