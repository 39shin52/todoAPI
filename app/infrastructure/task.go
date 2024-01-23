package infrastructure

import (
	"context"
	"database/sql"
	"log"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/domain/repository"
)

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (tr *taskRepositoryImpl) InsertTask(ctx context.Context, task *entity.Task) error {
	req := `insert into tasks (id, user_id,title, description, is_complete) values (?,?,?,?,?)`

	if _, err := tr.db.ExecContext(ctx, req, task.ID, task.UserID, task.Title, task.Description, task.IsComplete); err != nil {
		return err
	}

	return nil
}

func (tr *taskRepositoryImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	req := `update tasks set title=?, description=?, is_complete=?, updated_at=? where id=?`

	if _, err := tr.db.ExecContext(ctx, req, task.Title, task.Description, task.IsComplete, task.Updated_at, task.ID); err != nil {
		return err
	}

	return nil
}

func (tr *taskRepositoryImpl) DeleteTask(ctx context.Context, taskID string) error {
	req := `delete from tasks where id=?`

	if _, err := tr.db.ExecContext(ctx, req, taskID); err != nil {
		return err
	}
	return nil
}

func (tr *taskRepositoryImpl) SearchTaskByTaskID(taskID string) (*entity.Task, error) {
	task := new(entity.Task)
	req := `select id, user_id, title, description, is_complete, created_at, updated_at from tasks where id=?`

	err := tr.db.QueryRow(req, taskID).Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.IsComplete, &task.Created_at, &task.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("taskID %s not found in db", taskID)
		} else {
			log.Printf("some error occured: %v", err)
		}

		return nil, err
	}

	return task, nil
}

func (tr *taskRepositoryImpl) SearchTaskByTitle(title string) (*entity.Task, error) {
	task := new(entity.Task)
	req := `select task_id,title,description from tasks where title=?`

	row := tr.db.QueryRow(req, title)
	if err := row.Scan(&task.UserID, &task.ID, &task.Title, &task.Description, &task.IsComplete); err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *taskRepositoryImpl) SelectTasks(userID string) ([]entity.Task, error) {
	tasks := make([]entity.Task, 0)
	req := `select id, title, description, is_complete, created_at, updated_at from tasks where user_id = ?`

	rows, err := tr.db.Query(req, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no rows on %s", userID)
		} else {
			log.Printf("some error occur: %v", err)
		}

		return nil, err
	}

	for rows.Next() {
		var task entity.Task
		if err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsComplete, &task.Created_at, &task.Updated_at); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil

}
