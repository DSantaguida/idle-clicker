package password

import "github.com/dsantaguida/idle-clicker/pkg/idle_errors"

const lengthRequirement = 8

func CheckPasswordCompliance(password string) error {
	if len(password) < lengthRequirement {
		return idle_errors.ErrPasswordNotCompliant
	}
	return nil
}