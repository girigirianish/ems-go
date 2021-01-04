package util

import (
	"net/http"

	genericerrors "github.com/girigirianish/ems-go/internal/generic_errors"
	"github.com/labstack/gommon/log"

	"gopkg.in/go-playground/validator.v9"
)

// Size ...
const (
	MB = 1 << 20
)

// Sizer ...
type Sizer interface {
	Size() int64
}

// GetStatusCode ...
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	log.Error(err)
	switch err {
	case genericerrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case genericerrors.ErrNotFound:
		return http.StatusNotFound
	case genericerrors.ErrConflict:
		return http.StatusConflict
	case genericerrors.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// IsRequestValid ...
func IsRequestValid(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
