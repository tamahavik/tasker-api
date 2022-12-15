package utils

import "github.com/spf13/viper"

type Configuration interface {
	ReadConfiguration() error
}

type configurationStruct struct {
}

func NewConfiguration() Configuration {
	return &configurationStruct{}
}

func (c *configurationStruct) ReadConfiguration() error {
	viper.SetConfigType("env")
	viper.AddConfigPath("../.")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
