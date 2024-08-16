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

//JWT errors
var ErrUnknownClaimsType = errors.New("unknown claims type")