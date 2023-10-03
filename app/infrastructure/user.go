package infrastructure

import (
	"context"
	"database/sql"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/domain/repository"
	"github.com/39shin52/todoAPI/app/domain/repository/transaction"
)

type userRepositoryImpl struct {
	db *sql.DB
	t  transaction.Transaction
}

func NewUserRepository(db *sql.DB, t transaction.Transaction) repository.UserRepository {
	return &userRepositoryImpl{db: db, t: t}
}

func (ur *userRepositoryImpl) SelectUser(name string) (*entity.User, error) {
	user := new(entity.User)

	row := ur.db.QueryRow("SELECT user_id, user_name, mail, work from user where user_name=?", name)
	if err := row.Scan(&user.ID, &user.UserName, &user.Mail, &user.Work); err != nil {
		return nil, err
	}

	return user, nil
}
func (ur *userRepositoryImpl) SelectUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)

	rows, err := ur.db.Query("select user_id, password, token, user_name, mail, work from user")
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
	_, err := ur.t.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) (interface{}, error) {
		_, err := tx.Exec("delete from user where user_name=?", name)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}
func (ur *userRepositoryImpl) UpdateUser(ctx context.Context, user entity.User) error {
	_, err := ur.t.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) (interface{}, error) {
		_, err := tx.Exec("update user set user_name=?, email=?, work=?", user.UserName, user.Mail, user.Work)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}
func (ur *userRepositoryImpl) InsertUser(ctx context.Context, user entity.User) error {
	_, err := ur.t.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) (interface{}, error) {
		_, err := tx.Exec("insert into user (user_id,password,token,user_name,mail,work) values (?,?,?,?,?,?)", user.ID, user.Password, user.Token, user.UserName, user.Mail, user.Work)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}
