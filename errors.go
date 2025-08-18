package monoacquiring

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrUnexpectedHTTPStatus      = errors.New("unexpected http status code")
	ErrBadRequestHTTPStatus      = errors.New("bad request http status code")
	ErrForbiddenHTTPStatus       = errors.New("forbidden http status code")
	ErrNotFoundHTTPStatus        = errors.New("not found http status code")
	ErrInternalHTTPStatus        = errors.New("internal server error http status code")
	ErrTooManyRequestsHTTPStatus = errors.New("too many requests http status code")
	ErrMethodNotAllowedStatus    = errors.New("method not allowed")
)

type RequestError struct {
	Err     error
	Code    string
	Message string
}

func (e *RequestError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s (code: %s): %s", e.Message, e.Code, e.Err.Error())
	}

	return fmt.Sprintf("%s (code: %s)", e.Message, e.Code)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}

func newRequestError(err error, code, message string) *RequestError {
	return &RequestError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}
