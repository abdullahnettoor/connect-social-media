package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DbUri      string `mapstructure:"NEO4J_URI"`
	DbUsername string `mapstructure:"NEO4J_USERNAME"`
	DbPassword string `mapstructure:"NEO4J_PASSWORD"`
}

func LoadConfig() (c *Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Println(err)
		return
	}

	return c, err
}
