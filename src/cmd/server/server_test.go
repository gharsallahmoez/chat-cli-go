package main

import (
	"github.com/gharsallahmoez/chat/src/config"
	"github.com/gharsallahmoez/chat/src/service"
	"github.com/gharsallahmoez/chat/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_getServer(t *testing.T) {
	logger := utils.GetLogger()
	t.Run("invalid config", func(t *testing.T) {
		conf := config.Config{Server: config.Server{Type: "unknown"}}

		_, err := getServer(&conf, logger)
		if err == nil {
			t.Errorf("expected error got: %v", err)
		}
	})
	t.Run("valid config", func(t *testing.T) {
		conf, _ := config.MakeConfig()
		_, err := getServer(conf, logger)
		if err != nil {
			t.Errorf("expected success got error:%v", err)
		}
	})
}

func Test_realMain(t *testing.T) {
	logger := utils.GetLogger()
	t.Run("invalid runner: wrong port", func(t *testing.T) {
		conf, _ := config.MakeConfig()
		conf.Server.Port = "wrongPort"
		wrongRunner, _ := service.Create(conf, logger)
		go func() {
			time.Sleep(1 * time.Second)
			stop <- true
			time.Sleep(1 * time.Second)
		}()
		assert.Panics(t, func() {
			realMain(wrongRunner, logger)
		}, "starting runner should panic")

	})
	t.Run("valid runner", func(t *testing.T) {
		conf, _ := config.MakeConfig()
		validRunner, _ := service.Create(conf, logger)
		assert.NotPanics(t, func() {
			go func() {
				time.Sleep(1 * time.Second)
				stop <- true
				time.Sleep(1 * time.Second)
			}()
			realMain(validRunner, logger)
		}, "starting runner shouldn't panic")
	})
}

