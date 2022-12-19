package main

import "github.com/spf13/viper"

func ReadConfiguration() error {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
