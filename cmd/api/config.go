package main

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var (
	ErrRedisAddressNotFound = errors.New("REDIS_ADDRESS not found in configuration file and env var")
	ErrAPIPortNotFound      = errors.New("GRPC_LISTENER_PORT not found in configuration file and env var")
	ErrDBNameNotFound       = errors.New("DB_NAME not found in configuration file and env var")
	ErrDBPasswordNotFound   = errors.New("DB_PASSWORD not found in configuration file and env var")
	ErrDBUserNotFound       = errors.New("DB_USER not found in configuration file and env var")
	ErrDBPortNotFound       = errors.New("DB_PORT not found in configuration file and env var")
	ErrDBHostNotFound       = errors.New("DB_HOST not found in configuration file and env var")
)

type Config struct {
	GRPCListenerPort string `mapstructure:"GRPC_LISTENER_PORT"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBName           string `mapstructure:"DB_NAME"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	RedisAddr        string `mapstructure:"REDIS_ADDRESS"`
	Debug            bool   `mapstructure:"DEBUG"`
}

func LoadConfig(path string) (Config, error) { //nolint: cyclop
	v := viper.New()

	v.AddConfigPath(path)
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil { // nolint: nestif
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) { // nolint: nestif
			if err = v.BindEnv("DB_HOST"); err != nil {
				return Config{}, ErrDBHostNotFound
			}

			if err = v.BindEnv("DB_PORT"); err != nil {
				return Config{}, ErrDBPortNotFound
			}

			if err = v.BindEnv("DB_USER"); err != nil {
				return Config{}, ErrDBUserNotFound
			}

			if err = v.BindEnv("DB_PASSWORD"); err != nil {
				return Config{}, ErrDBPasswordNotFound
			}

			if err = v.BindEnv("DB_NAME"); err != nil {
				return Config{}, ErrDBNameNotFound
			}

			if err = v.BindEnv("GRPC_LISTENER_PORT"); err != nil {
				return Config{}, ErrAPIPortNotFound
			}

			if err = v.BindEnv("REDIS_ADDRESS"); err != nil {
				return Config{}, ErrRedisAddressNotFound
			}
		}
	}

	var config Config

	if err := v.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("config unmarshal error: %w", err)
	}

	return config, nil
}
