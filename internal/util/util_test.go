package util_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	genericerrors "github.com/girigirianish/ems-go/internal/generic_errors"
	"github.com/girigirianish/ems-go/internal/user/domain/entities"
	"github.com/girigirianish/ems-go/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestGetStatusCode(t *testing.T) {
	var response int

	response = util.GetStatusCode(nil)
	assert.Equal(t, response, http.StatusOK)

	response = util.GetStatusCode(genericerrors.ErrInternalServerError)
	assert.Equal(t, response, http.StatusInternalServerError)

	response = util.GetStatusCode(genericerrors.ErrNotFound)
	assert.Equal(t, response, http.StatusNotFound)

	response = util.GetStatusCode(genericerrors.ErrConflict)
	assert.Equal(t, response, http.StatusConflict)

	response = util.GetStatusCode(genericerrors.ErrBadParamInput)
	assert.Equal(t, response, http.StatusBadRequest)

	response = util.GetStatusCode(errors.New("unknown"))
	assert.Equal(t, response, http.StatusInternalServerError)

}

func TestIsRequestValid(t *testing.T) {
	mockUserEntity := entities.UserEntity{}
	valid, err := util.IsRequestValid(mockUserEntity)

	assert.Error(t, err)
	assert.False(t, valid)

	mockUserEntity = entities.UserEntity{
		Email:     "giri.girianish@gmail.com",
		Password:  "Content",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	valid, err = util.IsRequestValid(mockUserEntity)
	assert.NoError(t, err)
	assert.True(t, valid)
}
