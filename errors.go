package monoacquiring

import "github.com/pkg/errors"

var (
	ErrBlankAPIKey    = errors.New("api key is blank")
	ErrBlankBaseURL   = errors.New("base URL is blank")
	ErrInvalidBaseURL = errors.New("base URL is invalid")
)
