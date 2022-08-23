package handler

type CreateTodoRequest struct {
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}
