package idle_errors

import "errors"

var ErrBankNotExist = errors.New("bank does not exist")
var ErrBankAlreadyExists = errors.New("bank already exists")

var ErrUsernameTaken = errors.New("username has already been used")
var ErrUserNotFound = errors.New("user not found")
var ErrPasswordNotUpdated = errors.New("password failed to update")

var ErrUnknownClaimsType = errors.New("unknown claims type")