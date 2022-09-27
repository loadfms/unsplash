package main

import (
	"fmt"

	"github.com/loadfms/unsplash/services"
)

func main() {
	configs, err := services.LoadConfig("config.yaml")

	if err != nil {
		fmt.Println("There's a error with your config.yaml file.")
	}

	fmt.Println("Fetching some random and beautifull wallpaper for you! Just a sec...")
	services.Run(configs)
}
