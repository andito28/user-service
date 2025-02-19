package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordInCorrect    = errors.New("password incorrect")
	ErrUserNameExists       = errors.New("username already exists")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUserNameExists,
	ErrPasswordDoesNotMatch,
}
