package middleware

import (
	"net/http"
	"strings"

	"github.com/girigirianish/ems-go/internal/user/domain/usecases"
	"github.com/labstack/echo/v4"
)

// UserAuthMiddleWare godoc
type UserAuthMiddleWare struct {
}

const unAuthorizedHTTPMessage = "Unauthorized"

// CORS godoc
func (userAuthMiddleware *UserAuthMiddleWare) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// IsLoggedIn godoc
func (userAuthMiddleware *UserAuthMiddleWare) IsLoggedIn(usecase usecases.UserUseCase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return echo.NewHTTPError((http.StatusUnauthorized), unAuthorizedHTTPMessage)
			}

			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 {
				return echo.NewHTTPError((http.StatusUnauthorized), unAuthorizedHTTPMessage)
			}

			if headerParts[0] != "Bearer" {
				return echo.NewHTTPError((http.StatusUnauthorized), "")
			}

			signedUser, err := usecase.ParseToken(c.Request().Context(), headerParts[1])
			if err != nil {

				return echo.NewHTTPError(http.StatusUnauthorized, unAuthorizedHTTPMessage)
			}
			c.Set("currentSignedUser", signedUser)
			return next(c)
		}

	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *UserAuthMiddleWare {
	return &UserAuthMiddleWare{}
}
