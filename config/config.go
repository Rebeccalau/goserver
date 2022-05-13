package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Environment string
	First       string
	Other       string
}

func NewConfiguration(env string) Configuration {
	config := Configuration{}
	path, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/configfiles/%sconfig.json", path, env)

	fmt.Println(filePath)
	file, _ := os.Open(filePath)

	error := json.NewDecoder(file).Decode(&config)

	if error != nil {
		fmt.Println("error:", error)
	}

	fmt.Println("Configuration Loaded")
	return config
}
