package usecase

import (
	"context"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/domain/repository"
	"github.com/39shin52/todoAPI/app/domain/repository/transaction"
	"github.com/39shin52/todoAPI/app/interfaces/request"
	"github.com/google/uuid"
)

type TaskUsecase struct {
	txAdmin        *transaction.TxAdmin
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository, txAdmin *transaction.TxAdmin) *TaskUsecase {
	return &TaskUsecase{taskRepository: taskRepository, txAdmin: txAdmin}
}

func (tu *TaskUsecase) CreateTask(ctx context.Context, r request.CreateTaskRequest, userId string) (string, error) {
	taskId := uuid.NewString()

	task := &entity.Task{
		UserID:      userId,
		ID:          taskId,
		Title:       r.Title,
		Description: r.Description,
		IsComplete:  r.IsComplete,
	}

	if err := tu.taskRepository.InsertTask(ctx, task); err != nil {
		return "", err
	}

	return taskId, nil
}

func (tu *TaskUsecase) ReadTasks(userID string) ([]entity.Task, error) {
	tasks, err := tu.taskRepository.SelectTasks(userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tu *TaskUsecase) ReadTask(taskID string) (*entity.Task, error) {
	task, err := tu.taskRepository.SearchTaskByTaskID(taskID)
	if err != nil {
		return nil, err
	}

	return task, err
}

func (tu *TaskUsecase) DeleteTask(ctx context.Context, taskID string) error {
	task, err := tu.taskRepository.SearchTaskByTaskID(taskID)
	if err != nil {
		return err
	}

	if err = tu.taskRepository.DeleteTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (tu *TaskUsecase) UpdateTask(ctx context.Context, taskID string) error {
	task, err := tu.taskRepository.SearchTaskByTaskID(taskID)
	if err != nil {
		return err
	}

	if err = tu.taskRepository.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}
