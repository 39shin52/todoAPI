package infrastructure

import (
	"database/sql"

	"github.com/39shin52/todoAPI/app/domain/repository"
)

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (tr *taskRepositoryImpl) InsertTask() {

}
func (tr *taskRepositoryImpl) UpdateTask() {

}
func (tr *taskRepositoryImpl) DeleteTask() {

}
func (tr *taskRepositoryImpl) SearchTask() {

}
