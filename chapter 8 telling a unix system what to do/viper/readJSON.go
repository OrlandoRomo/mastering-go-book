package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	//Define the config file we want to read
	viper.SetConfigType("json")
	//Use the path of the file we want to read
	viper.SetConfigFile("./config.json")
	fmt.Println("Using file ", viper.ConfigFileUsed())
	// Read the file
	viper.ReadInConfig()

	if viper.IsSet("viper_config.version") {
		fmt.Println("viper_config.version", viper.Get("viper_config.version"))
	} else {
		fmt.Println("viper_config.version is not set")
	}

	if viper.IsSet("viper_config.language") {
		fmt.Println("viper_config.language: ", viper.Get("viper_config.language"))
	} else {
		fmt.Println("viper_config.language is not set")
	}

	if viper.IsSet("viper_config.language_version") {
		fmt.Println("viper_config.language_version: ", viper.Get("viper_config.language_version"))
	} else {
		fmt.Println("viper_config.language_version is not set")
	}

	if !viper.IsSet("viper_config.version_app") {
		fmt.Println("viper_config.version_app is not set")
	}
}
