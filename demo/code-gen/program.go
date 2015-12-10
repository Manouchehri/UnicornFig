package main

import (
	"./config"
	"fmt"
)

func main() {
	configuration, err := config.LoadConfigJson("config/out.json")
	fmt.Println(configuration, err)
}
