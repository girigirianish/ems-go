package http

import (
	_userHttpDeliveryMiddleware "github.com/girigirianish/ems-go/internal/user/delivery/http/middleware"
	"github.com/girigirianish/ems-go/internal/user/domain/usecases"
	"github.com/labstack/echo/v4"
)

// RegisterHTTPEndpoints godoc
func RegisterHTTPEndpoints(e *echo.Echo, usecase usecases.UserUseCase) {
	handler := NewUserHandler(usecase)
	userAuthMiddleWare := _userHttpDeliveryMiddleware.InitMiddleware()

	// Public
	e.POST("login", handler.SignIn)
	e.POST("register", handler.SignUp)
	e.POST("user", handler.CreateUserDetail)

	// Login required
	e.GET("users", handler.GetAllUserDetails, userAuthMiddleWare.IsLoggedIn(usecase))
	e.GET("user/education", handler.GetUserEducationDetails, userAuthMiddleWare.IsLoggedIn(usecase))
	e.GET("user/experience", handler.GetUserExperienceDetails, userAuthMiddleWare.IsLoggedIn(usecase))
	e.GET("user/permissions", handler.GetUserPermissions, userAuthMiddleWare.IsLoggedIn(usecase))
}
