package apperror

import "net/http"

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewError(msg string, status int) *Error {
	return &Error{msg, status}
}

func (e *Error) Error() string {
	return e.Message
}

func NotFound(msg string) error {
	return NewError(msg, http.StatusNotFound)
}
func BadRequest(msg string) error {
	return NewError(msg, http.StatusBadRequest)
}
func InternalServer(msg string) error {
	return NewError(msg, http.StatusInternalServerError)
}
func Unauthorized(msg string) error {
	return NewError(msg, http.StatusUnauthorized)
}
