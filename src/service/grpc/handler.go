package grpc

import (
	"fmt"
	"github.com/gharsallahmoez/chat/src/config"
	pb "github.com/gharsallahmoez/chat/src/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type SvcRunner struct {
	Config *config.Config
	*logrus.Logger
	Users map[string]struct{}
}

// NewRunner creates a runner
func NewRunner(conf *config.Config, logger *logrus.Logger) *SvcRunner {
	users := make(map[string]struct{})
	return &SvcRunner{
		Config: conf,
		Logger: logger,
		Users: users,
	}
}

func (runner *SvcRunner) GetLogHandler() *logrus.Logger {
	return runner.Logger
}

func (runner *SvcRunner) GetConfig() *config.Config {
	return runner.Config
}

// Start starts the runner , this method will be called in main function to run the server
func (runner *SvcRunner) Start(stop chan bool) error {
	// create listener
	listener, err := net.Listen(
		"tcp",
		":"+ runner.Config.Server.Port,
	)
	if err != nil {
		return err
	}

	// TODO : add interceptors
	// TODO : add server helth check

	server := grpc.NewServer()
	pb.RegisterServicesServer(server,runner)

	go func() {
		runner.Logger.Info("Starting server...")
		err = server.Serve(listener)
		if err != nil {
			runner.Logger.Errorf("server can't be started, err: %v", err)
			panic(fmt.Sprintf("server can't be started, err: %v", err))
		}
	}()

	go func() {
		for {
			select {
			case value := <-stop:
				if value == true {
					runner.Logger.Info("server stopped")
					server.GracefulStop()
				}
			}
		}
	}()
	return nil
}

