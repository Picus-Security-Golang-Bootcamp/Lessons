package httpErrors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/h4yfans/patika-bookstore/internal/api"
	"gorm.io/gorm"
)

var (
	InternalServerError = errors.New("Internal Server Error")
	NotFound            = errors.New("Not Found")
	RequestTimeoutError = errors.New("Request Timeout")
	CannotBindGivenData = errors.New("Could not bind given data")
	ValidationError     = errors.New("Validation failed for given payload")
	UniqueError         = errors.New("Item should be unique on database")
)

type RestError api.APIResponse

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.Code, e.Message, e.Details)
}

func (e RestError) Status() int {
	return int(e.Code)
}

func (e RestError) Causes() interface{} {
	return e.Details
}

func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		Code:    int64(status),
		Message: err,
		Details: causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		Code:    http.StatusInternalServerError,
		Message: InternalServerError.Error(),
		Details: causes,
	}
	return result
}

// ParseErrors Parser of error string messages returns RestError
func ParseErrors(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, NotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
	case errors.Is(err, CannotBindGivenData):
		return NewRestError(http.StatusBadRequest, CannotBindGivenData.Error(), err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NewRestError(http.StatusNotFound, gorm.ErrRecordNotFound.Error(), err)
	case strings.Contains(err.Error(), "validation"):
		return NewRestError(http.StatusBadRequest, ValidationError.Error(), err)
	case strings.Contains(err.Error(), "23505"):
		return NewRestError(http.StatusBadRequest, UniqueError.Error(), err)

	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
