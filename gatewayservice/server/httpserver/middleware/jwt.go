package middleware

import (
	"github.com/behnambm/todo/gatewayservice/types"
	"github.com/behnambm/todo/gatewayservice/types/constants"
	"github.com/labstack/echo/v4"
	"strconv"
)

type AuthService interface {
	IsValidWithClaim(string) (map[string]string, bool)
}

type UserService interface {
	GetUserByID(int64) (types.User, error)
}

// Auth will extract the JWT token from Authorization header and validate the token and user
// and if the user is authenticated the User object will be stored in the request context for further use
func Auth(userSrv UserService, authSrv AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isAuthenticated := false

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader != "" {
				claim, valid := authSrv.IsValidWithClaim(authHeader)
				if valid {
					uid, err := strconv.Atoi(claim["uid"])
					if err == nil {
						currentUser, userErr := userSrv.GetUserByID(int64(uid))
						if userErr == nil {
							isAuthenticated = true
							c.Set(constants.CurrentUserKey, currentUser)
						}
					}
				}
			}
			c.Set(constants.IsAuthenticatedKey, isAuthenticated)

			return next(c)
		}
	}
}
