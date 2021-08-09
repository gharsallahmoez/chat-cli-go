package utils

import (
	common "github.com/gharsallahmoez/chat/src/testdata"
)


// TTIsValidUsername is table test for the test TestIsValidUsername.
type TTIsValidUsername struct {
	Name    string
	Input   string
	IsValid bool
}

// TTIsEmptyString is table test for the test TestIsEmptyString.
type TTIsEmptyString struct {
	Name    string
	Input   string
	IsEmpty bool
}

// CreateTTIsValidUsername creates table test for IsValidUsername test.
func CreateTTIsValidUsername() (tt []TTIsValidUsername) {
	tt = []TTIsValidUsername{
		{
			Name:     "valid username",
			Input:  "valid",
			IsValid: true,
		},
		{
			Name:     "empty username",
			Input:  "",
			IsValid: false,
		},
		{
			Name:     "more than 20 characters username",
			Input: common.RandomString(50),
			IsValid: false,
		},
	}
	return
}

// CreateTTIsEmptyString creates table test for IsEmptyString test.
func CreateTTIsEmptyString() (tt []TTIsEmptyString) {
	tt = []TTIsEmptyString{
		{
			Name:     "valid username",
			Input:  "valid",
			IsEmpty: false,
		},
		{
			Name:     "empty string",
			Input:  "",
			IsEmpty: true,
		},
	}
	return
}
