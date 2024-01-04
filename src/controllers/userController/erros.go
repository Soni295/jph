package controller

import "errors"

var (
	ErrUserNameRequired  = errors.New(`field "name" is required`)
	ErrUserEmailRequired = errors.New(`field "email" is required`)
)