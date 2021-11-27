package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	flag.String("name", "Orlando", "Name of the user")
	flag.Int("age", 24, "Age of the user")
	flag.Bool("is-handsome", true, "wether the user is handsome or not")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	fmt.Println("Name: ", viper.Get("name"))
	fmt.Println("Age: ", viper.Get("age"))
	fmt.Println("Handsome: ", viper.Get("is-handsome"))
}
