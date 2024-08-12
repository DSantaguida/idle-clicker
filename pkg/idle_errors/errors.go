package idle_errors

import "errors"

var ErrBankNotExist = errors.New("bank does not exist")
var ErrBankAlreadyExists = errors.New("Bank already exists")
