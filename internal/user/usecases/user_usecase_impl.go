package usecases

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	genericerrors "github.com/girigirianish/ems-go/internal/generic_errors"
	auth "github.com/girigirianish/ems-go/internal/user"
	"github.com/girigirianish/ems-go/internal/user/domain/entities"
	authRepositories "github.com/girigirianish/ems-go/internal/user/domain/repositories"
	userUsecases "github.com/girigirianish/ems-go/internal/user/domain/usecases"
)

// AuthClaims godoc
type AuthClaims struct {
	jwt.StandardClaims
	User *entities.UserEntity `json:"user"`
}

// UserUseCaseImpl godoc
type UserUseCaseImpl struct {
	userRepo       authRepositories.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

// NewUserUseCaseImpl godoc
func NewUserUseCaseImpl(userRepo authRepositories.UserRepository, hashSalt string, signingKey []byte, expireDuration time.Duration) userUsecases.UserUseCase {
	return &UserUseCaseImpl{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * expireDuration,
	}
}

// SignUp godoc
func (usecases *UserUseCaseImpl) SignUp(ctx context.Context, email string, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(usecases.hashSalt))
	user := &entities.UserEntity{
		Email:     email,
		Password:  fmt.Sprintf("%x", pwd.Sum(nil)),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	return usecases.userRepo.CreateUser(ctx, user)
}

// SignIn godoc
func (usecases *UserUseCaseImpl) SignIn(ctx context.Context, email string, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(usecases.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := usecases.userRepo.GetUser(ctx, email, password)
	if err != nil {
		return "", genericerrors.ErrNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(usecases.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(usecases.signingKey)
}

// GetAllUsersDetails godoc
func (usecases *UserUseCaseImpl) GetAllUsersDetails(ctx context.Context) (*[]entities.UserDetailEntity, error) {
	userDetails, err := usecases.userRepo.GetAllUsersDetail(ctx)
	if err != nil {
		return nil, genericerrors.ErrNotFound
	}
	return userDetails, nil
}

// GetUserEducationDetail godoc
func (usecases *UserUseCaseImpl) GetUserEducationDetail(ctx context.Context, id int64) (*[]entities.UserEducationDetail, error) {
	userEduDetails, err := usecases.userRepo.GetUserEducationDetail(ctx, id)
	if err != nil {
		return nil, genericerrors.ErrNotFound
	}
	return userEduDetails, nil
}

// GetUserExperienceDetail godoc
func (usecases *UserUseCaseImpl) GetUserExperienceDetail(ctx context.Context, id int64) (*[]entities.UserExperienceDetail, error) {
	userExpDetails, err := usecases.userRepo.GetUserExperienceDetail(ctx, id)
	if err != nil {
		return nil, genericerrors.ErrNotFound
	}
	return userExpDetails, nil
}

// ParseToken godoc
func (usecases *UserUseCaseImpl) ParseToken(ctx context.Context, accessToken string) (*entities.UserEntity, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return usecases.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}
	return nil, auth.ErrInvalidAccessToken
}

// CreateUserDetail godoc
func (usecases *UserUseCaseImpl) CreateUserDetail(
	ctx context.Context,
	userDetail *entities.UserDetailEntity,
	userEduDetail []entities.UserEducationDetail,
	userExpDetail []entities.UserExperienceDetail,
) (err error) {
	e := usecases.userRepo.CreateUserDetail(ctx, userDetail)
	if e != nil {
		return e
	}
	userEduError := usecases.userRepo.CreateUserEducationDetails(ctx, userEduDetail)
	if userEduError != nil {
		return userEduError
	}
	userExpError := usecases.userRepo.CreateUserExperienceDetails(ctx, userExpDetail)
	return userExpError
}

// GetUserPermissions godoc hardcoded permission (Dynamic is out of scope of the task)
func (usecases *UserUseCaseImpl) GetUserPermissions(ctx context.Context, roleID int64) [1]string {
	var permission [1]string

	if roleID == 1 {
		permission[0] = "ViewEmployeeDetails"
		return permission
	}
	return permission
}
