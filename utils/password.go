package utils

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	ret, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return string(ret), err
	}
	return string(ret), err
}

func IsPasswordValid(password string) bool {
	isUpper, isLower, isDigit := false, false, false
	if len(password) < 8 || len(password) > 70 {
		return false
	} else {
		for _, r := range password {
			if unicode.IsUpper(r) && !isUpper {
				isUpper = true
			}
			if unicode.IsLower(r) && !isLower {
				isLower = true
			}
			if unicode.IsDigit(r) && !isDigit {
				isDigit = true
			}
		}
	}
	return isUpper && isLower && isDigit
}
