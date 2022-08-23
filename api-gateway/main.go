package main

import (
	"api-gateway/handler"
	"api-gateway/service"
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	userServiceHTTPURL = "http://localhost:8000"
	userServiceGRPCURL = "localhost:9000"
	todoServiceURL     = "http://localhost:8001"
	todoServiceGRPCURL = "localhost:9001"
)

func main() {
	h := handler.New(service.New(userServiceGRPCURL, todoServiceGRPCURL))

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/user/{id}/todos", h.GetUserTodos)
	r.Post("/user/register", h.RegisterUser)
	r.Post("/todo", h.CreateTodo)

	fmt.Println("Listening at :8080")
	http.ListenAndServe("localhost:8080", r)
}
