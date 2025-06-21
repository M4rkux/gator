package main

import (
	"fmt"

	"github.com/m4rkux/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading the config file:", err)
	}

	err = config.SetUser("Marcus", cfg)
	if err != nil {
		fmt.Println("Error setting the current user:", err)
	}

	cfgUpdated, err := config.Read()
	if err != nil {
		fmt.Println("Error reading the config file:", err)
	}

	fmt.Println(cfgUpdated)
}
