package repository

import (
	"context"

	"github.com/39shin52/todoAPI/app/domain/entity"
)

type TaskRepository interface {
	InsertTask(context.Context, *entity.Task) error
	UpdateTask(context.Context, *entity.Task) error
	DeleteTask(context.Context, *entity.Task) error
	SearchTaskByTaskID(string) (*entity.Task, error)
	SearchTaskByTitle(string) (*entity.Task, error)
	SelectTasks(string) ([]entity.Task, error)
}
