package entity

type Todo struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	Done   bool   `json:"done"`
	UserID int    `json:"user_id"`
}
