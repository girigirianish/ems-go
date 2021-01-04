package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/girigirianish/ems-go/internal/user/domain/entities"
	"github.com/girigirianish/ems-go/internal/user/domain/usecases"
	"github.com/girigirianish/ems-go/internal/user/models"
	"github.com/girigirianish/ems-go/internal/util"
	"github.com/labstack/echo/v4"
)

// UserAPIResponse godoc
type UserAPIResponse struct {
	Message string `json:"message"`
}

// UserHandler godoc
type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

// SignInput godoc
type SignInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInResponse godoc
type SignInResponse struct {
	Token string `json:"token"`
}

// NewUserHandler godoc
func NewUserHandler(usecase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: usecase,
	}
}

// SignUp godoc
func (userHandler *UserHandler) SignUp(echoContext echo.Context) error {
	var user entities.UserEntity
	err := echoContext.Bind(&user)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := echoContext.Request().Context()
	err = userHandler.UserUseCase.SignUp(ctx, user.Email, user.Password)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), UserAPIResponse{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, UserAPIResponse{Message: "success"})
}

// SignIn godoc
func (userHandler *UserHandler) SignIn(echoContext echo.Context) error {
	userDetails := new(SignInput)
	err := echoContext.Bind(userDetails)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := echoContext.Request().Context()
	token, err := userHandler.UserUseCase.SignIn(ctx, userDetails.Email, userDetails.Password)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), UserAPIResponse{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, SignInResponse{Token: token})
}

// GetAllUserDetails godoc
func (userHandler *UserHandler) GetAllUserDetails(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	users, err := userHandler.UserUseCase.GetAllUsersDetails(ctx)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), UserAPIResponse{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, users)
}

// GetUserEducationDetails godoc
func (userHandler *UserHandler) GetUserEducationDetails(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	idP, err := strconv.Atoi(echoContext.QueryParam("id"))
	id := int64(idP)
	users, err := userHandler.UserUseCase.GetUserEducationDetail(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), UserAPIResponse{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, users)
}

// GetUserExperienceDetails godoc
func (userHandler *UserHandler) GetUserExperienceDetails(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	idP, err := strconv.Atoi(echoContext.QueryParam("id"))
	id := int64(idP)
	users, err := userHandler.UserUseCase.GetUserExperienceDetail(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), UserAPIResponse{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, users)
}

// CreateUserDetail godoc
func (userHandler *UserHandler) CreateUserDetail(echoContext echo.Context) error {
	body, err := ioutil.ReadAll(echoContext.Request().Body)
	var userDetail models.UserDetailsReqResponse = models.UserDetailsReqResponse{}
	json.Unmarshal([]byte(body), &userDetail)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	ctx := echoContext.Request().Context()
	e := userHandler.UserUseCase.CreateUserDetail(
		ctx,
		userDetail.PersonalDetails,
		userDetail.EducationDetails,
		userDetail.ExperienceDetails,
	)
	if e != nil {
		return echoContext.JSON(util.GetStatusCode(err), UserAPIResponse{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, UserAPIResponse{Message: "success"})
}

// GetUserPermissions godoc
func (userHandler *UserHandler) GetUserPermissions(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	idP, err := strconv.Atoi(echoContext.QueryParam("roleId"))
	roleID := int64(idP)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), UserAPIResponse{Message: err.Error()})
	}
	var permissions = userHandler.UserUseCase.GetUserPermissions(ctx, roleID)
	return echoContext.JSON(http.StatusOK, permissions)
}
