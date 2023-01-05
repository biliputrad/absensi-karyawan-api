package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type Config struct {
	GinMode    string `mapstructure:"GIN_MODE"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbTz       string `mapstructure:"DB_TZ"`
	DbLogLevel string `mapstructure:"DB_LOG_LEVEL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func GetDBLogLevel(configLogLevel string) logger.LogLevel {
	var logLevel logger.LogLevel
	if configLogLevel == "info" {
		logLevel = logger.Info
	} else if configLogLevel == "warn" {
		logLevel = logger.Warn
	} else if configLogLevel == "error" {
		logLevel = logger.Error
	}

	return logLevel
}
