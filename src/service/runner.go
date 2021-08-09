package service

import (
	"errors"
	"github.com/gharsallahmoez/chat/src/config"
	"github.com/gharsallahmoez/chat/src/service/grpc"
	"github.com/sirupsen/logrus"
)

// Runner holds start function te be implemented by a runner
type Runner interface {
	Start(stop chan bool) error
}

// Create creates a runner of type defined in config
func Create(conf *config.Config, logger *logrus.Logger) (Runner, error) {
	var srv Runner

	switch conf.Server.Type {
	case "grpc":
		srv = grpc.NewRunner(conf,logger)
	default:
		return nil, errors.New("InvalidServerTypeError(#{conf.Server.Type})")
	}
	return srv, nil
}

