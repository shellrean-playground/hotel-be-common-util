package provider

import (
	"github.com/shellrean-playground/hotel-be-common-util/config"
	"github.com/shellrean-playground/hotel-be-common-util/configuration"
)

func GetConfiguration() *config.Config {
	configPath, err := configuration.ParseFlags()
	if err != nil {
		panic(err)
	}
	conf, err := configuration.New(configPath)
	if err != nil {
		panic(err)
	}
	return conf
}
