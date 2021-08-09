package config_test

import (
	"github.com/gharsallahmoez/chat/src/config"
	tt "github.com/gharsallahmoez/chat/src/testdata/config"
	"os"
	"testing"
)

const ConfEnvVar = "CONFIGOR_ENV"

// Test getting configuration from config.yaml file and parse it to config struct
func TestMakeConfig(t *testing.T) {
	for _, testCase := range tt.CreateTTMakeConf() {
		t.Run(testCase.Name, func(t *testing.T) {
			os.Setenv(ConfEnvVar, testCase.EnvVar)
			conf, err := config.MakeConfig()
			if err != nil ||
				(testCase.IsDev && conf.Parameters.Env != "dev") ||
				(testCase.IsProd && conf.Parameters.Env != "prod") {
				t.Errorf("expected loading config for %v environment, got error: %v", testCase.EnvVar, err)
			}
		})
	}
}
