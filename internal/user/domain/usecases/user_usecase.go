package usecases

import (
	"context"

	"github.com/girigirianish/ems-go/internal/user/domain/entities"
)

// UserUseCase godoc
type UserUseCase interface {
	SignIn(ctx context.Context, email string, password string) (string, error)
	SignUp(ctx context.Context, email string, password string) error
	ParseToken(ctx context.Context, accessToken string) (*entities.UserEntity, error)
	GetAllUsersDetails(ctx context.Context) (*[]entities.UserDetailEntity, error)
	GetUserEducationDetail(ctx context.Context, id int64) (*[]entities.UserEducationDetail, error)
	GetUserExperienceDetail(ctx context.Context, id int64) (*[]entities.UserExperienceDetail, error)
	CreateUserDetail(ctx context.Context,
		userDetail *entities.UserDetailEntity,
		userEduDetail []entities.UserEducationDetail,
		userExpDetail []entities.UserExperienceDetail,
	) error
	GetUserPermissions(ctx context.Context, roleID int64) [1]string
}
