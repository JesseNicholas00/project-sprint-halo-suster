package auth

import "errors"

var (
	ErrNipNotFound = errors.New(
		"authRepository: no such nip found",
	)
	ErrDuplicateUser = errors.New(
		"authRepository: duplicate id or nip found",
	)
)
