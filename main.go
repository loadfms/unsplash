package main

import (
	"fmt"

	"github.com/loadfms/unsplash/services"
)

func main() {
	configs, err := services.LoadConfig("config.yaml")

	if err != nil {
		fmt.Println("There's an error with your config.yaml file.")
		return
	}

	fmt.Println("Fetching some random and beautiful wallpaper for you! Just a sec...")
	err = services.Run(configs)
	if err != nil {
		fmt.Println("There was an error!")
		return
	}
}
