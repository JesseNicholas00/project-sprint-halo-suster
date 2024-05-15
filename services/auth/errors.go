package auth

import "errors"

var (
	ErrNipAlreadyExists = errors.New(
		"authService: nip already exists",
	)

	ErrUserNotFound = errors.New(
		"authService: no such user found",
	)
	ErrUserHasNoAccess = errors.New(
		"authService: user doesn't have access",
	)

	ErrInvalidCredentials = errors.New(
		"authService: invalid credentials",
	)
	ErrTokenInvalid = errors.New(
		"authService: invalid access token",
	)
	ErrTokenExpired = errors.New(
		"authService: token expired",
	)
)
