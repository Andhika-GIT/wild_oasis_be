package config

import (
	"log"

	"github.com/spf13/viper"
)


func NewViper() *viper.Viper{
	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	if err != nil {
        log.Fatalf("Error reading .env file, %s", err)
    }

	return config

}