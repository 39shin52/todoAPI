package repository

type TaskRepository interface {
	InsertTask()
	UpdateTask()
	DeleteTask()
	SearchTask()
}
