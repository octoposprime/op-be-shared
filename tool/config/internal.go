package tconfig

import (
	"os"

	"gopkg.in/yaml.v2"
)

type InternalConfig struct {
	Debug   bool `yaml:"debug"`
	Local   bool `yaml:"local"`
	Restapi struct {
		ProbeHost   string `yaml:"probe_host"`
		ProbePort   string `yaml:"probe_port"`
		GraphqlHost string `yaml:"graphql_host"`
		GraphqlPort string `yaml:"graphql_port"`
	} `yaml:"restapi"`
	Grpc struct {
		LoggerHost         string `yaml:"logger_host"`
		LoggerPort         string `yaml:"logger_port"`
		UserHost           string `yaml:"user_host"`
		UserPort           string `yaml:"user_port"`
		AuthenticationHost string `yaml:"authentication_host"`
		AuthenticationPort string `yaml:"authentication_port"`
		AuthorizationHost  string `yaml:"authorization_host"`
		AuthorizationPort  string `yaml:"authorization_port"`
	} `yaml:"grpc"`
}

var InternalConfigInstance *InternalConfig

func GetInternalConfigInstance() *InternalConfig {
	if InternalConfigInstance == nil {
		InternalConfigInstance = &InternalConfig{}
		InternalConfigInstance.ReadConfig()
	}
	return InternalConfigInstance
}

var InternalConfigPath string = "config/internal.yml"
var InternalConfigTestPath string = "config/internal_test.yml"
var InternalConfigLocalPath string = "config/internal_local.yml"

func (c *InternalConfig) ReadConfig() {

	configPath := InternalConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = InternalConfigLocalPath
		}
	} else {
		if os.Getenv("TEST") != "" {
			if os.Getenv("TEST") == "true" {
				configPath = InternalConfigTestPath
			}
		}
	}

	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

func IsDebug() bool {
	return GetInternalConfigInstance().Debug
}
