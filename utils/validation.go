package utils

import (
	"errors"
	"regexp"
)

func EmailValidation(email string) (string, error) {

	emailPattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	validEmail := emailPattern.MatchString(email)
	if !validEmail {
		return "", errors.New("Email not valid")
	}
	return email, nil
}
