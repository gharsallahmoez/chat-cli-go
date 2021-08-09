package config

import (
	"github.com/jinzhu/configor"
	"path"
	"runtime"
	"time"
)

type Parameters struct {
	Env     string `default:"dev" env:"CONFIGOR_ENV"`
	Timeout time.Duration
}

type Server struct {
	Type     string
	Host     string        `default:"localhost" env:"SERVER_HOST"`
	Port     string        `default:"50053" env:"SERVER_PORT"`
	Deadline time.Duration `env:"GRPC_DEADLINE"`
}

type Chat struct {
	Type       string `default:"dynamodb" env:"DB"`
	DbName     string
	TableName  string
	Host       string
	Username   string `env:"DB_USER"`
	Password   string `env:"DB_PASS"`
}

type Database struct {
	Chat Chat
}

type Config struct {
	Parameters Parameters
	Server     Server
	Database   Database
}

// MakeConfig sets the application config, change configuration according to env variable, and set the host of database
// returns config,error
func MakeConfig() (conf *Config, err error) {
	var configFilePath string
	config := configor.New(&configor.Config{})
	switch config.GetEnvironment() {
	case "prod":
		configFilePath = "../../config/config.prod.yml"
	default:
		configFilePath = "../../config/config.dev.yml"
	}
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), configFilePath)
	// Load configuration from yml file
	conf = new(Config)
	err = config.Load(conf, filepath)
	// return the configuration and the error
	return conf,err
}
