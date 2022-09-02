package config

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigFile(`config.json`)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
