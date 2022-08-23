package repository

import (
	"context"
	"todo-service/todo"
)

type DB struct {
	todos map[int32][]todo.Todo // [user-id] => list of todos
	genID func() int32
}

func NewDB() *DB {
	return &DB{
		todos: make(map[int32][]todo.Todo),
		genID: NewIDGenerator(),
	}
}

func (d *DB) CreateTodo(ctx context.Context, t todo.Todo) (todo.Todo, error) {
	t.ID = d.genID()

	d.todos[t.UserID] = append(d.todos[t.UserID], t)

	return t, nil
}

func (d *DB) GetTodosByUserID(ctx context.Context, userID int32) ([]todo.Todo, error) {
	return d.todos[userID], nil
}

func NewIDGenerator() func() int32 {
	var id int32

	return func() int32 {
		id++
		return id
	}
}
