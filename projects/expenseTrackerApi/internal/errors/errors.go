package errors

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type AppError struct {
	Status  int
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}

	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func New(status int, code, message string, err error) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func BadRequest(code, message string, err error) *AppError {
	return New(http.StatusBadRequest, code, message, err)
}

func Unauthorized(code, message string, err error) *AppError {
	return New(http.StatusUnauthorized, code, message, err)
}

func NotFound(code, message string, err error) *AppError {
	return New(http.StatusNotFound, code, message, err)
}

func Conflict(code, message string, err error) *AppError {
	return New(http.StatusConflict, code, message, err)
}

func Internal(err error) *AppError {
	return New(http.StatusInternalServerError, "internal_error", "internal server error", err)
}

func From(err error) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return NotFound("resource_not_found", "resource not found", err)
	}

	return Internal(err)
}
