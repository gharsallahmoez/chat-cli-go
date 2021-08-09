package database_test

import (
	"github.com/gharsallahmoez/chat/src/config"
	"github.com/gharsallahmoez/chat/src/database"
	tt "github.com/gharsallahmoez/chat/src/testdata/database"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	tableTest := tt.CreateTTHandler()
	for _, tc := range tableTest {
		conf := config.Database{
			Chat: tc.ChatDB,
		}
		t.Run(tc.Name, func(t *testing.T) {
			_, err := database.Create(&conf.Chat)
			if err != nil && !tc.HasError {
				t.Errorf("expected success , got error: %v", err)
			}
			if err == nil && tc.HasError {
				t.Error("expected error, got nil")
			}
		})
	}
}

