package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/domain/repository"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (ur *userRepositoryImpl) SelectUser(name string) (*entity.User, error) {
	req := "select user_id, user_name, mail, work from users where user_name = ?"
	user := new(entity.User)

	err := ur.db.QueryRow(req, name).Scan(&user.ID, &user.UserName, &user.Mail, &user.Work)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user's name %s  is not found in db", name)
		} else {
			log.Printf("some error occured: %v", err)
		}

		return nil, err
	}

	return user, nil
}
func (ur *userRepositoryImpl) SelectUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)

	rows, err := ur.db.Query("select user_id, password, token, user_name, mail, work from users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u entity.User
		if err := rows.Scan(&u.ID, &u.Password, &u.Token, &u.UserName, &u.Mail, &u.Work); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepositoryImpl) DeleteUser(ctx context.Context, name string) error {
	req := "delete from users where user_name=?"

	if _, err := ur.db.ExecContext(ctx, req, name); err != nil {
		return err
	}

	return nil
}

func (ur *userRepositoryImpl) UpdateUser(ctx context.Context, user entity.User) error {
	req := "update users set user_name=?, email=?, work=?"

	if _, err := ur.db.ExecContext(ctx, req, user.UserName, user.Mail, user.Work); err != nil {
		return err
	}

	return nil
}

func (ur *userRepositoryImpl) InsertUser(ctx context.Context, user entity.User) error {
	req := "insert into users (user_id,password,token,user_name,mail,work) values (?,?,?,?,?,?)"

	if _, err := ur.db.ExecContext(ctx, req, user.ID, user.Password, user.Token, user.UserName, user.Mail, user.Work); err != nil {
		return err
	}

	return nil
}
