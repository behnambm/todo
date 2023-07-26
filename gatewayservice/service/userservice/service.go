package userservice

import (
	"fmt"
	"github.com/behnambm/todo/common/utils/hash"
	"github.com/behnambm/todo/gatewayservice/types"
)

type CommandRepo interface {
	CreateUser(types.User) error
}

type QueryRepo interface {
	GetUserByID(int64) (types.User, error)
}

type UserService struct {
	command CommandRepo
	query   QueryRepo
}

func New(cmdRepo CommandRepo, queryRepo QueryRepo) UserService {
	return UserService{
		command: cmdRepo,
		query:   queryRepo,
	}
}

func (us UserService) GetUserByID(userId int64) (types.User, error) {
	user, err := us.query.GetUserByID(userId)
	if err != nil {
		return types.User{}, fmt.Errorf("[User Service] GetUserByID - %w", err)
	}

	return user, nil
}

func (us UserService) CreateUser(user types.User) error {
	// hash the plain text password
	hashedPassword, err := hash.String(user.Password)
	if err != nil {
		return fmt.Errorf("[User Service] CreateUser - %w", err)
	}
	// store the hashed password
	user.Password = hashedPassword

	createErr := us.command.CreateUser(user)
	if createErr != nil {
		return fmt.Errorf("[User Service] CreateUser - %w", createErr)
	}

	return nil
}
