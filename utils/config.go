package utils

import "github.com/spf13/viper"

func GetConfig(nameConfig, def string) string {
	if viper.IsSet(nameConfig) {
		return viper.GetString(nameConfig)
	}

	return def
}
