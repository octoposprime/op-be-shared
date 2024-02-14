package tconfig

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	Redis struct {
		Enabled  bool   `yaml:"enabled"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
}

var RedisConfigPath string = "config/redis.yml"
var RedisConfigTestPath string = "config/redis_test.yml"
var RedisConfigLocalPath string = "config/redis_local.yml"

var RedisConfigInstance *RedisConfig

func GetRedisConfigInstance() *RedisConfig {
	if RedisConfigInstance == nil {
		RedisConfigInstance = &RedisConfig{}
		RedisConfigInstance.ReadConfig()
	}
	return RedisConfigInstance
}

func (c *RedisConfig) ReadConfig() {
	configPath := RedisConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = RedisConfigLocalPath
		}
	} else {
		if os.Getenv("TEST") != "" {
			if os.Getenv("TEST") == "true" {
				configPath = RedisConfigTestPath
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
	c.Redis.Password = os.Getenv("REDIS_PASSWORD")
}
