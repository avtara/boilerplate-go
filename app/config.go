package app

import (
	"github.com/spf13/viper"
)

func (cfg *App) InitViper() (err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	return
}
