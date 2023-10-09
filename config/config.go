package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Api    ApiConfig      `mapstructure:",squash"`
	Db     DatabaseConfig `mapstructure:",squash"`
	Logger Logger         `mapstructure:",squash"`
}

type ApiConfig struct {
	Binary         string `mapstructure:"API_BINARY" validate:"required"`
	Mode           string `mapstructure:"API_MODE" validate:"required"`
	Port           int    `mapstructure:"API_PORT" validate:"required"`
	Version        string `mapstructure:"API_VERSION" validate:"required"`
	OriginsAllowed string `mapstructure:"API_ORIGINS_ALLOWED" validate:"required"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"DB_DRIVER" validate:"required"`
	Host     string `mapstructure:"DB_HOST" validate:"required"`
	Port     string `mapstructure:"DB_PORT" validate:"required"`
	User     string `mapstructure:"DB_USER" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
}

type Logger struct {
	Development bool   `mapstructure:"LOGGER_DEVELOPMENT" validate:"required"`
	Encoding    string `mapstructure:"LOGGER_ENCODING" validate:"required"`
	Level       string `mapstructure:"LOGGER_LEVEL" validate:"required"`
}

// Load reads in an env file and loads into a config struct
func Load(path string) (config *Config, err error) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(".env." + env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	validate := validator.New()
	if err = validate.Struct(config); err != nil {
		return
	}

	return
}
