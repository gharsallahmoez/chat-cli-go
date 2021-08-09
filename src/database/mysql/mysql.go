package mysql

import "github.com/gharsallahmoez/chat/src/config"

type Repo struct {
	conf *config.Chat
}

// NewRepo creates a mysql client and returns a Repo
func NewRepo(dbConf *config.Chat) *Repo {
	repo := Repo{
		conf:     dbConf,
	}
	return &repo
}
