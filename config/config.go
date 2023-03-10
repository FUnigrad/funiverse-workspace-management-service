package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Enviroment string `mapstructure:"ENV"`
	Port       string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName(os.Getenv("ENV"))

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
