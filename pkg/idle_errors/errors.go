package idle_errors

import "errors"

//Bank errors
var ErrBankNotExist = errors.New("bank does not exist")
var ErrBankAlreadyExists = errors.New("bank already exists")

//Authentication errors
var ErrUsernameTaken = errors.New("username has already been used")
var ErrUserNotFound = errors.New("user not found")
var ErrPasswordNotUpdated = errors.New("password failed to update")
var ErrPasswordNotCompliant = errors.New("password is not compliant")
var ErrIncorrectPassword = errors.New("could not find a user with a matching username and password")

//JWT errors
var ErrUnknownClaimsType = errors.New("unknown claims type")
var ErrExpiredToken = errors.New("jwt is expired")
var ErrTokenNotInHeader = errors.New("failed to get login token")
var ErrNoMetadata = errors.New("failed to get incoming metadata")