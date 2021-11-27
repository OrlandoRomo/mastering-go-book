package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	var config *string = flag.String("c", "config_yaml", "path of the configuration file")
	flag.Parse()
	_, err := os.Stat(*config)
	if err == nil {
		fmt.Println("Using user specified config file")
		viper.SetConfigFile(*config)
	}
	if err != nil {
		// Use the default name which is config_yaml from the -c option
		viper.SetConfigName(*config)
		// Find the configuration in /tmp or $HOME or in the current directory
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Using config file ", viper.ConfigFileUsed())

	if viper.IsSet("version") {
		fmt.Println("version", viper.Get("version"))
	} else {
		fmt.Println("version is not set")
	}

	if viper.IsSet("services.backend.volumes") {
		fmt.Println("services.backend.volumes: ", viper.Get("services.backend.volumes"))
	} else {
		fmt.Println("services.backend.volumes is not set")
	}

	if viper.IsSet("services.db.volumes") {
		fmt.Println("services.db.volumes: ", viper.Get("services.db.volumes"))
	} else {
		fmt.Println("services.db.volumes is not set")
	}

	if !viper.IsSet("service.redis") {
		fmt.Println("service.redis is not set")
	}
}
