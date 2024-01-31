package tconfig

import (
	"os"

	"github.com/joho/godotenv"
)

type JwtConfig struct {
	Jwt struct {
		SecretKey string `yaml:"secretkey"`
	} `yaml:"jwt"`
}

//var JwtConfigPath string = "config/jwt.yml"
//var JwtConfigTestPath string = "config/jwt_test.yml"
//var JwtConfigLocalPath string = "config/jwt_local.yml"

var JwtConfigInstance *JwtConfig

func GetJwtConfigInstance() *JwtConfig {
	if JwtConfigInstance == nil {
		JwtConfigInstance = &JwtConfig{}
		JwtConfigInstance.ReadConfig()
	}
	return JwtConfigInstance
}

func (c *JwtConfig) ReadConfig() {
	/*configPath := JwtConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = JwtConfigLocalPath
		}
	} else {
		if os.Getenv("TEST") != "" {
			if os.Getenv("TEST") == "true" {
				configPath = JwtConfigTestPath
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
	}*/
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			if err := godotenv.Load(); err != nil {
				panic("Error loading .env file")
			}

		}
	}
	c.Jwt.SecretKey = os.Getenv("JWT_SECRET_KEY")
}
