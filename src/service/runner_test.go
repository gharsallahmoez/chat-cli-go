package service_test

import (
	"github.com/gharsallahmoez/chat/src/config"
	svc "github.com/gharsallahmoez/chat/src/service"
	"github.com/gharsallahmoez/chat/src/utils"
	"testing"
)

func TestCreate(t *testing.T) {

	conf, _ := config.MakeConfig()
	log := utils.GetLogger()

	t.Run("failure in runner creation", func(t *testing.T) {
		serverRunner := config.Server{
			Type: "unknown",
		}
		wrongConf := config.Config{
			Server: serverRunner,
		}
		runner, err := svc.Create(&wrongConf, log)
		if err == nil || runner != nil {
			t.Error("expected error, got nil")
		}
	})
	t.Run("success runner creation", func(t *testing.T) {
		_, err := svc.Create(conf, log)
		if err != nil {
			t.Errorf("expected error, got %v", err)
		}
	})
}
