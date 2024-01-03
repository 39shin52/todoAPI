package response

import "time"

type GetTasksResponse struct {
	tasks []GetTaskResponse
}

type GetTaskResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsComplete  string    `json:"is_complete"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
