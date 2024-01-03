package request

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  string `json:"is_complete"`
}

type UpdateTaskRequest struct {
}

type GetTasksRequest struct {
	ID string `json:"id"`
}
