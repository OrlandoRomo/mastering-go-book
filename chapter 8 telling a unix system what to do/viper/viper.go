package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	max := viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS ", max)
	viper.Set("GOMAXPROCS", 10)
	max = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS ", max)

	viper.BindEnv("NEW_VARIABLE")
	newVariable := viper.Get("NEW_VARIABLE")
	if newVariable == nil {
		fmt.Println("NEW_VARIABLE not defined")
		return
	}

	fmt.Println("NEW_VARIABLE ", newVariable)
}
