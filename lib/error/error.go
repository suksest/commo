package error

import (
	"errors"
)

// ErrResponse for returning error message
type ErrResponse struct {
	Message string `json:"message"`
}

// Errors
var (
	// Request
	ErrEmptyRequestBody = errors.New("request.empty_request_body")
	ErrBadInput         = errors.New("request.input_invalid")

	// Auth
	ErrIdentifierNotFound = errors.New("auth.identifier_not_found")
	ErrPasswordWrong      = errors.New("auth.password_wrong")
	ErrIdentifierEmpty    = errors.New("auth.identifier_empty")
	ErrPasswordEmpty      = errors.New("auth.password_empty")
	ErrSigningJWT         = errors.New("auth.signing_token_error")
	ErrNotRegistered      = errors.New("auth.account_not_registered")
	ErrPhoneNotUnique     = errors.New("auth.phone_is_already_registered")
	ErrInvalidToken       = errors.New("auth.invalid_token")

	ErrConnection     = errors.New("db.connection_error")
	ErrRecordNotFound = errors.New("db.record_not_found")
)
