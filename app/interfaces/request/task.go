package request

type UpdateTaskRequest struct {
	UserID      string `json:"user_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	IsComplete  int    `json:"is_complete,omitempty"`
}
type CreateTaskRequest struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetTasksRequest struct {
	UserID string `json:"user_id"`
}
