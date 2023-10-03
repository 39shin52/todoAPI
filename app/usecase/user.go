package usecase

import (
	"context"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/domain/repository"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

// contextはinterfaceで宣言します
func NewUserRepository(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (uu *UserUsecase) SelectUser(name string) (*entity.User, error) {
	return uu.userRepository.SelectUser(name)
}

func (uu *UserUsecase) SelectUsers() ([]entity.User, error) {
	return uu.userRepository.SelectUsers()
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, user entity.User) error {
	return uu.userRepository.UpdateUser(ctx, user)
}

func (uu *UserUsecase) DeleteUser(ctx context.Context, name string) error {
	return uu.userRepository.DeleteUser(ctx, name)
}
