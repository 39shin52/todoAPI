package usecase

import (
	"context"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/domain/repository"
	"github.com/39shin52/todoAPI/app/domain/repository/transaction"
)

type UserUsecase struct {
	txAdmin        *transaction.TxAdmin
	userRepository repository.UserRepository
}

// contextはinterfaceで宣言します
func NewUserUsecase(userRepository repository.UserRepository, txAdmin *transaction.TxAdmin) *UserUsecase {
	return &UserUsecase{userRepository: userRepository, txAdmin: txAdmin}
}

func (uu *UserUsecase) SelectUser(name string) (*entity.User, error) {
	user, err := uu.userRepository.SelectUser(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UserUsecase) SelectUsers() ([]entity.User, error) {
	users, err := uu.userRepository.SelectUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, user entity.User) error {
	return uu.userRepository.UpdateUser(ctx, user)
}

func (uu *UserUsecase) DeleteUser(ctx context.Context, name string) error {
	return uu.userRepository.DeleteUser(ctx, name)
}
