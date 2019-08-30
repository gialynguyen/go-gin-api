package config

import (
	"github.com/spf13/viper"
	"log"
	_ "path/filepath"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("Error on parsing config file ", err)
	}
}
func GetConfig() *viper.Viper {
	return config
}
