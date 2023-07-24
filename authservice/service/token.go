package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
)

type TokenService struct {
	SecretKey  string
	SignMethod *jwt.SigningMethodHMAC
}

func New(secret string) TokenService {
	return TokenService{
		SecretKey:  secret,
		SignMethod: jwt.SigningMethodHS256,
	}
}

// GetToken is used to generate a JWT token containing UserId in claims
func (s TokenService) GetToken(userID int) (string, error) {
	token := jwt.NewWithClaims(s.SignMethod, jwt.MapClaims{
		"uid": userID,
	})

	return token.SignedString([]byte(s.SecretKey))
}

func (s TokenService) GetClaim(tokenString string) (jwt.MapClaims, error) {
	token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("[Service] GetClaim - Unexpected signing method: %v\n", token.Header["alg"])
		}

		return []byte(s.SecretKey), nil
	})

	if parseErr != nil {
		return nil, fmt.Errorf("[Service] GetClaim - claim parse error %w", parseErr)
	}

	if !token.Valid {
		return nil, fmt.Errorf("[Service] GetClaim - invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("[Service] GetClaim - claim type is not valid")
	}

	return claims, nil
}

func (s TokenService) IsTokenValid(tokenString string) bool {
	_, claimErr := s.GetClaim(tokenString)

	return claimErr == nil
}
