package utils

import "strings"

// check if a given username is valid
func IsValidUsername(input string) bool {
	return !isStringEmpty(input) && len(strings.TrimSpace(input))<21
}

// isStringEmpty checks if a string is empty.
func isStringEmpty(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}