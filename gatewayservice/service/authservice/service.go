package authservice

import (
	"fmt"
	"github.com/behnambm/todo/common/utils/hash"
	"github.com/behnambm/todo/gatewayservice/types"
	"log"
)

type CommandRepo interface {
	CreateUser(types.User) error
}

type QueryRepo interface {
	GetUserByEmail(string) (types.User, error)
	GetToken(int64) (string, error)
	IsValidWithClaim(string) (map[string]string, bool)
}

type AuthService struct {
	command CommandRepo
	query   QueryRepo
}

func New(cmdRepo CommandRepo, queryRepo QueryRepo) AuthService {
	return AuthService{
		command: cmdRepo,
		query:   queryRepo,
	}
}

func (s AuthService) LoginUser(email string, password string) (string, error) {
	user, err := s.query.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("[AUTH SERVICE] LoginUser - unable to fetch user - %w", err)
	}

	hashedInputPassword, hashErr := hash.String(password)
	if hashErr != nil {
		log.Println("[AUTH SERVICE] LoginUser - unable to hash password ", hashErr)

		return "", fmt.Errorf("[AUTH SERVICE] LoginUser - internal error")
	}

	if hashedInputPassword != user.Password {
		return "", fmt.Errorf("[AUTH SERVICE] LoginUser - invalid credentials")
	}

	token, tokenErr := s.query.GetToken(user.ID)
	if tokenErr != nil {
		return "", fmt.Errorf("[AUTH SERVICE] LoginUser - unable to fetch token - %w", tokenErr)
	}

	return token, nil
}

func (s AuthService) Register(user types.User) error {
	_, err := s.query.GetUserByEmail(user.Email)
	if err == nil {
		return fmt.Errorf("[AUTH SERVICE] Register - user already exists - %w", err)
	}

	createErr := s.command.CreateUser(user)
	if createErr != nil {
		return fmt.Errorf("[AUTH SERVICE] Register - unable to create user - %w", createErr)
	}

	return nil
}

func (s AuthService) IsValidWithClaim(token string) (map[string]string, bool) {
	claimMap, valid := s.query.IsValidWithClaim(token)
	if !valid {
		return nil, false
	}

	return claimMap, true
}
