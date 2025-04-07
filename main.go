package main

import (
	"blog-aggregator/internal/config"
	"fmt"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		panic(err)
	}

	err = configFile.SetUser("RandomUser")
	if err != nil {
		panic(err)
	}

	configFile, err = config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println(configFile)
}
