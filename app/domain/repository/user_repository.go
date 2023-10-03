package repository

import (
	"context"

	"github.com/39shin52/todoAPI/app/domain/entity"
)

type UserRepository interface {
	SelectUser(string) (*entity.User, error)
	SelectUsers() ([]entity.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, entity.User) error
	InsertUser(context.Context, entity.User) error
}
