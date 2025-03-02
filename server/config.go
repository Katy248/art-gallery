package main

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func SetupConfiguration() {
	viper.SetConfigName("art-gallery-conf")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$XDG_CONFIG_DIR/art-gallery/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Infof("Config file not found: %s", err)
		} else {
			log.Errorf("Error while reading config file: %s", err)
			// Config file was found but another error was produced
		}
	}
}
