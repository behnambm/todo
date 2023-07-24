package service

import (
	"fmt"
	"github.com/behnambm/todo/common/utils/hash"
	"github.com/behnambm/todo/userservice/types"
)

type UserRepo interface {
	GetUserByEmail(string) (types.User, error)
	CreateUser(types.User) (types.User, error)
}

type UserService struct {
	repo UserRepo
}

func New(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetUserByEmail(email string) (types.User, error) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return types.User{}, fmt.Errorf("[Service] GetUserByEmail - %w", err)
	}

	return user, nil
}

func (us *UserService) CreateUser(user types.User) (types.User, error) {
	// hash the plain text password
	hashedPassword, err := hash.String(user.Password)
	if err != nil {
		return types.User{}, fmt.Errorf("[Service] CreateUser - %w", err)
	}
	// store the hashed password
	user.Password = hashedPassword

	newUser, createErr := us.repo.CreateUser(user)
	if createErr != nil {
		return types.User{}, fmt.Errorf("[Service] CreateUser - %w", createErr)
	}

	return newUser, nil
}
