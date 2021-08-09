package main

import (
	"errors"
	"github.com/gharsallahmoez/chat/src/config"
	"github.com/gharsallahmoez/chat/src/service"
	"github.com/gharsallahmoez/chat/src/utils"
	"github.com/sirupsen/logrus"
)
var (
	stop = make(chan bool)
)
func getServer(conf *config.Config, logger *logrus.Logger) (service.Runner, error) {
	srv, err := service.Create(conf, logger)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

func realMain(srv service.Runner, logger *logrus.Logger) {
	logger.Debug("Configuration parsed successfully")

	logger.Debug("Starting serviceRequest service...")
	err := srv.Start(stop)
	if err != nil {
		logger.Panicf(errors.New("failed to start the service").Error())
	}
	select {
	case <-stop:
	}
}

func main() {
	logger := utils.GetLogger()
	conf, err := config.MakeConfig()
	if err != nil {
		logger.Fatalf(errors.New( "failed to create the configuration").Error())
	}
	logger.Debug("Configuration parsed successfully")

	srv, err := getServer(conf, logger)
	if err != nil {
		logger.Fatal(errors.New( "failed to create the server").Error())
	}
	realMain(srv, logger)
}
