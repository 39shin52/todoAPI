package request

type CreateTaskRequest struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
}

type GetTasksRequest struct {
	UserID string `json:"user_id"`
}
