package utils

import (
	tt "github.com/gharsallahmoez/chat/src/testdata/utils"
	"testing"
)

func TestIsValidUsername(t *testing.T) {
	t.Parallel()
	for _, tc := range tt.CreateTTIsValidUsername() {
		t.Run(tc.Name, func(t *testing.T) {
			ok := IsValidUsername(tc.Input)
			if ok && !tc.IsValid {
				t.Errorf("expect error, got success")
			}
			if !ok && tc.IsValid {
				t.Errorf("expect success, got error")
			}
		})
	}
}


func TestIsStringEmpty(t *testing.T) {
	t.Parallel()
	for _, tc := range tt.CreateTTIsEmptyString() {
		t.Run(tc.Name, func(t *testing.T) {
			ok := isStringEmpty(tc.Input)
			if ok && !tc.IsEmpty {
				t.Errorf("expect error, got success")
			}
			if !ok && tc.IsEmpty {
				t.Errorf("expect success, got error")
			}
		})
	}
}
