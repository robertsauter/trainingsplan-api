package lib

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`

	ServerPort   string `mapstructure:"PORT"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	loadingError := viper.ReadInConfig()
	if loadingError != nil {
		return nil, errors.New("could not load config")
	}

	var config Config
	unmarshalError := viper.Unmarshal(&config)
	if unmarshalError != nil {
		return nil, errors.New("could not unpack config")
	}

	return &config, nil
}
