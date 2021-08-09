package database

import (
	"errors"
	"fmt"
	"github.com/gharsallahmoez/chat/src/config"
	"github.com/gharsallahmoez/chat/src/database/mysql"
)

type ChatHandler interface {
	// db functions
}

// DbHandler holds the handler for chat model
type DbHandler interface {
	ChatHandler
}

// Create creates db handler based on the given config
func Create(chatDB *config.Chat) (DbHandler, error) {
	var chatHandler ChatHandler
	var err error
	switch chatDB.Type {
	case "mysql":
		chatHandler = mysql.NewRepo(chatDB)
	default:
		return nil, errors.New(fmt.Sprintf("%s is an unknown database type", chatDB.Type))
	}

	handlers := struct {
		ChatHandler
	}{
		chatHandler,
	}
	return handlers, err
}
