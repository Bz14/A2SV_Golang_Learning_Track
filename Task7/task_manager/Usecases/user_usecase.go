package usecases

import (
	"errors"
	"fmt"
	domain "task_manager/Domain"
)

type UserUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(repository domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: repository,
	}
}

func (userUseCase *UserUseCase) Register(user domain.User) (interface{}, error) {
	if user.UserName == "" || user.Password == "" || user.Role == ""{
		fmt.Println("Password" , user.Password)
		return nil, errors.New("invalid credentials")
	}
	userSign, err := userUseCase.userRepository.Register(user)
	if err != nil {
		return nil, err
	}
	return userSign, nil
}

func (userUseCase *UserUseCase) Login(user domain.User) (string, error) {
	if user.UserName == "" || user.Password == "" {
		return "", errors.New("invalid credentials")
	}
	userLogin, err := userUseCase.userRepository.Login(user)
	if err != nil {
		return "", err
	}
	return userLogin, nil
}
