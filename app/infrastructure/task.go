package infrastructure

import (
	"context"
	"database/sql"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/domain/repository"
	"github.com/39shin52/todoAPI/app/domain/repository/transaction"
)

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB, t transaction.TxAdmin) repository.TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (tr *taskRepositoryImpl) InsertTask(ctx context.Context, task *entity.Task) error {
	req := `insert into tasks (user_id,task_id,title,description,is_complete) values (?,?,?,?,?)`

	if _, err := tr.db.ExecContext(ctx, req, task.UserID, task.ID, task.Title, task.Description, task.IsComplete); err != nil {
		return err
	}

	return nil
}

func (tr *taskRepositoryImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	req := `update`

	if _, err := tr.db.ExecContext(ctx, req); err != nil {
		return err
	}

	return nil
}

func (tr *taskRepositoryImpl) DeleteTask(ctx context.Context, task *entity.Task) error {
	req := `delete`

	if _, err := tr.db.ExecContext(ctx, req); err != nil {
		return err
	}
	return nil
}

func (tr *taskRepositoryImpl) SearchTaskByTaskID(taskID string) (*entity.Task, error) {
	task := new(entity.Task)
	req := `select task_id,title,descripion from tasks where task_id=?`

	row := tr.db.QueryRow(req, taskID)
	if err := row.Scan(&task.UserID, &task.ID, &task.Title, &task.Description, &task.IsComplete); err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *taskRepositoryImpl) SearchTaskByTitle(title string) (*entity.Task, error) {
	task := new(entity.Task)
	req := `select task_id,title,descripion from tasks where title=?`

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
