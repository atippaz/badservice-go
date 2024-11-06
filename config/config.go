package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type (
	Database struct {
		Url          string `mapstructure:"url" validate:"required"`
		DatabaseName string `mapstructure:"databaseName" validate:"required"`
	}

	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
	}

	Config struct {
		Database *Database `mapstructure:"database" validate:"required"`
		Server   *Server   `mapstructure:"server" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validate := validator.New()

		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
