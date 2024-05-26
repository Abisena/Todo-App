package helper

import (
	"strings"
)

func ValidateData(username string, email string) (bool, bool) {
	isValidUsername := len(username) >= 2 && 2 < len(username)
	isValidEmail := strings.Contains("@", email)
	return isValidUsername, isValidEmail
}