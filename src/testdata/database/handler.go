package db

import "github.com/gharsallahmoez/chat/src/config"

// TTCreateHandler represents table test structure of CreateHandler test
type TTCreateHandler struct {
	Name       string
	ChatDB config.Chat
	HasError   bool
}

// CreateTTHandler creates table test for Create test
func CreateTTHandler() []TTCreateHandler {
	conf, _ := config.MakeConfig()
	return []TTCreateHandler{
		{
			Name: "valid config for chat with mysql",
			ChatDB: config.Chat{
				Type:     "mysql",
				Host:     conf.Database.Chat.Host,
				Username: conf.Database.Chat.Username,
				Password: conf.Database.Chat.Password,
			},
			HasError: false,
		},
		{
			Name:       "unsupported db for chat",
			ChatDB: config.Chat{Type: "invalid"},
			HasError:   true,
		},
	}
}
