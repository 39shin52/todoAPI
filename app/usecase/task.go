package usecase

import (
	"context"
	"fmt"
	"time"

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

func (tu *TaskUsecase) CreateTask(ctx context.Context, r request.CreateTaskRequest) (string, error) {
	taskId := uuid.NewString()

	task := &entity.Task{
		ID:          taskId,
		UserID:      r.UserID,
		Title:       r.Title,
		Description: r.Description,
		IsComplete:  0,
	}

	if err := tu.txAdmin.Transaction(ctx, func(ctx context.Context) error {
		if err := tu.taskRepository.InsertTask(ctx, task); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return fmt.Sprintf("faied to create task: %v", err), err
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
	_, err := tu.taskRepository.SearchTaskByTaskID(taskID)
	if err != nil {
		return err
	}

	if err := tu.txAdmin.Transaction(ctx, func(ctx context.Context) error {
		if err := tu.taskRepository.DeleteTask(ctx, taskID); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}

	return nil
}

func (tu *TaskUsecase) UpdateTask(ctx context.Context, taskID string, task *entity.Task) error {
	newTask, err := tu.taskRepository.SearchTaskByTaskID(taskID)
	if err != nil {
		return err
	}

	if task.Title != "" {
		newTask.Title = task.Title
	}
	if task.Description != "" {
		newTask.Description = task.Description
	}
	if newTask.IsComplete != task.IsComplete {
		newTask.IsComplete = task.IsComplete
	}
	newTask.Updated_at = time.Now()

	if err := tu.txAdmin.Transaction(ctx, func(ctx context.Context) error {
		if err := tu.taskRepository.UpdateTask(ctx, newTask); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	return nil
}
