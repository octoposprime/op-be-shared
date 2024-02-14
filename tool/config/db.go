package tconfig

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	PostgresDb struct {
		Enabled  bool   `yaml:"enabled"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgresdb"`
}

var DbConfigPath string = "config/postgres.yml"
var DbConfigTestPath string = "config/postgres_test.yml"
var DbConfigLocalPath string = "config/postgres_local.yml"

var DbConfigInstance *DbConfig

func GetDbConfigInstance() *DbConfig {
	if DbConfigInstance == nil {
		DbConfigInstance = &DbConfig{}
		DbConfigInstance.ReadConfig()
	}
	return DbConfigInstance
}

func (c *DbConfig) ReadConfig() {
	configPath := DbConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = DbConfigLocalPath
		}
	} else {
		if os.Getenv("TEST") != "" {
			if os.Getenv("TEST") == "true" {
				configPath = DbConfigTestPath
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

	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			if err := godotenv.Load(); err != nil {
				panic("Error loading .env file")
			}
		}
	}
	c.PostgresDb.UserName = os.Getenv("POSTGRES_USERNAME")
	c.PostgresDb.Password = os.Getenv("POSTGRES_PASSWORD")
}
