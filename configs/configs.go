package configs

import (
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Configuration struct {
		Database `yaml:"database"`
	}

	Database struct {
		MaxPoolSize int    `yaml:"max_pool_size"`
		URL         string `yaml:"url"`
	}
)

func GetConfigurations() *Configuration {
	var config Configuration
	err := cleanenv.ReadConfig(GetConfigPath(), &config)
	if err != nil {
		panic(err)
	}

	return &config
}

func GetConfigPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b) + "/config.yml"
}
