package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DbUri          string `mapstructure:"NEO4J_URI"`
	DbContainerUri string `mapstructure:"NEO4J_CONTAINER_URI"`
	DbUsername     string `mapstructure:"NEO4J_USERNAME"`
	DbPassword     string `mapstructure:"NEO4J_PASSWORD"`

	ContentCloudUri string `mapstructure:"IMG_CLOUD_URL"`
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
