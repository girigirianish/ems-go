package repositories

import (
	"context"

	"github.com/girigirianish/ems-go/internal/user/domain/entities"
)

// UserRepository godoc
type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.UserEntity) error
	GetUser(ctx context.Context, email string, password string) (*entities.UserEntity, error)
	GetAllUsersDetail(ctx context.Context) (*[]entities.UserDetailEntity, error)
	GetUserEducationDetail(ctx context.Context, id int64) (*[]entities.UserEducationDetail, error)
	GetUserExperienceDetail(ctx context.Context, id int64) (*[]entities.UserExperienceDetail, error)
	CreateUserDetail(ctx context.Context, user *entities.UserDetailEntity) error
	CreateUserEducationDetails(ctx context.Context, userEduDetails []entities.UserEducationDetail) error
	CreateUserExperienceDetails(ctx context.Context, userEduDetails []entities.UserExperienceDetail) error
}
